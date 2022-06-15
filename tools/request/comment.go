package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"forum/models"
	"forum/tools/authorization"
	"io/ioutil"
	"net/http"
	"os"
)

//Function to get all the comments from the api
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

//Function to get a comment using an id from the api
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

//Function to create a comment in the api
func PostComment(comment models.Comment, SID string) error {
	url := os.Getenv("url_api") + "comment"
	client := &http.Client{}
	commentBytes, err := json.Marshal(comment)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(commentBytes))
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
	err = json.Unmarshal(reqBody, &comment)
	if err != nil {
		return err
	}
	return nil
}

//Function to get all the comments of a post using an id
func GetCommentsByPostId(id string) ([]models.Comment, error) {
	AllComments := []models.Comment{}
	url := os.Getenv("url_api") + "comment/GetCommentByPost/" + id
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
	err = json.Unmarshal(reqBody, &AllComments)
	if err != nil {
		return nil, err
	}
	return AllComments, nil
}

//Function to upvote or downvote a comment
func LikeComment(id, SID, votes string) error {
	url := os.Getenv("url_api") + "comment/" + id
	if votes == "upvote" {
		url += "/upvote"
	} else if votes == "downvote" {
		url += "/downvote"
	} else {
		return nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
	err = json.Unmarshal(reqBody, &votes)
	if err != nil {
		return err
	}
	return nil
}
