package sourceworker

import (
	"bitbucket.org/bigobject/going-definitions/flow"
)

// Definition Going source worker definition
type Definition struct {
	Name         string
	Desc         string
	Version      string
	ParamsSchema string
}

// Setting Going source worker setting
type Setting struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Version      string `json:"version"`
	Path         string `json:"path"`
	ParamsSchema string `json:"paramsSchema"`
}

// SourceWorker Going source worker
type SourceWorker struct {
	Key          string
	Name         string
	Desc         string
	Version      string
	Creator      flow.SourceWorkerCreator
	ParamsSchema string
	Path         string
}
