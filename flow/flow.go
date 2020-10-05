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

	"github.com/bigobject-inc/going-definition/v2/logger"
	"github.com/bigobject-inc/going-definition/v2/notification"
	"github.com/bigobject-inc/going-definition/v2/sharedmemory"
)

// Flow Going Flow interface
type Flow interface {
	GetID() string
	GetLogger() logger.Logger
	GetNotification(nID string) (notification.Notification, error)
	GetSetting() Setting
	GetSharedMemory(sID string) (sharedmemory.SharedMemory, error)
	GetSourceWorkerCoreByID(id string) (SourceWorkerCore, error)
	GetSourceWorkerCores() []SourceWorkerCore
	GetSourceWorkerCoresByFilter(filter SourceWorkerCoreFilter) []SourceWorkerCore
	GetStatus() string
	GetStatusMessage() string
	GetWorkerCoreByID(id string) (WorkerCore, error)
	GetWorkerCores() []WorkerCore
	GetWorkerCoresByFilter(filter WorkerCoreFilter) []WorkerCore
	SetNotification(creatorName, dataSourceName string, params interface{}) (string, error)
	SetSharedMemory(creatorName, dataSourceName string, params interface{}) (string, error)
	Init(setting Setting, logPath, logLevel string) error
	Defer()
	Start(ctx context.Context) error
	Stop() error
}
