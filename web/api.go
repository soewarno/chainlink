package web

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/manyminds/api2go/jsonapi"
)

const (
	// PaginationDefault is the number of records to supply from a paginated
	// request when no size param is supplied.
	PaginationDefault = 25

	// MediaType is the response header for JSONAPI documents.
	MediaType = "application/vnd.api+json"

	// KeyNextLink is the name of the key that contains the HREF for the next
	// document in a paginated response.
	KeyNextLink = "next"
	// KeyPreviousLink is the name of the key that contains the HREF for the
	// previous document in a paginated response.
	KeyPreviousLink = "prev"
)

// ParsePaginatedRequest parses the parameters that control pagination for a
// collection request, returning the size and offset if specified, or a
// sensible default.
func ParsePaginatedRequest(sizeParam, pageParam string) (int, int, int, error) {
	var err error
	page := 1
	size := PaginationDefault

	if sizeParam != "" {
		if size, err = strconv.Atoi(sizeParam); err != nil || size < 1 {
			return 0, 0, 0, fmt.Errorf("invalid size param, error: %+v", err)
		}
	}

	if pageParam != "" {
		if page, err = strconv.Atoi(pageParam); err != nil || page < 1 {
			return 0, 0, 0, fmt.Errorf("invalid page param, error: %+v", err)
		}
	}

	offset := (page - 1) * size
	return size, page, offset, nil
}

func paginationLink(url url.URL, size, page int) jsonapi.Link {
	query := url.Query()
	query.Add("size", strconv.Itoa(size))
	query.Add("page", strconv.Itoa(page))
	url.RawQuery = query.Encode()
	return jsonapi.Link{Href: url.String()}
}

func nextLink(url url.URL, size, page int) jsonapi.Link {
	return paginationLink(url, size, page+1)
}

func prevLink(url url.URL, size, page int) jsonapi.Link {
	return paginationLink(url, size, page-1)
}

// NewJSONAPIResponse returns a JSONAPI response for a single resource.
func NewJSONAPIResponse(resource interface{}) ([]byte, error) {
	document, err := jsonapi.MarshalToStruct(resource, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource to struct: %+v", err)
	}

	return json.Marshal(document)
}

// NewPaginatedResponse returns a jsonapi.Document with links to next and previous collection pages
func NewPaginatedResponse(url url.URL, size, page, count int, resource interface{}) ([]byte, error) {
	document, err := jsonapi.MarshalToStruct(resource, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource to struct: %+v", err)
	}

	document.Meta = make(jsonapi.Meta)
	document.Meta["count"] = count

	document.Links = make(jsonapi.Links)
	if count > size {
		if page*size < count {
			document.Links[KeyNextLink] = nextLink(url, size, page)
		}
		if page > 1 {
			document.Links[KeyPreviousLink] = prevLink(url, size, page)
		}
	}
	return json.Marshal(document)
}

// ParsePaginatedResponse parse a JSONAPI response for a document with links
func ParsePaginatedResponse(input []byte, resource interface{}, links *jsonapi.Links) error {
	err := ParseJSONAPIResponse(input, resource)
	if err != nil {
		return err
	}

	// Unmarshal using the stdlib Unmarshal to extract the links part of the document
	document := jsonapi.Document{}
	err = json.Unmarshal(input, &document)
	if err != nil {
		return fmt.Errorf("unable to unmarshal links: %+v", err)
	}

	*links = document.Links
	return nil
}

// ParseJSONAPIResponse parses the bytes of the root document and unmarshals it
// into the given resource.
func ParseJSONAPIResponse(input []byte, resource interface{}) error {
	// as is api2go will discard the links
	err := jsonapi.Unmarshal(input, resource)
	if err != nil {
		return fmt.Errorf("web: unable to unmarshal data, %+v", err)
	}

	return nil
}
