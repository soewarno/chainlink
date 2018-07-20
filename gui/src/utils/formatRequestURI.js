import url from 'url'

export default (path, query = {}, options = {}) => {
  let formatOptions = Object.assign(
    {pathname: path, query: query},
    options
  )

  return url.format(formatOptions)
}
