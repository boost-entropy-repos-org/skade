package main

import (
	"flag"
	"github.com/Mindslave/skade/internal/engine"
)

func main() {
	engine := engine.NewAnalysisService()
	fileName := flag.String("File", "", "Name of the suspicious file")
	flag.Parse()
	engine.Scan(*fileName)
}
