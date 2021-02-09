package main

import "ps/pkg/process"

type SubmitProcessRequest struct {
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type SubmitProcessResponse struct {
	Pid string `json:"pid"`
}

type GetProcessStatusResponse struct {
	process.ProcessInfo `json:"status"`
}
