package access_log

import (
	"fmt"
	"reflect"

	"github.com/eolinker/apinto/output"

	"github.com/eolinker/eosc"
)

type Driver struct {
	profession string
	name       string
	label      string
	desc       string
	workers    eosc.IWorkers
	configType reflect.Type
}

func (d *Driver) Check(v interface{}, workers map[eosc.RequireId]eosc.IWorker) error {
	_, err := d.check(v)
	if err != nil {
		return err
	}
	return nil
}

func (d *Driver) check(v interface{}) (*Config, error) {
	conf, ok := v.(*Config)
	if !ok {
		return nil, eosc.ErrorConfigFieldUnknown
	}
	return conf, nil
}

func (d *Driver) ConfigType() reflect.Type {
	return d.configType
}

func (d *Driver) getList(auths []eosc.RequireId) ([]output.IEntryOutput, error) {
	ls := make([]output.IEntryOutput, 0, len(auths))
	for _, id := range auths {
		worker, has := d.workers.Get(string(id))
		if !has {
			return nil, fmt.Errorf("%s:%w", id, eosc.ErrorWorkerNotExits)
		}

		outPut, ok := worker.(output.IEntryOutput)
		if !ok {
			return nil, fmt.Errorf("%s:worker not implement IEntryOutput", string(id))
		}

		ls = append(ls, outPut)

	}
	return ls, nil
}

func (d *Driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error) {
	conf, err := d.check(v)
	if err != nil {
		return nil, err
	}
	list, err := d.getList(conf.Output)
	if err != nil {
		return nil, err
	}
	o := &accessLog{
		Driver: d,
		id:     id,
		output: list,
	}

	return o, nil
}
