package hello

import "sync"

type inst struct {
	greetSvc *greetService
}

var ins *inst
var insOnce sync.Once

func getInst() *inst {
	insOnce.Do(func() {
		ins = &inst{
			greetSvc: newGreetService(),
		}
	})
	return ins
}
