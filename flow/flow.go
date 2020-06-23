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

package flow

import (
	"context"

	"github.com/bigobject-inc/going-definition"
	"github.com/bigobject-inc/going-definition/logger"
	"github.com/bigobject-inc/going-definition/notification"
	"github.com/bigobject-inc/going-definition/sharedmemory"
)

// Flow Going Flow interface
type Flow interface {
	GetID() string
	GetLogger() logger.Logger
	GetNotification(nID string) (notification.Notification, error)
	GetSetting() Setting
	GetSharedMemory(sID string) (sharedmemory.SharedMemory, error)
	GetSourceWorkerCore() SourceWorkerCore
	GetStatus() string
	GetWorkerCoreByID(id string) (WorkerCore, error)
	GetWorkerCores() []WorkerCore
	SetNotification(creatorName, dataSourceName string, params interface{}) (string, error)
	SetSharedMemory(creatorName, dataSourceName string, params interface{}) (string, error)
	Init(setting Setting, logPath, logLevel string) error
	Defer()
	Start(ctx context.Context) error
	Stop() error
}

// SourceWorker Going Source worker interface
type SourceWorker interface {
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(swc SourceWorkerCore, params string) error
	BeforeStart() error
	Start(ctx context.Context) error
	Stop() error
}

// SourceWorkerCore Going source Worker core interface
type SourceWorkerCore interface {
	GetChannelOut() []*Channel
	GetFlow() Flow
	GetID() string
	GetNotification(nID string) (notification.Notification, error)
	GetSetting() SettingSourceWorker
	GetSharedMemory(sID string) (sharedmemory.SharedMemory, error)
	GetSourceWorker() SourceWorker
	GetStatus() string
	SendDataNext(data interface{}) error
	SetChannelOut(chanOut *Channel) error
	SetNotification(creatorName, dataSourceName string, params interface{}) (string, error)
	SetSharedMemory(creatorName, dataSourceName string, params interface{}) (string, error)
	Init(sw SourceWorker, f Flow) error
	Defer()
	Start(ctx context.Context) error
	Stop() error
}

// Worker Going worker interface
type Worker interface {
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(wc WorkerCore, params string) error
	Defer()
	BeforeProcess() error
	Process(data interface{}) (interface{}, error)
	Timeout(seconds int, data interface{}) error
}

// WorkerCore Going Worker core interface
type WorkerCore interface {
	GetChannelIn() *Channel
	GetChannelOut() []*Channel
	GetFlow() Flow
	GetID() string
	GetNotification(nID string) (notification.Notification, error)
	GetSetting() SettingWorker
	GetSharedMemory(sID string) (sharedmemory.SharedMemory, error)
	GetStatus() string
	GetWorker() Worker
	SendDataNext(data interface{}) error
	SetChannelIn(chanIn *Channel) error
	SetChannelOut(chanOut *Channel) error
	SetNotification(creatorName, dataSourceName string, params interface{}) (string, error)
	SetSharedMemory(creatorName, dataSourceName string, params interface{}) (string, error)
	SetTimeout(seconds int, data interface{})
	Init(w Worker, flow Flow) error
	Defer()
	Start(ctx context.Context) error
	Stop() error
}

// Channel flow worker channel
type Channel struct {
	ID string
	C  chan interface{}
}

// Setting Going flow setting
type Setting struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	Desc          string              `json:"desc"`
	ChannelBuffer int                 `json:"channelBuffer"`
	SourceWorker  SettingSourceWorker `json:"sourceWorker"`
	Workers       []SettingWorker     `json:"workers"`
}

// SettingSourceWorker Going flow setting source worker
type SettingSourceWorker struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Version string `json:"version"`
	Params  string `json:"params"`
}

// SettingWorker Going flow setting worker
type SettingWorker struct {
	ID      string   `json:"id"`
	Parents []string `json:"parents"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Desc    string   `json:"desc"`
	Version string   `json:"version"`
	Params  string   `json:"params"`
}

// SourceWorkerCreator Going flow source worker creator
type SourceWorkerCreator func(logger logger.Logger) SourceWorker

// WorkerCreator Going flow worker creator
type WorkerCreator func(logger logger.Logger) Worker
