package main

import (
	linters "github.com/dc7303/recovergoroutine/recovergoroutine"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
