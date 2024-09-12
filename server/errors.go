package server

// e - для конвертации ошибки string в json
func e(e error) map[string]string {
	return map[string]string{
		"error": e.Error(),
	}
}
