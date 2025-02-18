package handlers

func ApiResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{"message": "success", "data": data}
}
