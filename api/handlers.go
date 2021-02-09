package main

import (
	"encoding/json"
	"net/http"

	"ps/pkg/scheduler"
)

type Handler struct {
	s scheduler.Scheduler
}

func (handler *Handler) SubmitProcess(w http.ResponseWriter, r *http.Request) {
	var spr SubmitProcessRequest
	err := json.NewDecoder(r.Body).Decode(&spr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if spr.Path == "" {
		http.Error(w, "path is required", http.StatusBadRequest)
		return
	}
	pid := handler.s.SubmitProcess(spr.Path, spr.Args...)
	resp := SubmitProcessResponse{Pid: pid}
	json.NewEncoder(w).Encode(resp)
}

func (handler *Handler) CancelProcess(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get(":pid")
	if pid == "" {
		http.Error(w, "pid is required", http.StatusBadRequest)
		return
	}
	err := handler.s.CancelProcess(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (handler *Handler) GetProcessStatus(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get(":pid")
	if pid == "" {
		http.Error(w, "pid is required", http.StatusBadRequest)
		return
	}

	status, err := handler.s.ProcessStatus(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := GetProcessStatusResponse{*status}
	json.NewEncoder(w).Encode(resp)
}
