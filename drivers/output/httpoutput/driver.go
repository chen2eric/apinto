package httpoutput

import (
	"reflect"

	"github.com/eolinker/eosc"
)

type Driver struct {
	configType reflect.Type
}

func (d *Driver) ConfigType() reflect.Type {
	return d.configType
}

func Check(v interface{}) (*Config, error) {
	conf, ok := v.(*Config)
	if !ok {
		return nil, errConfigType
	}

	httpConf := conf
	if httpConf.Method == "" {
		return nil, errMethod
	}
	switch httpConf.Method {
	case "GET", "POST", "HEAD", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE":
	default:
		return nil, errMethod
	}

	if httpConf.Url == "" {
		return nil, errUrlNull
	}

	if httpConf.Type == "" {
		httpConf.Type = "line"
	}

	switch httpConf.Type {
	case "line", "json":
	default:
		return nil, errFormatterType
	}

	if len(httpConf.Formatter) == 0 {
		return nil, errFormatterConf
	}

	return httpConf, nil
}

func (d *Driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error) {
	conf, err := Check(v)
	if err != nil {
		return nil, err
	}
	worker := &HttpOutput{
		id:     id,
		config: conf,
	}

	return worker, nil
}
