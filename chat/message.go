package chat

import "time"

type Message struct {
	Author string    `json:"author"`
	Body   string    `json:"body"`
	Time   time.Time `json:"time"`
}

func (self *Message) String() string {
	return "<" + self.Author + ">" + " " + self.Body
}
