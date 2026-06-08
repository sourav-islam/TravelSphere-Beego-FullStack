package utils

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(data interface{}) JSONResponse {
	return JSONResponse{Success: true, Data: data}
}

func SuccessMessageResponse(message string, data interface{}) JSONResponse {
	return JSONResponse{Success: true, Message: message, Data: data}
}

func ErrorResponse(err string) JSONResponse {
	return JSONResponse{Success: false, Error: err}
}

func ValidationErrorResponse(msg string) JSONResponse {
	return JSONResponse{Success: false, Error: msg, Message: "Validation failed"}
}
