package tx

type OptionFunc interface {
	apply(*Options)
}

func OptionApply(o *Options, opts ...OptionFunc) {
	for _, opt := range opts {
		opt.apply(o)
	}
}

type optionFn func(*Options)

func (fn optionFn) apply(o *Options) {
	fn(o)
}

type Options struct {
}
