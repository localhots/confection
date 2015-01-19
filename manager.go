package confection

import (
	"flag"
	"sync"
)

type (
	Manager struct {
		mux          *sync.Mutex
		conf         *config
		file         *configFile
		unmarshaller func([]byte) interface{}
	}
)

var (
	configPath string
	serverPort int
)

func SetupFlags() {
	flag.StringVar(&configPath, "config", "config.json", "Path to config file")
	flag.IntVar(&serverPort, "config-port", 5050, "Config manager http port")
}

func New(conf interface{}, unmarshaller func([]byte) interface{}) *Manager {
	mgr := &Manager{
		mux: &sync.Mutex{},
		conf: &config{
			config: conf,
		},
		file: &configFile{
			path: configPath,
		},
		unmarshaller: unmarshaller,
	}
	mgr.bootstrap()

	return mgr
}

func (m *Manager) Config() interface{} {
	return m.conf.config
}

func (m *Manager) StartServer() {
	srv := &server{
		manager: m,
		port:    serverPort,
	}
	srv.start()
}

func (m *Manager) bootstrap() {
	if m.file.isExist() {
		b, err := m.file.read()
		if err != nil {
			panic(err)
		}

		m.importJson(b)
	} else {
		b, err := m.conf.dump()
		if err != nil {
			panic(err)
		}

		err = m.file.write(b)
		if err != nil {
			panic(err)
		}
	}
}

func (m *Manager) importJson(b []byte) {
	m.conf.config = m.unmarshaller(b)
}
