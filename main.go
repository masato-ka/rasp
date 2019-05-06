package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
	"rasp/cmd/build"
)

func main() {

	version := "0.1.3"

	subcommands.Register(subcommands.HelpCommand(), "")
	//subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&build.BuildCmd{}, "")

	v := flag.Bool("version", false, "version")

	flag.Parse()
	if *v {
		fmt.Fprintf(os.Stdout, "rasp %s \n", version)
		os.Exit(0)
	}

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
