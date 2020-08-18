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

package notification

import (
	"github.com/bigobject-inc/golib/logger"
)

// Notification notification
type Notification interface {
	Close() error
	Publish(topic, data string) error
	Subscribe(topic string, handle SubscribeHandle) (string, error)
	Unsubscribe(sID string) error
}

// DataWithWorkersName notification data with workers name
type DataWithWorkersName struct {
	Data        string
	WorkersName []string
}

// Creator shared memory creator
type Creator func(l logger.Logger, dsn string, params interface{}) (Notification, error)

// NotificationsMap notifications map
type NotificationsMap map[string]Notification

// SubscribeHandle flow subscribe handle
type SubscribeHandle func(data string) error
