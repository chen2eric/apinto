package nsq

import (
	"github.com/eolinker/apinto/output"
	"github.com/eolinker/eosc"
	"reflect"
)

var _ output.IEntryOutput = (*NsqOutput)(nil)
var _ eosc.IWorker = (*NsqOutput)(nil)

type NsqOutput struct {
	id        string
	write     *Writer
	config    *Config
	isRunning bool
}

func (n *NsqOutput) Output(entry eosc.IEntry) error {
	w := n.write
	if w != nil {
		return w.output(entry)
	}
	return eosc.ErrorWorkerNotRunning
}

func (n *NsqOutput) Reset(conf interface{}, workers map[eosc.RequireId]eosc.IWorker) error {
	cfg, err := Check(conf)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(cfg, n.config) {
		return nil
	}
	n.config = cfg
	if n.isRunning {
		w := n.write
		if w == nil {
			w = NewWriter()
		}

		err = w.reset(cfg)
		if err != nil {
			return err
		}
		n.write = w
	}

	return nil
}

func (n *NsqOutput) Stop() error {
	w := n.write
	if w != nil {
		return w.stop()
	}
	return nil
}

func (n *NsqOutput) Id() string {
	return n.id
}

func (n *NsqOutput) Start() error {
	n.isRunning = true
	w := n.write
	if w == nil {
		w = NewWriter()
	}
	err := w.reset(n.config)
	if err != nil {
		return err
	}
	n.write = w

	return nil
}

func (n *NsqOutput) CheckSkill(skill string) bool {
	return output.CheckSkill(skill)
}
