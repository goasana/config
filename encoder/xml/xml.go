package xml

import (
	"encoding/xml"

	"github.com/goasana/config/encoder"
)

func init() {
	e := NewEncoder()
	encoder.Register(e.String(), e)
}

const XML encoder.Provider =  "xml"

type xmlEncoder struct{}

func Encode(v interface{}, hasIndent bool) ([]byte, error) {
	if hasIndent {
		return xml.MarshalIndent(v, "", " ")
	}
	return xml.Marshal(v)
}

func (j xmlEncoder) Encode(v interface{}, hasIndent ...bool) ([]byte, error) {
	return Encode(v, len(hasIndent) > 0 && hasIndent[0])
}

func Decode(d []byte, v interface{}) error {
	return xml.Unmarshal(d, v)
}

func (j xmlEncoder) Decode(d []byte, v interface{}) error {
	return Decode(d, v)
}

func (j xmlEncoder) String() encoder.Provider {
	return XML
}

func NewEncoder() encoder.Encoder {
	return xmlEncoder{}
}
