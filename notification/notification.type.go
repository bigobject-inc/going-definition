package notification

import (
	"github.com/bigobject-inc/going-definition/v2/logger"
)

// DataWithWorkersName notification data with workers name
type DataWithWorkersName struct {
	Data        interface{}
	WorkersName []string
}

// Creator shared memory creator
type Creator func(l logger.Logger, dsn string, params interface{}) (Notification, error)

// NotificationsMap notifications map
type NotificationsMap map[string]Notification

// SubscribeHandle flow subscribe handle
type SubscribeHandle func(data string) error
