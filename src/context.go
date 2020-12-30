package main

import (
	"context"
	"fmt"

	"sample.com/external"
)

func contexttest() {
	fmt.Println("--------------------context--------------------")
	ctx := context.Background()
	valueCtx := external.WithRequestID(ctx, 123)
	reqID, ok := external.GetRequestID(valueCtx)
	if !ok {
		fmt.Println("リクエストIDなし")
		return
	}
	fmt.Printf("requestID:%v\n", reqID)

	errReqID, ok := external.GetRequestID(ctx)
	if !ok {
		fmt.Printf("requestID:%v\n", errReqID)
		fmt.Println("リクエストIDなし")
		return
	}
}
