package codecs

import (
	"coreapi/primitives"
	"encoding/json"
	"coreapi/document"
	"fmt"
	"log"
)

type typedPrimitive struct {
	Type string `json:"_type"`
}

func (t *typedPrimitive) GetType () string {
	return t.Type
}

type JSONCodec struct {}

func (JSONCodec) Decode(data []byte) (ret primitives.Primitive, e error) {
	var raw interface{}
	e = json.Unmarshal(data, &raw)
	if e != nil {
		panic(fmt.Sprintf("JSON Unmarshal failed: %s", e))
	}
	ret, e = convert(raw)
	return ret, e
}

type FakeError struct {}

func (FakeError) Error () string { return "fake" }

func (JSONCodec) Encode(_ primitives.Primitive) ([]byte, error) {
	return []byte(""), FakeError{}
}

func convert(raw interface{}) (primitives.Primitive, error) {
	defer func () {
		if err := recover(); err != nil {
			log.Println("failed converting:", raw, "downstream error", err)
			panic("died")
		}
	}()
	var ret primitives.Primitive
	// try a map
	as_map, _ := raw.(map[string]interface{})
	if val, ok := as_map["_type"]; ok {
		println(val)
		switch val {
		case "document":
			var title, url string
			fields := make(map[string]primitives.Primitive)
			links := make(map[string]primitives.Link)
			for key, value := range as_map {
				if key != "_type" && key != "_meta" {
					next, err := convert(value)
					println(fmt.Sprintf("got next %s from value %s", next, value))
					if err != nil {
						return nil, err
					}
					if next.GetType() == "link" {
						links[key] = next.(primitives.Link)
					} else {
						fields[key] = next
					}
				}
			}
			if meta, ok := as_map["_meta"].(map[string]string); ok {
				title = meta["title"]
				url = meta["url"]
			} else {
				title = ""
				url = ""
			}
			ret = document.NewDocument(
				title,
				url,
				fields,
				links,
			)
			return ret, nil
		case "link":
			ret = primitives.NewLink()
			return ret, nil
		}
	}
	if slice, ok := raw.([]interface{}); ok {
		converted := make(primitives.Array, 0)
		for _, r := range slice {
			c, e := convert(r)
			if e != nil {
				panic("unhandled error in array conversion")
			}
			log.Printf("converting array %s, tried %s and got %s", slice, r, c)
			converted = append(converted, c)
		}
		return converted, nil
	}
	if num, ok := raw.(primitives.Number); ok {
		return num, nil
	}
	if b, ok := raw.(bool); ok {
		return primitives.Boolean(b), nil
	}
	if s, ok := raw.(string); ok {
		return primitives.String(s), nil
	}
 	panic(fmt.Sprintf("unhandled value: %s", raw))
}