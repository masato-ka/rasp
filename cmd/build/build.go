package build

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
	"rasp/lib"
	"strings"
)

type BuildCmd struct {
	repo string
	cmd  string
	dst  string
	skip bool
}

func (*BuildCmd) Name() string     { return "build" }
func (*BuildCmd) Synopsis() string { return "Print args to stdout." }
func (*BuildCmd) Usage() string {
	return `build [-repo] <path text> [-cmd] <cmd text>:
get the code from repo and run the cmd for build.
`
}

func (b *BuildCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&b.repo, "repo", "", "repository path")
	f.StringVar(&b.cmd, "cmd", "", "command")
	f.StringVar(&b.dst, "dst", "./", "distination")
	f.BoolVar(&b.skip, "skip-build", false, "Skip build flag")
}

func (b *BuildCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if b.repo == "" {
		fmt.Fprintf(os.Stderr, "Usage error. \n")
		return subcommands.ExitUsageError
	}

	p, err := lib.Getrepo(b.repo, b.dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed fetch resource from repo : %+v \n", b.repo)
		return subcommands.ExitFailure
	}

	if !b.skip {
		_, buildErr := b.build(p)
		if buildErr != nil {
			fmt.Fprintf(os.Stderr, "Failed build: %+v \n", err)
			return subcommands.ExitFailure
		}
		fmt.Fprintf(os.Stdout, "complete build.\n")
	}

	return subcommands.ExitSuccess
}

func (b *BuildCmd) build(p string) (string, error) {

	d := lib.Detect(p)

	cmd := ""
	var args []string

	if b.cmd != "" {
		s := strings.Split(b.cmd, " ")
		cmd = s[0]
		args = s[1:]
	}

	if d == "python" {
		cmd = "pip"
		args = append(args, "install", "-e", ".")
	}

	if d == "go" {
		cmd = "go"
		args = append(args, "build")
	}

	if d != "" || b.cmd != "" {
		ps, err := lib.Execution(cmd, args, p, os.Stdout, os.Stderr)
		if err != nil {
			return "", err
		}
		return ps.String(), nil
	}

	return "pass", nil

}
