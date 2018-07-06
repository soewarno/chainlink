package adapters

import (
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
)

var (
	TaskTypeCopy       = models.NewTaskType("copy")
	TaskTypeEthBytes32 = models.NewTaskType("ethbytes32")
	TaskTypeEthInt256  = models.NewTaskType("ethint256")
	TaskTypeEthUint256 = models.NewTaskType("ethuint256")
	TaskTypeEthTx      = models.NewTaskType("ethtx")
	TaskTypeHTTPGet    = models.NewTaskType("httpget")
	TaskTypeHTTPPost   = models.NewTaskType("httppost")
	TaskTypeJSONParse  = models.NewTaskType("jsonparse")
	TaskTypeMultiply   = models.NewTaskType("multiply")
	TaskTypeNoOp       = models.NewTaskType("noop")
	TaskTypeNoOpPend   = models.NewTaskType("nooppend")
	TaskTypeSleep      = models.NewTaskType("sleep")
)

// Adapter interface applies to all core adapters.
// Each implementation must return a RunResult.
type Adapter interface {
	Perform(models.RunResult, *store.Store) models.RunResult
}

// AdapterWithMinConfs is the interface required for an adapter to be run in
// the job pipeline. In addition to the Adapter interface, implementers must
// specify the number of confirmations required before the Adapter can be run.
type AdapterWithMinConfs interface {
	Adapter
	MinConfs() uint64
}

// MinConfsWrappedAdapter allows for an adapter to be wrapped so that it meets
// the AdapterWithMinConfsInterface.
type MinConfsWrappedAdapter struct {
	Adapter
	ConfiguredConfirmations uint64
}

// MinConfs specifies the number of block confirmations
// needed to run the adapter.
func (wa MinConfsWrappedAdapter) MinConfs() uint64 {
	return wa.ConfiguredConfirmations
}

// For determines the adapter type to use for a given task.
func For(task models.TaskSpec, store *store.Store) (AdapterWithMinConfs, error) {
	var ac Adapter
	var err error
	switch task.Type {
	case TaskTypeCopy:
		ac = &Copy{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeEthBytes32:
		ac = &EthBytes32{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeEthInt256:
		ac = &EthInt256{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeEthUint256:
		ac = &EthUint256{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeEthTx:
		ac = &EthTx{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeHTTPGet:
		ac = &HTTPGet{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeHTTPPost:
		ac = &HTTPPost{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeJSONParse:
		ac = &JSONParse{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeMultiply:
		ac = &Multiply{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeNoOp:
		ac = &NoOp{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeNoOpPend:
		ac = &NoOpPend{}
		err = unmarshalParams(task.Params, ac)
	case TaskTypeSleep:
		ac = &Sleep{}
		err = unmarshalParams(task.Params, ac)
	default:
		bt, err := store.FindBridge(task.Type.String())
		if err != nil {
			return nil, fmt.Errorf("%s is not a supported adapter type", task.Type)
		}
		return &Bridge{bt}, nil
	}
	wa := MinConfsWrappedAdapter{
		Adapter:                 ac,
		ConfiguredConfirmations: store.Config.MinIncomingConfirmations,
	}
	return wa, err
}

func unmarshalParams(params models.JSON, dst interface{}) error {
	bytes, err := params.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}
