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
		Path       string      `json:"path"`
		Type       string      `json:"type"`
		Value      interface{} `json:"value"`
		IsRequired bool        `json:"is_required"`
		IsReadonly bool        `json:"is_readonly"`
		IsIgnored  bool        `json:"is_ignored"`
		Title      string      `json:"title"`
		Options    []string    `json:"options"`
	}
)

const (
	tJson     = "json"
	tTitle    = "title"
	tAttrs    = "attrs"
	tOptions  = "options"
	aRequired = "required"
	aReadonly = "readonly"
	aIgnored  = "ignored"
	sep       = ","
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
	out.WriteByte('\n')

	return out.Bytes(), nil
}

// TODO: function is too heavy, needs refactor
func (c *config) meta(prefix string) []*configField {
	var (
		fields = []*configField{}
		cval   = reflect.ValueOf(c.config)
		ctyp   = reflect.TypeOf(c.config)
		ckind  = cval.Kind()
	)

	if ckind != reflect.Struct {
		panic(fmt.Errorf("Config is expected to be a Struct, not %s", ckind.String()))
	}

	for i := 0; i < cval.NumField(); i++ {
		var (
			field = ctyp.Field(i)
			val   = cval.Field(i)
			kind  = val.Kind()

			jsonKey = field.Tag.Get(tJson)
			path    = strings.Join([]string{prefix, jsonKey}, "/")
			title   = field.Tag.Get(tTitle)
			attrs   = field.Tag.Get(tAttrs)
			options = field.Tag.Get(tOptions)

			cf = &configField{
				Path:  path,
				Type:  val.Kind().String(),
				Title: title,
			}
		)

		if len(attrs) > 0 {
			for _, attr := range strings.Split(attrs, sep) {
				if attr == aRequired {
					cf.IsRequired = true
				}
				if attr == aReadonly {
					cf.IsReadonly = true
				}
				if attr == aIgnored {
					cf.IsIgnored = true
				}
			}

		}

		// Substitute field name for title if none set
		if kind != reflect.Struct {
			cf.Value = val.Interface()
		}
		if title == "" {
			cf.Title = field.Name
		}
		if len(options) > 0 {
			cf.Options = strings.Split(options, sep)
		}

		fields = append(fields, cf)

		// Recursion here
		if kind == reflect.Struct {
			subconf := &config{
				config: val.Interface(),
			}
			fields = append(fields, subconf.meta(path)...)
		}
	}

	return fields
}
