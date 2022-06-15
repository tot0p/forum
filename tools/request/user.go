package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"forum/models"
	"forum/tools/authorization"
	"io/ioutil"
	"net/http"
	"os"
)

//Function to get all the users from the api
func GetAllUser(SID string) ([]models.User, error) {
	//init
	All_user := []models.User{}
	url := os.Getenv("url_api") + "users"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//request
	authorization.SetAuthorizationBearer(SID, req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return nil, errors.New(jsonReqBody["msg"].(string))
		}
		return nil, errors.New(jsonReqBody["err"].(string))
	}
	err = json.Unmarshal(reqBody, &All_user)
	if err != nil {
		return nil, err
	}
	return All_user, nil
}

//Function to get an user using an username from the api
func GetUserByName(name string) (models.User, error) {
	user := models.User{}
	url := os.Getenv("url_api") + "user/by-username/" + name
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return models.User{}, errors.New(jsonReqBody["msg"].(string))
		}
		return models.User{}, errors.New(jsonReqBody["err"].(string))
	}
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//Function to get your user from the api
func GetMe(SID string) (models.User, error) {
	user := models.User{}
	url := os.Getenv("url_api") + "user"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, err
	}
	authorization.SetAuthorizationBearer(SID, req)
	resp, err := client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return models.User{}, errors.New(jsonReqBody["msg"].(string))
		}
		return models.User{}, errors.New(jsonReqBody["err"].(string))
	}
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//Function to modify an user on the api
func PutUser(user models.User, SID string) error {
	url := os.Getenv("url_api") + "user"
	client := &http.Client{}
	modifiedUser := make(map[string]interface{})
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = json.Unmarshal(userBytes, &modifiedUser)
	if err != nil {
		return err
	}
	if modifiedUser["profilepicture"].(string) != "" {
		modifiedUser["profilepicture"] = fmt.Sprintf("%02x", user.ProfilePicture)
	} else {
		modifiedUser["profilepicture"] = ""
	}
	data, err := json.Marshal(modifiedUser)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
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
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		return err
	}
	return nil
}

//Function to get the username of an user from the api
func GetUserUsername(UUID string) string {
	client := &http.Client{}
	url := os.Getenv("url_api") + "username"
	params := make(map[string]string)
	params["UUID"] = UUID
	data, err := json.Marshal(params)
	if err != nil {
		return ""
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return ""
	}
	req.Header.Set("content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var jsonReqBody map[string]string
	json.Unmarshal(reqBody, &jsonReqBody)
	return jsonReqBody["username"]
}

//Function to delete an user using an id on the api
func DeleteUserById(id, SID string) error {
	url := os.Getenv("url_api") + "user/" + id
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

func GetSearchUser(s string) ([]models.User, error) {
	All_user := []models.User{}
	url := os.Getenv("url_api") + "user/search/" + s
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &All_user)
	if err != nil {
		return nil, err
	}
	return All_user, nil
}
