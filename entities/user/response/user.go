package response

// DisplayResponse -()
func DisplayResponse(data interface{}, message string, status ...string) map[string]interface{} {
	response := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}

	if status == nil {
		response["status"] = "OK"
	} else {
		response["status"] = status[0]
	}

	return response
}
