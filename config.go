package confection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type (
	config struct {
		config interface{}
	}
	configField struct {
		Path        string   `json:"path"`
		IsRequired  bool     `json:"is_required"`
		IsReadonly  bool     `json:"is_readonly"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Options     []string `json:"options"`
	}
)

const (
	tJson        = "json"
	tTitle       = "title"
	tDescription = "description"
	tAttrs       = "attrs"
	tOptions     = "options"
	aRequired    = "required"
	aReadonly    = "readonly"
	sep          = ","
)

func (c *config) dump() ([]byte, error) {
	var (
		out bytes.Buffer
		b   []byte
		err error
	)

	if b, err = json.Marshal(c.config); err != nil {
		return nil, err
	}
	// Indent with empty prefix and four spaces
	if err = json.Indent(&out, b, "", "    "); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c *config) load(b []byte) (err error) {
	return
}

// TODO: function draft, needs refactor
func (c *config) meta(prefix string) []*configField {
	var (
		fields = []*configField{}
		cval   = reflect.ValueOf(c.config)
		typ    = reflect.TypeOf(c.config)
		kind   = cval.Kind()
	)

	if kind != reflect.Struct {
		panic(fmt.Errorf("Config is expected to be a Struct, not %s", kind.String()))
	}

	for i := 0; i < cval.NumField(); i++ {
		var (
			field = typ.Field(i)
			val   = cval.Field(i)

			jsonKey     = field.Tag.Get(tJson)
			path        = strings.Join([]string{prefix, jsonKey}, "/")
			title       = field.Tag.Get(tTitle)
			description = field.Tag.Get(tDescription)
			attrs       = strings.Split(field.Tag.Get(tAttrs), sep)
			options     = strings.Split(field.Tag.Get(tOptions), sep)

			cf = &configField{
				Path:        path,
				Title:       title,
				Description: description,
			}
		)

		// Skip field if no tags are set
		if title == "" && len(attrs) == 0 && len(options) == 0 {
			continue
		}

		// Substitute field name for title if none set
		if title == "" {
			cf.Title = field.Name
		}

		for _, attr := range attrs {
			if attr == aRequired {
				cf.IsRequired = true
			}
			if attr == aReadonly {
				cf.IsReadonly = true
			}
		}

		fields = append(fields, cf)

		// Recursion here
		if val.Kind() == reflect.Struct {
			subconf := &config{
				config: val.Interface(),
			}
			fields = append(fields, subconf.meta(path)...)
		}
	}

	return fields
}
