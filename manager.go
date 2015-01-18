package confection

import (
	"flag"
	"sync"
)

type (
	Manager struct {
		mux  *sync.Mutex
		conf *config
		file *configFile
	}
)

var (
	configPath string
	serverPort int
)

func Setup() {
	flag.StringVar(&configPath, "config", "config.json", "Path to config file")
	flag.IntVar(&serverPort, "config-port", 5050, "Config manager http port")
}

func New(conf interface{}) *Manager {
	mgr := &Manager{
		mux: &sync.Mutex{},
		conf: &config{
			config: conf,
		},
		file: &configFile{
			path: configPath,
		},
	}

	return mgr
}

func (m *Manager) StartServer() {
	srv := &server{
		manager: m,
		port:    serverPort,
	}
	srv.start()
}
