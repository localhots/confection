package confection

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type (
	configFile struct {
		path string
	}
)

func (cf *configFile) write(b []byte) (err error) {
	var (
		fd *os.File
		n  int // Bytes written successfully
	)

	if cf.isExist() {
		fd, err = os.OpenFile(cf.path, os.O_TRUNC|os.O_WRONLY, 0633)
	} else {
		if err = cf.mkdirp(); err != nil {
			return
		}

		fd, err = os.Create(cf.path)
	}
	if err != nil {
		return
	}
	defer fd.Close()

	n, err = fd.Write(b)
	if err == nil && n != len(b) {
		return fmt.Errorf("Failed to write config file: written %d/%d bytes", n, len(b))
	}

	return
}

func (cf *configFile) read() ([]byte, error) {
	if cf.isExist() {
		return ioutil.ReadFile(cf.path)
	} else {
		return nil, fmt.Errorf("Config file does not exist")
	}
}

func (cf *configFile) isExist() bool {
	_, err := os.Stat(cf.path)
	return (err == nil)
}

func (cf *configFile) mkdirp() error {
	dir := path.Dir(cf.path)
	return os.MkdirAll(dir, 0755)
}
