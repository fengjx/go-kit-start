package endpoint

import "sync"

type Inst struct {
	GreeterEndpoints *greeterEndpoints
}

var ins *Inst
var insOnce sync.Once

func GetInst() *Inst {
	insOnce.Do(func() {
		ins = &Inst{
			GreeterEndpoints: newGreeterEndpoints(),
		}
	})
	return ins
}
