package transports

import "coreapi/primitives"

type Transport interface {
	Transition(primitives.Link, map[string]primitives.Primitive) (primitives.Primitive, error)
}