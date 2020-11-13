package messages

type (
	//Location is a end point which used to receive message
	Location interface {
		Send(id string, name string, data string) error
	}
)
