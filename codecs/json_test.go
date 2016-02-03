package codecs

import (
	"testing"
	"coreapi/document"
)

var exampleDocument = []byte(`{
    "_type": "document",
    "_meta": {
        "url": "/",
        "title": "Notes"
    },
    "notes": [
        {
            "_type": "document",
            "_meta": {
                "url": "/1de153fe-6747-41d3-bc0e-d9d7d87e448a",
                "title": "Note"
            },
            "complete": false,
            "description": "Email venue about conference dates",
            "delete": {
                "_type": "link",
                "trans": "delete"
            },
            "edit": {
                "_type": "link",
                "trans": "update",
                "fields": [
                    "description",
                    "complete"
				]
            }
        }
    ],
    "add_note": {
        "_type": "link",
        "trans": "action",
        "fields": [
            {
                "name": "description",
                "required": true
            }
        ]
    }
}`)

func TestDecodeDocument(t *testing.T) {
	doc, err := JSONCodec{}.Decode(exampleDocument)
	if err != nil {
		t.Errorf("error decoding test document: %s", err)
		return
	}
	if doc == nil {
		t.Errorf("error, document was nil")
		return
	}
	if doc.GetType() != "document" {
		t.Errorf("didn't get document, got %s instead from %s", doc.GetType(), doc)
	}
	document := doc.(*document.Document)
	descdocument.Get("description")
	if document.GetTitle() != "Notes" {
		t.Fatalf("Wrong title, got %s instead of Notes", document.GetTitle())
	}
}