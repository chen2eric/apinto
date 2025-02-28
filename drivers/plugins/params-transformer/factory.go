package params_transformer

import (
	"github.com/eolinker/eosc/log"
	"github.com/eolinker/eosc/utils/schema"
	"reflect"

	"github.com/eolinker/eosc"
)

const (
	Name = "params_transformer"
)

func Register(register eosc.IExtenderDriverRegister) {
	log.Debug("register params_transformer is ", Name)
	register.RegisterExtenderDriver(Name, NewFactory())
}

type Factory struct {
}

func (f *Factory) Render() interface{} {
	render, err := schema.Generate(reflect.TypeOf((*Config)(nil)), nil)
	if err != nil {
		return nil
	}
	return render
}
func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Create(profession string, name string, label string, desc string, params map[string]interface{}) (eosc.IExtenderDriver, error) {
	d := &Driver{
		profession: profession,
		name:       name,
		label:      label,
		desc:       desc,
		configType: reflect.TypeOf((*Config)(nil)),
	}

	return d, nil
}
