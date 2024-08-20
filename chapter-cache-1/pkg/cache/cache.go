package main

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		isSuccess := true
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <-ctx.Done():
		// 其他RPC调用调用失败
		return ctx.Err()
	case e := <-err:
		// 本RPC调用失败，返回错误信息
		return e
	case <-result:
		// 本RPC调用成功，不返回错误信息
		return nil
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	// RPC1调用
	err := Rpc(ctx, "http://rpc_1_url")
	if err != nil {
		return
	}

	// RPC2调用
	go func() {
		err := Rpc(ctx, "http://rpc_2_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC3调用
	go func() {
		err := Rpc(ctx, "http://rpc_3_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC4调用
	go func() {
		err := Rpc(ctx, "http://rpc_4_url")
		if err != nil {
			cancel()
		}
	}()

	time.Sleep(5 * time.Second)
}
