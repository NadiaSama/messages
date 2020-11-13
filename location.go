package messages

type (
	//Location is a end point which used to receive message
	Location interface {
		Send(msg Message, data string) error
	}
)
