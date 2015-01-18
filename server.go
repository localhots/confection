package confection

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type (
	server struct {
		manager *Manager
		port    int
	}
)

func (s *server) start() {
	portStr := ":" + strconv.Itoa(s.port)

	fmt.Println("Configuration server is available at http://127.0.0.1" + portStr)
	if err := http.ListenAndServe(portStr, s); err != nil {
		panic(err)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	jsn, err := json.Marshal(s.manager.conf.meta(""))
	if err != nil {
		panic(err)
	}

	w.Write(jsn)
}
