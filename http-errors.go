// Copyright (c) 2018. Flying Jamnik Authors
// license that can be found in the LICENSE file.

package errors

import "errors"

var (
	InternalServerErr      = errors.New("Internal server error!")
	UnauthorizedRequestErr = errors.New("Unauthorized request error!")
)
