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
)

// Errors going errors
type Errors struct {
	AlreadyInitialized error
	NotInitialized     error
	AlreadyRunning     error
	NotRunning         error
	ValueNotFound      error
}

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
