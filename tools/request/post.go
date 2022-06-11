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

func GetAllPost() ([]models.Post, error) {
	AllPosts := []models.Post{}
	url := os.Getenv("url_api") + "posts"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
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
	err = json.Unmarshal(reqBody, &AllPosts)
	if err != nil {
		return nil, err
	}
	return AllPosts, nil
}

func GetPostById(id string) (models.Post, error) {
	Post := models.Post{}
	url := os.Getenv("url_api") + "post/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Post{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Post{}, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Post{}, err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return models.Post{}, errors.New(jsonReqBody["msg"].(string))
		}
		return models.Post{}, errors.New(jsonReqBody["err"].(string))
	}
	err = json.Unmarshal(reqBody, &Post)
	if err != nil {
		return models.Post{}, err
	}
	return Post, nil
}

func PutPost(post models.Post, SID string) error {
	url := os.Getenv("url_api") + "post"
	client := &http.Client{}
	modifiedPost := make(map[string]interface{})
	postBytes, err := json.Marshal(post)
	if err != nil {
		return err
	}
	err = json.Unmarshal(postBytes, &modifiedPost)
	if err != nil {
		return err
	}
	if modifiedPost["image"].(string) != "" {
		modifiedPost["image"] = fmt.Sprintf("%02x", post.Image)
	} else {
		modifiedPost["image"] = ""
	}

	data, err := json.Marshal(modifiedPost)
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
	err = json.Unmarshal(reqBody, &post)
	if err != nil {
		return err
	}
	return nil
}
