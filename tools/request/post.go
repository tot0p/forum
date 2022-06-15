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

//Function to get all the posts from the api
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

//Function to get a post using an id from the api
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

//Function to modify a post in the api using an id
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

//Function to delete a post in the api using an id
func DeletePostById(id, SID string) error {
	url := os.Getenv("url_api") + "post/" + id
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

//Function to get all the posts related to a subject using an id from the api
func GetPostsBySubjectId(id string) ([]models.Post, error) {
	url := os.Getenv("url_api") + "post/GetPostsBySubject/" + id
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonReqBody = []models.Post{}
	err = json.Unmarshal(reqBody, &jsonReqBody)
	if err != nil {
		return nil, err
	}
	return jsonReqBody, nil
}

//Function to get all the posts created by an user using an id from the api
func GetPostsByUserId(id string) ([]models.Post, error) {
	url := os.Getenv("url_api") + "post/GetPostsByUser/" + id
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonReqBody = []models.Post{}
	err = json.Unmarshal(reqBody, &jsonReqBody)
	if err != nil {
		return nil, err
	}
	return jsonReqBody, nil
}

//Function to upvote a post in the api
func LikePost(id, SID, votes string) (string, error) {
	url := os.Getenv("url_api") + "post/" + id
	if votes == "upvote" {
		url += "/upvote"
	} else if votes == "downvote" {
		url += "/downvote"
	} else {
		return "", nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("content-Type", "application/json")
	authorization.SetAuthorizationBearer(SID, req)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	result := ""
	err = json.Unmarshal(reqBody, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func GetSearchPosts(s string) ([]models.Post, error) {
	AllPosts := []models.Post{}
	url := os.Getenv("url_api") + "post/search/" + s
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
	err = json.Unmarshal(reqBody, &AllPosts)
	if err != nil {
		return nil, err
	}
	return AllPosts, nil
}

func GetPostsByTag(s string) ([]models.Post, error) {
	AllPosts := []models.Post{}
	url := os.Getenv("url_api") + "post/GetPostsByTag/" + s
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
	err = json.Unmarshal(reqBody, &AllPosts)
	if err != nil {
		return nil, err
	}
	return AllPosts, nil
}
