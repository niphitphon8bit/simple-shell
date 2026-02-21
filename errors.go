package main

import "errors"

var errExit = errors.New("exit requested")
var errNoPath = errors.New("path required")
var errNoCommandFound = errors.New("executable file not found in $PATH")
