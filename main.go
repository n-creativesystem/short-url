package main

import (
	"github.com/n-creativesystem/short-url/pkg/cmd"
	_ "github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/driver"
)

func main() {
	cmd.Execute()
}
