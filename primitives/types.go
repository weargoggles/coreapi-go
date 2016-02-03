package primitives

type Link struct {}

func (Link) GetType() string {
	return "link"
}

func NewLink () Link {
	return Link{}
}

type Error struct {
	message string
}

func (Error) GetType() string {
	return "error"
}

func (e Error) Error() string {
	return e.message
}

type Number float64

func (Number) GetType() string  {
	return "number"
}

type String string

func (String) GetType() string {
	return "string"
}

type Boolean bool

func (Boolean) GetType() string {
	return "boolean"
}

type Array []Primitive

func (Array) GetType() string {
	return "array"
}