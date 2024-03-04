package json

import (
	"bytes"
	"encoding/json"
	"errors"
)

func Byte2Struct(body []byte, ss interface{}) error {
	err := json.Unmarshal(body, ss)
	if nil != err {
		return err
	}
	return nil
}

func Json2Struct(body string, ss interface{}) error {
	err := json.Unmarshal([]byte(body), ss)
	if nil != err {
		return err
	}
	return nil
}

func Struct2Json(ss interface{}) (string, error) {
	res, err := json.Marshal(ss)
	if nil != err {
		return "", err
	}
	return string(res), nil
}

func Json2Map(parm interface{}) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	switch s_type := parm.(type) {
	case string:
		decoder := json.NewDecoder(bytes.NewReader([]byte(string(s_type))))
		decoder.UseNumber()
		err := decoder.Decode(&resultMap)
		if err != nil {
			return nil, err
		}
	case map[string]interface{}:
		resultMap = parm.(map[string]interface{})
	default:
		return nil, errors.New("ToMap is err. input not string, not map")
	}
	return resultMap, nil
}
