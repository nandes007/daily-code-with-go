package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// context.WithValue()
// Signature: withValue(parent Context, key, val interface{}) (ctx Context)
func contextWithValue() {
	fmt.Println("\nContext With Value")
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", injectMsgID(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}

// context.WithCacel()
// Signature: func WithCancal(parent Context) (ctx Context, cancel CancelFunc)
func contextWithCancel() {
	fmt.Println("Context With Cancel")
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)
}

// context.WithTimeout()
// Used for time-based cancellation. The signature of the function is:
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// context.WithTimeout function will
// - Will return a copy of the parentContext with the new done channel.
// - Accept a timeout duration after which this done channel will be closed and context will be cancelled.
// - A cancel function which can be called in case the context needs to be cancelled before timeout.
func contextWithTimeout() {
	fmt.Println("\nContext with timeout")
	ctx := context.Background()
	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	go task1(cancelCtx)
	time.Sleep(time.Second * 4)
}

// context.WithDeadline()
// Used for deadline-based cancellation. The signature of the function is
// func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
// context.WithDeadline function
// - Will return a copy of the parentContext with the new done channel.
// - Accept a deadline after which this done channel will be closed and context will be canceled
// - A cancel function which can be called in case the context needs to be canceled before the deadline is reached
func contextWithDeadline() {
	fmt.Println("\nContext With Deadline")
	ctx := context.Background()
	cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
	defer cancel()
	go task(cancelCtx)
	time.Sleep(time.Second * 6)
}
