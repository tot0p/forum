package request

import (
	"encoding/json"
	"errors"
	"forum/models"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAllComment() ([]models.Comment, error) {
	AllComments := []models.Comment{}
	url := os.Getenv("url_api") + "comments"
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
	err = json.Unmarshal(reqBody, &AllComments)
	if err != nil {
		return nil, err
	}
	return AllComments, nil
}

func GetCommentById(id string) (models.Comment, error) {
	Comment := models.Comment{}
	url := os.Getenv("url_api") + "comment/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Comment{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Comment{}, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Comment{}, err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return models.Comment{}, errors.New(jsonReqBody["msg"].(string))
		}
		return models.Comment{}, errors.New(jsonReqBody["err"].(string))
	}
	err = json.Unmarshal(reqBody, &Comment)
	if err != nil {
		return models.Comment{}, err
	}
	return Comment, nil
}
