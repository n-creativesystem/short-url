package flags

import "github.com/spf13/pflag"

func MustGetInt(name string, flag *pflag.FlagSet) int {
	v, err := flag.GetInt(name)
	if err != nil {
		panic(err)
	}
	return v
}
