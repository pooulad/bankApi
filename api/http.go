type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

