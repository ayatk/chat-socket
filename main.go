package main

import (
	"context"
	"flag"
	"os"

	"github.com/ayatk/chat-socket/cmd"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&cmd.ServerCmd{}, "")
	subcommands.Register(&cmd.ClientCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
