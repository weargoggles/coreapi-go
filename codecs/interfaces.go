package codecs

import "coreapi/primitives"

type Codec interface {
	Decode([]byte) (primitives.Primitive, error)
	Encode(primitives.Primitive) ([]byte, error)
}