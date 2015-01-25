# Confection

Confection is a configuration manager plugin for Go projects.

## Configuration

First you need to add tags to fields of the config struct.

* `json` — required for proper serialization
* `title` — human readable field name (optional)
* `attrs` — field attributes: `required`, `readonly`, `ignored`; separated by comma
* `options` — list of supported values, separated by comma

Required attributes will block `manager.RequireConfig()` call until field gets a value

Ignored attributes are not displayed.

Readonly attributes are displayed but disabled.

You also need to pass an unmarshalling function as shown below.

```go
package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/localhots/confection"
)

type (
	Config struct {
		AppName     string `json:"app_name" attrs:"required" title:"Application Name"`
		BuildNumber int    `json:"build_number" attrs:"readonly" title:"Build Number"`
	}
)

func init() {
	confection.SetupFlags()
	flag.Parse()
}

func main() {
	conf := Config{}
	manager := confection.New(conf, func(b []byte) interface{} {
		var newConf Config
		if err := json.Unmarshal(b, &newConf); err != nil {
			panic(err)
		}
		return newConf
	})
	manager.StartServer()
	manager.RequireConfig()

	fmt.Println("Ready to work!")
}
```

[Full example](https://github.com/localhots/confection/blob/master/demo/demo.go)

## Demo Screenshot

<img src="https://raw.githubusercontent.com/localhots/confection/master/demo/demo.png" width="590">