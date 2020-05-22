package worker

import (
	"bitbucket.org/bigobject/going-definitions/flow"
	"bitbucket.org/bigobject/going-definitions/logger"
)

// Definition Going worker definition
type Definition struct {
	Name         string
	Desc         string
	Version      string
	ParamsSchema string
}

// Setting Going worker setting
type Setting struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Version      string `json:"version"`
	Path         string `json:"path"`
	ParamsSchema string `json:"paramsSchema"`
}

// Worker Going worker
type Worker struct {
	Key          string
	Name         string
	Desc         string
	Version      string
	Creator      Creator
	ParamsSchema string
	Path         string
}

// Creator Going worker creator
type Creator func(logger logger.Logger) flow.Worker
