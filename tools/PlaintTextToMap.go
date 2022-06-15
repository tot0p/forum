package tools

import (
	"strings"
)

//Function to convert a plain text to a map
func PlaintTextToMap(Request []byte) map[string]string {
	Str := string(Request)
	Datas := strings.Split(Str, "&")
	Data := make(map[string]string)
	var key string
	var value string
	var Splited []string
	for i := 0; i < len(Datas); i++ {
		Str = Datas[i]
		Splited = strings.Split(Str, "=")
		if len(Splited) == 2 {
			key = Splited[0]
			value = Splited[1]
			Data[key] = value
		}
	}
	return Data
}
