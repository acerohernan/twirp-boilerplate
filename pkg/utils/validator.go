package utils

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/proto"
)

type ProtoValidator struct {
	val *protovalidate.Validator
}

func NewProtoValidator() (*ProtoValidator, error) {
	val, err := protovalidate.New()

	if err != nil {
		return nil, err
	}

	return &ProtoValidator{
		val: val,
	}, nil
}

func (v *ProtoValidator) Validate(msg proto.Message) error {
	return v.val.Validate(msg)
}
