package migration

type Option interface {
	apply(*MigrationOption)
}

type MigrationOption struct {
	ignore bool
}

func (o *MigrationOption) Ignore() bool {
	return o.ignore
}

func NewMigration(opts ...Option) *MigrationOption {
	o := &MigrationOption{}
	for _, opt := range opts {
		opt.apply(o)
	}
	return o
}

type migrationOptionFn func(*MigrationOption)

func (fn migrationOptionFn) apply(opt *MigrationOption) {
	fn(opt)
}

func WithIgnore() Option {
	return migrationOptionFn(func(mo *MigrationOption) {
		mo.ignore = true
	})
}
