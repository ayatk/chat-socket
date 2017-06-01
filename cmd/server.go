package cmd

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/ayatk/chat-socket/chat"
	"github.com/google/subcommands"
)

type ServerCmd struct {
	port int
}

func (*ServerCmd) Name() string     { return "server" }
func (*ServerCmd) Synopsis() string { return "Start headless server." }
func (*ServerCmd) Usage() string {
	return `server [-port]:
  Start headless server.
`
}

func (s *ServerCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&s.port, "port", 1234, "set server port.")
}

func (s *ServerCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	log.SetFlags(log.Ltime)

	// websocket server
	server := chat.NewServer("/chat")
	go server.Listen()

	// static files
	//http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.port), nil))
	return subcommands.ExitSuccess
}
