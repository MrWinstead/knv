package storage

import (
	"bytes"
	"encoding/gob"
	"github.com/pkg/errors"
)

var (
	_ Codec = &gobCodec{}
)

type gobCodec struct {
	instanceFactory InstanceFactory
}

func NewGobCodec(f InstanceFactory) Codec {
	created := &gobCodec{
		instanceFactory: f,
	}
	return created
}

func (*gobCodec) Marshal(v interface{}) ([]byte, error) {
	output := bytes.Buffer{}
	encodeErr := gob.NewEncoder(&output).Encode(v)
	if nil != encodeErr {
		err := errors.Wrap(encodeErr, "could not encode value using gob")
		return nil, err
	}
	return output.Bytes(), nil
}

func (g *gobCodec) Unmarshal(raw []byte) (interface{}, error) {
	instance := g.instanceFactory()
	unmarshalErr := gob.NewDecoder(bytes.NewBuffer(raw)).Decode(instance)
	if nil != unmarshalErr {
		err := errors.Wrap(unmarshalErr, "could not unmarshal value using gob")
		return nil, err
	}
	return instance, nil
}
