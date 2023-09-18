package sentry

import "time"

type ginOption struct {
	rePanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type GinOption interface {
	apply(opt *ginOption)
}

type ginOptionFn func(opt *ginOption)

func (fn ginOptionFn) apply(opt *ginOption) {
	fn(opt)
}

func defaultGinOption() ginOption {
	return ginOption{
		rePanic: true,
	}
}

func WithRePanic(rePanic bool) GinOption {
	return ginOptionFn(func(opt *ginOption) {
		opt.rePanic = rePanic
	})
}

func WithWaitForDelivery(waitForDelivery bool) GinOption {
	return ginOptionFn(func(opt *ginOption) {
		opt.waitForDelivery = waitForDelivery
	})
}

func WithTimeout(t time.Duration) GinOption {
	return ginOptionFn(func(opt *ginOption) {
		opt.timeout = t
	})
}
