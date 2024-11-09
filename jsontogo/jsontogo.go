package jsontogo

import (
	"bytes"

	"github.com/ChimeraCoder/gojson"
	json_to_go "github.com/kumakichi/json-to-go"
)

func JsonToGoWithPkg(data string) (string, error) {
	var parser gojson.Parser = gojson.ParseJson
	if output, err := gojson.Generate(bytes.NewBufferString(data), parser, "Payload", "", []string{"json"}, true, true); err != nil {
		return "", err
	} else {
		return string(output), nil
	}
}

func JsonToGo(data string, structName string) (string, error) {
	output, err := json_to_go.Parse(data, json_to_go.Options{TypeName: structName})
	return output, err
}
