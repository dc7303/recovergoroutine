package main

import (
	"golang.org/x/tools/go/analysis"

	linters "github.com/dc7303/recovergoroutine"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
