package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/ayatk/chat-socket/chat"
	"github.com/google/subcommands"
	"github.com/marcusolsson/tui-go"
	"golang.org/x/net/websocket"
)

type ClientCmd struct {
	host string
	port int
	name string
}

func (*ClientCmd) Name() string     { return "client" }
func (*ClientCmd) Synopsis() string { return "Start client." }
func (*ClientCmd) Usage() string {
	return `client -host <hostname> [-port <port num> default: 1234] -name <username>:
  Start client.
`
}

func (s *ClientCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.host, "host", "", "Connection server host name.")
	f.IntVar(&s.port, "port", 1234, "set client port.")
	f.StringVar(&s.name, "name", "", "set your name.")

}

func (s *ClientCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	ws, err := websocket.Dial(
		"ws://"+s.host+":"+strconv.Itoa(s.port)+"/chat", "", "http://"+s.host+":"+strconv.Itoa(s.port))
	if err != nil {
		panic(err.Error())
	}

	history := tui.NewVBox()
	history.SetBorder(true)
	history.Append(tui.NewSpacer())

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chatBox := tui.NewVBox(history, inputBox)
	chatBox.SetSizePolicy(tui.Expanding, tui.Expanding)

	go func() {
		for {
			var msg = make([]byte, 1024)
			if n, err := ws.Read(msg); err != nil {
				panic("Read: " + err.Error())
			} else {
				resp := &chat.Message{}
				json.Unmarshal(msg[0:n], resp)

				history.Append(tui.NewHBox(
					tui.NewLabel(resp.Time.Format("2006-01-02 15:04:05")),
					tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", resp.Author))),
					tui.NewLabel(resp.Body),
					tui.NewSpacer(),
				))
			}
		}
	}()

	input.OnSubmit(func(e *tui.Entry) {
		if len(e.Text()) == 0 {
			return
		}
		message := &chat.Message{
			Author: s.name,
			Body:   e.Text(),
			Time:   time.Now(),
		}

		jsonMassage, _ := json.Marshal(*message)

		if _, err := ws.Write(jsonMassage); err != nil {
			panic("Write: " + err.Error())
		}
		input.SetText("")
	})

	root := tui.NewHBox(chatBox)

	ui := tui.New(root)
	ui.SetKeybinding(tui.KeyEsc, func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}

	return subcommands.ExitSuccess
}
