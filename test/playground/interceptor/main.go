package main

import (
	"context"
	"fmt"
)

type handler func(ctx context.Context)
type invoker func(ctx context.Context, interceptors []interceptor, h handler) error
type interceptor func(ctx context.Context, h handler, ivk invoker) error

func main() {
	var ctx context.Context
	var ceps []interceptor
	var h = func(ctx context.Context) {
		fmt.Println("do something handler")
	}

	var inter1 = func(ctx context.Context, h handler, ivk invoker) error {
		//h(ctx)
		fmt.Println("interceptor 1")
		return ivk(ctx, ceps, h)
	}

	var inter2 = func(ctx context.Context, h handler, ivk invoker) error {
		//h(ctx)
		fmt.Println("interceptor 2")
		return ivk(ctx, ceps, h)
	}

	var inter3 = func(ctx context.Context, h handler, ivk invoker) error {
		//h(ctx)
		fmt.Println("interceptor 3")
		return ivk(ctx, ceps, h)
	}

	ceps = append(ceps, inter1, inter2, inter3)

	var ivk = func(ctx context.Context, interceptors []interceptor, h handler) error {
		fmt.Println("ivk handler start")
		return nil
	}
	cep := getChainInterceptor(ctx, ceps, ivk)
	_ = cep(ctx, h, ivk)
}

// 链式调用
// interceptor1 -> interceptor2-> interceptor3 -> ...... -> handler
// interceptor
func getChainInterceptor(ctx context.Context, interceptors []interceptor, ivk invoker) interceptor {
	if len(interceptors) == 0 {
		return nil
	}
	if len(interceptors) == 1 {
		return interceptors[0]
	}
	return func(ctx context.Context, h handler, ivk invoker) error {
		fmt.Println("getChainInterceptor")
		return interceptors[0](ctx, h, getInvoker(ctx, interceptors, 0, ivk))
	}
}

// get Last Invoker
func getInvoker(ctx context.Context, interceptors []interceptor, cur int, ivk invoker) invoker {
	fmt.Println("getInvoker in cur", cur)
	// last time invoke will return real ivk
	if cur == len(interceptors)-1 {
		return ivk
	}
	return func(ctx context.Context, interceptors []interceptor, h handler) error {
		return interceptors[cur+1](ctx, h, getInvoker(ctx, interceptors, cur+1, ivk))
	}
}
