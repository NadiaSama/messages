package messages

type (
	Property map[string]interface{}
	//Context define common property which shared by messages. context
	//can derive from a context
	Context struct {
		property  Property
		locations []Location
	}
)

//Deriv create a new context.
func (c *Context) Deriv(props Property, locs ...Location) *Context {
	ret := &Context{
		locations: make([]Location, len(c.locations)),
		property:  make(Property, len(c.property)),
	}
	copy(ret.locations, c.locations)
	for k, v := range c.property {
		ret.property[k] = v
	}

	for k, v := range props {
		ret.property[k] = v
	}

	ret.locations = append(ret.locations, locs...)
	return ret
}

func (c *Context) Send(title string, prop Property) error {
	return nil
}

func format(base Property) ([]string, error) {
	ret := []string{}
	for k, v := range base {

	}
}
