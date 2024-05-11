package health

import "net/http"

func Read(response http.ResponseWriter, _ *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"status":"ok"}`))
}
