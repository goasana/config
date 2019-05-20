package yaml

import (
	"github.com/ghodss/yaml"
	"github.com/goasana/config/encoder"
)

func init() {
	e := NewEncoder()
	encoder.Register(e.String(), e)
	encoder.Register(encoder.YML, e)
}

type yamlEncoder struct{}

func Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (j yamlEncoder) Encode(v interface{}, hasIndent ...bool) ([]byte, error) {
	return Encode(v)
}

func Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

func (j yamlEncoder) Decode(d []byte, v interface{}) error {
	return Decode(d, v)
}

func (j yamlEncoder) String() string {
	return encoder.YAML
}

func NewEncoder() encoder.Encoder {
	return yamlEncoder{}
}
