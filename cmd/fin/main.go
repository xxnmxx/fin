package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/xxnmxx/fin"
)

func main() {
	run()
}

func run() {
	compCmd := flag.NewFlagSet("comp", flag.ExitOnError)
	compEach := compCmd.Bool("e", false, "if true print each duration's value")
	args := os.Args
	if len(args) < 2 {
		fmt.Println("must have subcommands")
		os.Exit(1)
	}
	switch args[1] {
	case "comp":
		compCmd.Parse(args[2:])
		if len(args[2:]) != 3 {
			fmt.Println("params: principal float64, rate float64, duration int")
			os.Exit(1)
		}
		p, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("params: principal float64, rate float64, duration int")
			os.Exit(1)
		}
		r, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			fmt.Println("params: principal float64, rate float64, duration int")
			os.Exit(1)
		}
		d, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println("params: principal float64, rate float64, duration int")
			os.Exit(1)
		}
		if *compEach {
			fin.CompoundEach(p, r, d)
			break
		}
		ret := fin.Compound(p, r, d)
		fmt.Println(ret)
	case "test":
		// implements help flag with usage
	default:
		fmt.Println("must have subcommands")

	}
}
