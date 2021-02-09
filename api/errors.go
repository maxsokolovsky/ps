package main

import (
	"errors"
)

var ErrPathIsRequired = errors.New("path is required")
var ErrPidIsRequired = errors.New("pid is required")
