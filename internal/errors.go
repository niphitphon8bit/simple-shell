package internal

import "errors"

var ErrExit = errors.New("exit requested")
var ErrNoPath = errors.New("path required")
