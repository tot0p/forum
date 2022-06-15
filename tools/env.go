package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Function to load the environnement
func LoadEnv(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	datas := string(data)
	sep := "\n"
	if strings.Contains(datas, "\r") {
		sep = "\r" + sep
	}
	var result = map[string]string{}
	for _, elem := range strings.Split(datas, sep) {
		if strings.Contains(elem, "=") {
			t := strings.SplitN(elem, "=", 2)
			if len(t) > 1 {
				result[t[0]] = t[1]
			} else {
				return errors.New("error : syntax error")
			}
		}
	}
	for key, value := range result {
		os.Setenv(key, value)
	}
	return nil
}
