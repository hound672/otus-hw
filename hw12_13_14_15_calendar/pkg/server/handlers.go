package server

import "net/http"

func helloWorld(resp http.ResponseWriter, _ *http.Request) {
	_, _ = resp.Write([]byte(`hello world`))
}
