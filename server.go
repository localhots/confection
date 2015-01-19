package confection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	switch req.URL.Path {
	case "/fields.json":
		s.fieldsHandler(w, req)
	case "/save":
		s.saveHandler(w, req)
	default:
		http.NotFound(w, req)
	}
}

func (s *server) fieldsHandler(w http.ResponseWriter, req *http.Request) {
	jsn, err := json.Marshal(s.manager.conf.meta(""))
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf8")
	w.Write(jsn)
}

func (s *server) saveHandler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	s.manager.importJson(b)
}
