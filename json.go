package main

import "encoding/json"

// CodeJSON is ...
type CodeJSON struct {
	Code    int
	Message string
	Data    interface{}
}

// ReJSON is ...
func ReJSON(data interface{}) (str string) {

	resp := new(CodeJSON)

	resp.Code = 0
	resp.Message = Config.Message.Success
	resp.Data = data

	jsons, err := json.Marshal(resp)

	CheckErr(err)

	str = string(jsons)

	return
}
