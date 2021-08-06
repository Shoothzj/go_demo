package timeout

import (
	"context"
	"time"
)

func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
}

func requestWorkSync(ctx context.Context, job interface{}) error {
	return hardWork(job)
}

func requestWorkSyncV1(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func requestWorkSyncV2(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	done := make(chan error, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case p := <-panicChan:
		panic(p)
	case <-ctx.Done():
		return ctx.Err()
	}
}
