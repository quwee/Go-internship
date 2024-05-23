package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	RunBakery(parseFlags())
}

func parseFlags() (n, m, k, t1, t2 int) {
	nArg := flag.Int("n", -1, "Number of goroutines of bake step")
	mArg := flag.Int("m", -1, "Number of goroutines of pack step")
	kArg := flag.Int("k", -1, "Number of cakes")
	t1Arg := flag.Int("t1", -1, "t1 parameter")
	t2Arg := flag.Int("t2", -1, "t2 parameter")

	flag.Parse()

	switch {
	case *nArg == -1:
		printErrorAndExit("n is not defined")
	case *nArg <= 0:
		printErrorAndExit("n must be > 0")
	case *mArg == -1:
		printErrorAndExit("m is not defined")
	case *mArg <= 0:
		printErrorAndExit("m must be > 0")
	case *kArg == -1:
		printErrorAndExit("k is not defined")
	case *kArg <= 0:
		printErrorAndExit("k must be > 0")
	case *t1Arg == -1:
		printErrorAndExit("t1 is not defined")
	case *t1Arg <= 0:
		printErrorAndExit("t1 must be > 0")
	case *t2Arg == -1:
		printErrorAndExit("t2 is not defined")
	case *t2Arg <= 0:
		printErrorAndExit("t2 must be > 0")
	case *t1Arg > *t2Arg:
		printErrorAndExit("t1 must be <= t2")
	}

	return *nArg, *mArg, *kArg, *t1Arg, *t2Arg
}

func printErrorAndExit(message string) {
	fmt.Println(message)
	fmt.Println("Usage of program:")
	flag.PrintDefaults()
	os.Exit(1)
}
