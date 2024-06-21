package types

type JSONResponse struct {
	Endpoint string
	Outputs  []FileOutput
}

type FileOutput struct {
	Filename string
	Output   string
	Success  bool
}
