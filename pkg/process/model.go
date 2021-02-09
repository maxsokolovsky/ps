package process

type ProcessInfo struct {
	Path      string   `json:"path"`
	Args      []string `json:"args"`
	IsRunning bool     `json:"isRunning"`
	Output    []byte   `json:"output"`
	Error     error    `json:"error"`
}
