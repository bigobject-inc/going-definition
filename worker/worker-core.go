package worker

import (
	"context"
)

// Core Going Worker core interface
type Core interface {
	GetID() string
	GetStatus() string
	// GetSetting() Setting
	GetChannelIn() *Channel
	GetChannelOut() []*Channel
	GetFromWorkerID() []string
	GetToWorkerID() []string
	GetWorker() Worker
	GetSharedMemoryValue(sID, key string) (string, error)
	SetWorker(worker Worker)
	SetChannelIn(channel *Channel)
	SetChannelOut(channel *Channel)
	SetFromWorkerID(id string)
	SetToWorkerID(id string)
	SetTimeout(seconds int, data interface{})
	SetSharedMemory(creatorName string, dataSourceName string, params interface{}) (string, error)
	SetSharedMemoryValue(sID, key, value string, expired int) error
	SendDataNext(data interface{}) error
	Start(ctx context.Context) error
	Stop() error
}

// Channel Going worker channel
type Channel struct {
	ID      string
	From    string
	To      string
	Channel chan interface{}
}
