package main

import (
	"os"
	"path/filepath"

	"github.com/golang/glog"

	"ripple/pkg/cmd"
	"ripple/pkg/cmd/ripple"
)

func main() {
	defer glog.Flush()

	baseName := filepath.Base(os.Args[0])

	err := ripple.NewCommand(baseName).Execute()
	cmd.CheckError(err)
}
