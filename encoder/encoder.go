package encoder

var encoders = make(map[Provider]Encoder)

type Provider string

type Encoder interface {
	Encode(interface{}, ...bool) ([]byte, error)
	Decode([]byte, interface{}) error
	String() Provider
}

func Register(name Provider, enc Encoder) {
	encoders[name] = enc
}

func GetEncoder(name Provider) Encoder {
	if _, ok := encoders[name]; !ok {
		panic("encoders: Register called twice for encoder: " + name)
	}
	return encoders[name]
}
