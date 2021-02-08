package main

type SubmitProcessRequest struct {
	Path string   `json:"path"`
	Args []string `json:"args"`
}

type SubmitProcessResponse struct {
	Pid string `json:"pid"`
}
