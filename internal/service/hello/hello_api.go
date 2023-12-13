package hello

import "context"

func SayHello(ctx context.Context, name string) string {
	return getInst().greetSvc.sayHello(ctx, name)
}
