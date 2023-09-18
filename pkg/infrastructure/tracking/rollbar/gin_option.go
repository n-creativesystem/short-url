package rollbar

type ginOption struct {
	onlyCrashes bool
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
		onlyCrashes: false,
	}
}

func WithOnlyCrashers(onlyCrashes bool) GinOption {
	return ginOptionFn(func(opt *ginOption) {
		opt.onlyCrashes = onlyCrashes
	})
}
