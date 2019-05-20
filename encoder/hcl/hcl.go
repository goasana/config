package hcl

import (
	"github.com/goasana/config/encoder"
	"github.com/goasana/config/encoder/json"
	"github.com/hashicorp/hcl"
)

func init()  {
	e := NewEncoder()
	encoder.Register(e.String(), e)
}

type hclEncoder struct{}

func Encode(v interface{}, hasIndent bool) ([]byte, error) {
	return json.Encode(v, hasIndent)
}

func (h hclEncoder) Encode(v interface{}, hasIndent ...bool) ([]byte, error) {
	return Encode(v, len(hasIndent) > 0 && hasIndent[0])
}

func Decode(d []byte, v interface{}) error {
	return hcl.Unmarshal(d, v)
}

func (h hclEncoder) Decode(d []byte, v interface{}) error {
	return Decode(d, v)
}

func (h hclEncoder) String() encoder.Provider {
	return encoder.HCL
}

func NewEncoder() encoder.Encoder {
	return hclEncoder{}
}