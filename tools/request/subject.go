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

func GetAllSubject() ([]models.Subject, error) {
	AllSubjects := []models.Subject{}
	url := os.Getenv("url_api") + "subjects"
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
	err = json.Unmarshal(reqBody, &AllSubjects)
	if err != nil {
		return nil, err
	}
	return AllSubjects, nil
}

func GetSubjectById(id string) (models.Subject, error) {
	Subject := models.Subject{}
	url := os.Getenv("url_api") + "subject/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Subject{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Subject{}, err
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Subject{}, err
	}
	var jsonReqBody map[string]interface{}
	json.Unmarshal(reqBody, &jsonReqBody)
	if _, ok := jsonReqBody["err"]; ok {
		if _, ok := jsonReqBody["msg"]; ok {
			return models.Subject{}, errors.New(jsonReqBody["msg"].(string))
		}
		return models.Subject{}, errors.New(jsonReqBody["err"].(string))
	}
	err = json.Unmarshal(reqBody, &Subject)
	if err != nil {
		return models.Subject{}, err
	}
	return Subject, nil
}

func PutSubject(subject models.Subject, SID string) error {
	url := os.Getenv("url_api") + "subject"
	client := &http.Client{}
	modifiedSubject := make(map[string]interface{})
	subjectBytes, err := json.Marshal(subject)
	if err != nil {
		return err
	}
	err = json.Unmarshal(subjectBytes, &modifiedSubject)
	if err != nil {
		return err
	}
	if modifiedSubject["image"].(string) != "" {
		modifiedSubject["image"] = fmt.Sprintf("%02x", subject.Image)
	} else {
		modifiedSubject["image"] = ""
	}

	data, err := json.Marshal(modifiedSubject)
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
	err = json.Unmarshal(reqBody, &subject)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSubjectById(id, SID string) error {
	url := os.Getenv("url_api") + "subject/" + id
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
