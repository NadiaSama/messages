package messages

type (
	//Message a text temlpate message instance which send to location
	Message interface {
		Name() string
		//Template return tmpl string which used to build template
		//instance. due to go do not support struct statis property
		//is not easy to impl a method which return template instance
		Template() string
	}
)
