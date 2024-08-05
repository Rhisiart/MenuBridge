package protocol

type Hub struct {
	Clients       chan Command
	Connection    chan *Client
	Disconnection chan *Client
}
