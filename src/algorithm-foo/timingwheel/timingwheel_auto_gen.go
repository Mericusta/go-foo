package moduleTimingWheel

import sgs "github.com/Mericusta/go-sgs"

var Service *ModuleTimingWheel

func (*ModuleTimingWheel) New(mos ...sgs.ModuleOption) *ModuleTimingWheel {
	mtw := &ModuleTimingWheel{}
	for _, mo := range mos {
		mo(mtw)
	}
	return mtw
}

func (*ModuleTimingWheel) WithTimingWheelRound(twr *timingWheelRound) sgs.ModuleOption {
	return func(m sgs.Module) { m.(*ModuleTimingWheel).rounds = append(m.(*ModuleTimingWheel).rounds, twr) }
}
