package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/dc7303/recovergoroutine/recovergoroutine"
)

func main() {
	singlechecker.Main(recovergoroutine.Analyzer)
}
