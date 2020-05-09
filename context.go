package alexa

//----------------------------------------------------------------------------------------------------------------------
type Context struct {
	memory map[string]interface{}
	Event  *Event
}

func (ctx *Context) SetSession(key string, value interface{}) {
	ctx.memory[key] = value
}

func (ctx *Context) GetFromSession(key string) (interface{}, bool) {
	value, exist := ctx.Event.Session.Attributes[key]
	return value, exist
}
