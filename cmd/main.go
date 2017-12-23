package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/drish/ben"
	"github.com/drish/ben/config"

	"github.com/drish/ben/utils"
)

var usage = `Usage: ben [options...]
Options:
  -o  output file. Default is ./ben-summary.html.
  -d  display benchmark results to stdout. Default is true
`

var defaultBenchmarkFile = "./ben-summary.html"

// TODO: add before environment
func main() {
	trap()

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	outputFlag := flag.String("o", defaultBenchmarkFile, "output summary file")
	displayFlag := flag.Bool("d", false, "display benchmark results to stdout")
	flag.Parse()

	fmt.Println(*displayFlag)

	c, err := config.ReadConfig("ben.json")
	if err != nil {
		utils.Fatal(err)
	}

	err = ben.New(c).Run(*outputFlag, *displayFlag)

	if err != nil {
		utils.Fatal(err)
	}
}

// trappy
func trap() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		println("\n")
		os.Exit(1)
	}()
}
