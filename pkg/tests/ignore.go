package tests

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

type ignoreUnexportedFieldsMatcher[T any] struct {
	x            T
	ignoreFields []cmp.Option
}

func NewIgnoreUnexportedFieldsMatcher[T any](x T, fields ...string) gomock.Matcher {
	mp := map[string]struct{}{}
	for _, field := range fields {
		mp[field] = struct{}{}
	}
	ignoreFields := cmp.FilterPath(func(p cmp.Path) bool {
		sf, ok := p.Index(-1).(cmp.StructField)
		if !ok {
			return false
		}
		_, ok = mp[sf.Name()]
		return ok
	}, cmp.Ignore())
	var v T
	return &ignoreUnexportedFieldsMatcher[T]{
		x: x,
		ignoreFields: []cmp.Option{
			cmp.AllowUnexported(v),
			ignoreFields,
		},
	}
}

func (m *ignoreUnexportedFieldsMatcher[T]) Matches(x interface{}) bool {
	if diff := cmp.Diff(m.x, x, m.ignoreFields...); diff != "" {
		return false
	}
	return true
}

func (m *ignoreUnexportedFieldsMatcher[T]) String() string {
	return fmt.Sprintf("is equal to %v", m.x)
}
