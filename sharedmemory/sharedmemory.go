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

package sharedmemory

// SharedMemory shared memory
type SharedMemory interface {
	Close() error
	Get(key string) (string, error)
	GetByKeys(keys []string) ([]string, error)
	Keys(pattern string) ([]string, error)
	Remove(key string) error
	RemoveByKeys(keys []string) error
	Put(key string, value string, expired int) error
}

// Creator shared memory creator
type Creator func(dsn string, params interface{}) (SharedMemory, error)

// SharedMemoriesMap shared memories map
type SharedMemoriesMap map[string]SharedMemory
