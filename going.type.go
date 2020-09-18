// ----------------------------------------------------------------------
// Copyright (c) 2019 BigObject Inc. ("BigObject")
// All Rights Reserved.
//
// Use of, copying, modifications to, and distribution of this software
// and its documentation without BigObject's
// written permission can result in the violation of U.S. Copyright
// and Patent laws. Violators will be prosecuted to the highest
// extent of the applicable laws.
//
// BIGOBJECT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
// THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
// TO THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE, OR NON-INFRINGEMENT.
//
// @author
// CloudChen <cloudchen@bigobject.io>
// 2020/06/03
// ----------------------------------------------------------------------

package going

import (
	"errors"
	"os"

	"github.com/bigobject-inc/golib/logger"
)

var (
	// Errs  errors
	Errs Errors = Errors{
		AlreadyInitialized: errors.New("Services is already initialized"),
		NotInitialized:     errors.New("Services is not initialized"),
		AlreadyRunning:     errors.New("Services is already running"),
		NotRunning:         errors.New("Services is not running"),
		ValueNotFound:      errors.New("Unable to find value"),
	}
)

const (
	// PluginSourceWorker Plugin source worker
	PluginSourceWorker = "sourceWorkers"
	// PluginWorker Plugin worker
	PluginWorker = "workers"
)

const (
	// StatusCreated Status created
	StatusCreated = "created"
	// StatusReady Status ready
	StatusReady = "ready"
	// StatusRunning Status running
	StatusRunning = "running"
	// StatusStop Status stop
	StatusStop = "stop"
	// StatusReceived Status received
	StatusReceived = "received"
	// StatusProcessed Status processed
	StatusProcessed = "processed"
	// StatusError Status error
	StatusError = "error"
)

// Errors going errors
type Errors struct {
	AlreadyInitialized error
	NotInitialized     error
	AlreadyRunning     error
	NotRunning         error
	ValueNotFound      error
}

// Logger going logger
type Logger struct {
	Logger  logger.Logger
	LoggerF *os.File
}

// Paths going paths
type Paths struct {
	Root     string
	Logs     string
	Settings string
	Plugin   string
}

// Plugin going plugin
type Plugin struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Path string `json:"path"`
}

// PluginInfo going plugin info
type PluginInfo struct {
	Name         string
	Desc         string
	Version      string
	ParamsSchema string
}
