package request

import (
	"encoding/json"
	"errors"
	"forum/tools/authorization"
	"io/ioutil"
	"net/http"
	"os"
)

// api/signout

func SignOutUser(SID string) error {
	url := os.Getenv("url_api") + "signout"
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("content-Type", "application/json")
	authorization.SetAuthorizationBearer(SID, req)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return errors.New(jsonReqBody["msg"].(string))
		}
		return errors.New(jsonReqBody["err"].(string))
	}
	return nil
}
