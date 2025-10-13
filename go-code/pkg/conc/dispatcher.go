package conc

// To do
type Message struct {
	Content string
}

type Receiver struct {
	Ch chan Message
}
