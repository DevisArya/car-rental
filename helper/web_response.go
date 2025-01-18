package helper

// response sukses dengan data
func NewResponseWithData(statusCode int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	}
}

// response sukses dengan data
func NewResponseWithDatas(statusCode int, message string, data interface{}, totalPage, totalCount, limit, offset int) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
		"pagination": map[string]int{
			"total_page":  totalPage,
			"total_count": totalCount,
			"limit":       limit,
			"offset":      offset,
		},
	}
}

// response sukses tanpa data
func NewResponse(statusCode int, message string) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
	}
}

// response error
func NewErrorResponse(statusCode int, errors []string) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": statusCode,
		"errors":     errors,
	}
}
