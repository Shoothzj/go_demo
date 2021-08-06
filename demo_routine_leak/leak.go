package demo_routine_leak

import "time"

func leak() error {
	go func() {
		time.Sleep(time.Minute)
	}()
	return nil
}
