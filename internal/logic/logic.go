package logic

import "sync"

type Inst struct {
	HelloLogic *helloLogic
}

var ins *Inst
var insOnce sync.Once

func GetInst() *Inst {
	insOnce.Do(func() {
		ins = &Inst{
			HelloLogic: newHelloLogic(),
		}
	})
	return ins
}
