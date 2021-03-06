// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Code generated by "stringer -type=CallType"; DO NOT EDIT

package testutil

import "fmt"

const _CallType_name = "CallGetCallPutCallWriteToCallReadFrom"

var _CallType_index = [...]uint8{0, 7, 14, 25, 37}

func (i CallType) String() string {
	if i < 0 || i >= CallType(len(_CallType_index)-1) {
		return fmt.Sprintf("CallType(%d)", i)
	}
	return _CallType_name[_CallType_index[i]:_CallType_index[i+1]]
}
