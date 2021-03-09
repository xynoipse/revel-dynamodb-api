package response

type (
	// JSONResponse struct for response json
	JSONResponse struct {
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
		Success bool        `json:"success"`
		Message string      `json:"message,omitempty"`
	}
)

// Success func return success with optional data JSONResponse
func Success(data interface{}, message string) JSONResponse {
	return JSONResponse{
		Data:    data,
		Success: true,
		Message: message,
	}
}

// Failed func return failed or error JSONResponse
func Failed(errors interface{}, message string) JSONResponse {
	return JSONResponse{
		Errors:  errors,
		Success: false,
		Message: message,
	}
}
