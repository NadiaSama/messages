package messages

type (
	//Message a text message instance which send to location
	Message interface {
		ID() string
		Name() string
	}
)
