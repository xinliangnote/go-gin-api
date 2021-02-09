package grpc

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var jsonMarshal = &jsonpb.Marshaler{
	OrigName:     true,
	EmitDefaults: true,
}

// ProtoMessage2JSON marshal protobuf message to json string
func ProtoMessage2JSON(message proto.Message) (string, error) {
	if message == nil {
		return "", errors.New("message required")
	}

	raw, err := jsonMarshal.MarshalToString(message)
	if err != nil {
		return "", errors.Wrap(err, "marshal protobuf message to json err")
	}

	return raw, nil
}

// ProtoMessage2Map marshal protobuf message to map[string]interface{}
func ProtoMessage2Map(message proto.Message) (map[string]interface{}, error) {
	if message == nil {
		return nil, errors.New("message required")
	}

	raw := bytes.NewBuffer(nil)
	if err := jsonMarshal.Marshal(raw, message); err != nil {
		return nil, errors.Wrap(err, "marshal protobuf message to map err")
	}

	var mp map[string]interface{}
	if err := json.Unmarshal(raw.Bytes(), &mp); err != nil {
		return nil, errors.Wrap(err, "marshal protobuf message to map err")
	}

	return mp, nil
}
