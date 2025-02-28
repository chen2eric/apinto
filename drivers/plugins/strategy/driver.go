package strategy

import (
	"github.com/eolinker/eosc"
	"reflect"
)

type Config struct {
}
type driver struct {
	configType reflect.Type
}

func (d *driver) ConfigType() reflect.Type {
	return d.configType
}

func (d *driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error) {
	return &Strategy{
		id:   id,
		name: name,
	}, nil
}
