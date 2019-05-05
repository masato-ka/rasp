package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"os"
	"rasp/cmd/build"
)

func main() {

	subcommands.Register(subcommands.HelpCommand(), "")
	//subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&build.BuildCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
