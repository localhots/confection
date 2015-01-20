package confection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GeertJohan/go.rice"
)

type (
	server struct {
		manager       *Manager
		port          int
		staticHandler http.Handler
	}
)

func (s *server) start() {
	s.staticHandler = http.FileServer(rice.MustFindBox("static").HTTPBox())

	portStr := ":" + strconv.Itoa(s.port)

	fmt.Println("Configuration server is available at http://127.0.0.1" + portStr)
	if err := http.ListenAndServe(portStr, s); err != nil {
		panic(err)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Rewrite like a boss
	if req.URL.Path == "/" {
		req.URL.Path = "/config.html"
	}

	// Route like a boss
	switch req.URL.Path {
	case "/config.html", "/app.js", "/bootstrap.min.css":
		s.staticHandler.ServeHTTP(w, req)
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
