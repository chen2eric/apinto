package strategy

import (
	"github.com/eolinker/apinto/strategy"
	"github.com/eolinker/eosc"
	eoscContext "github.com/eolinker/eosc/eocontext"
)

type Strategy struct {
	id   string
	name string
}

func (s *Strategy) DoFilter(ctx eoscContext.EoContext, next eoscContext.IChain) (err error) {
	return strategy.Strategy(ctx, next)
}

func (s *Strategy) Destroy() {
	return
}

func (s *Strategy) Id() string {
	return s.id
}

func (s *Strategy) Start() error {
	return nil
}

func (s *Strategy) Reset(conf interface{}, workers map[eosc.RequireId]eosc.IWorker) error {
	return nil
}

func (s *Strategy) Stop() error {
	return nil
}

func (s *Strategy) CheckSkill(skill string) bool {
	return eoscContext.FilterSkillName == skill
}
