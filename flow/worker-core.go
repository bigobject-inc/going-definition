package flow

import (
	"context"

	"github.com/bigobject-inc/going-definition/v2/notification"
	"github.com/bigobject-inc/going-definition/v2/sharedmemory"
)

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
	GetStoragePath() string
	GetPrivateStoragePath() string
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
