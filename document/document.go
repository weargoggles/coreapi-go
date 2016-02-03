package document

import (
	"coreapi/primitives"
	"coreapi/transports"
)

type Document struct {
	title string
	url string
	fields map[string]primitives.Primitive
	links map[string]primitives.Link
}

func (*Document) GetType() string {
	return "document"
}

func (d *Document) Get(key string) (primitives.Primitive, bool) {
	ret, ok := d.fields[key]
	return ret, ok
}

func (d *Document) GetTitle() string {
	return d.title
}

func (d *Document) Action(transport transports.Transport, name string, parameters map[string]primitives.Primitive) (primitives.Primitive, error) {
	link, ok := d.links[name]
	if ok {
		return transport.Transition(link, parameters)
	} else {
		return nil, nil
	}
}

func NewDocument(title string, url string, fields map[string]primitives.Primitive, links map[string]primitives.Link) *Document {
	return &Document{
		title: title,
		url: url,
		fields: fields,
		links: links,
	}
}