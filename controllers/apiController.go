package controllers

import (
	"encoding/json"
	"fmt"
	"forum/controllers/api"
	"forum/repository"
	"io/ioutil"
	"net/http"
	"strings"
)

type APIController struct {
}

func (a *APIController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	elements := strings.Replace(key, "/api", "", 1)
	repository.AddConnection(r.RemoteAddr)
	if elements == "/" {
		w.Write([]byte("Doc"))
	} else {
		w.Header().Set("content-Type", "application/json")
		fmt.Println("Method :", r.Method, "At :", r.URL.Path, "By :", r.RemoteAddr)
		switch r.Method {
		case "GET":
			var params = map[string]interface{}{}
			for k, elem := range r.URL.Query() {
				params[k] = strings.Join(elem, ",")
			}
			//docs
			//TODO Api docs
			//user
			pathDyn("/users", elements, params, api.GetAllUsers, w, r)
			pathDyn("/user", elements, params, api.GetUserMe, w, r)
			pathDyn("/username/:id", elements, params, api.GetUserUsername, w, r)
			pathDyn("/user/:id", elements, params, api.GetUserById, w, r)
			pathDyn("/user/search/:word", elements, params, api.SearchUser, w, r)
			pathDyn("/user/by-username/:username", elements, params, api.GetUserByUsername, w, r)
			//TODO Follow and unFollow
			//subject
			pathDyn("/subjects", elements, params, api.GetAllSubject, w, r)
			pathDyn("/subject/:id", elements, params, api.GetSubjectById, w, r)
			pathDyn("/subject/:id/upvote", elements, params, api.SubjectLike, w, r)
			pathDyn("/subject/:id/count", elements, params, api.SubjectCount, w, r)
			pathDyn("/subject/:id/downvote", elements, params, api.SubjectHate, w, r)
			pathDyn("/subject/search/:word", elements, params, api.SearchSubject, w, r)
			pathDyn("/subject/GetNBSubject/:nb", elements, params, api.GetNbRandomSubject, w, r)
			pathDyn("/subject/GetLastSubjectUpdate/:nb", elements, params, api.GetSubjectLastUpdate, w, r)

			//post
			pathDyn("/posts", elements, params, api.GetAllPost, w, r)
			pathDyn("/post/:id", elements, params, api.GetPostById, w, r)
			pathDyn("/post/search/:word", elements, params, api.SearchPost, w, r)
			pathDyn("/post/GetNBPost/:nb", elements, params, api.GetNbRandomPost, w, r)
			pathDyn("/post/GetLastPost/:nb", elements, params, api.GetLastPost, w, r)

			//comment
			pathDyn("/comments", elements, params, api.GetAllComment, w, r)
			pathDyn("/comment/:id", elements, params, api.GetCommentById, w, r)

			//count
			pathDyn("/count", elements, params, api.GetCount, w, r)
			pathDyn("/count/user", elements, params, api.GetNBUser, w, r)
			pathDyn("/count/post", elements, params, api.GetNBPost, w, r)
			pathDyn("/count/subject", elements, params, api.GetNBSubject, w, r)
			pathDyn("/count/session", elements, params, api.GetNBUserConnected, w, r)

		case "POST":
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"error\":\"500\"}"))
				return
			}
			// fmt.Println("Params :", string(reqBody))
			var params map[string]interface{}
			if r.Header.Get("Content-type") != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("{\"error\":\"the Content-type must be json\"}"))
				return
			} else {
				err = json.Unmarshal(reqBody, &params)
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"error\":\"500\"}"))
					return
				}
			}
			pathDyn("/user", elements, params, api.CreateUser, w, r)
			pathDyn("/subject", elements, params, api.CreateSubject, w, r)
			pathDyn("/post", elements, params, api.CreatePost, w, r)
			pathDyn("/comment", elements, params, api.CreateComment, w, r)
			pathDyn("/login", elements, params, api.UserLogin, w, r)

		case "PUT":
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"error\":\"500\"}"))
				return
			}
			// fmt.Println("Params :", string(reqBody))
			var params = map[string]interface{}{}
			if r.Header.Get("Content-type") != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("{\"error\":\"the Content-type must be json\"}"))
				return
			} else {
				err = json.Unmarshal(reqBody, &params)
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"error\":\"500\"}"))
					return
				}
			}
			pathDyn("/user", elements, params, api.PutUserById, w, r)
			pathDyn("/subject", elements, params, api.PutSubjectsById, w, r)
			pathDyn("/post", elements, params, api.PutPostById, w, r)
			pathDyn("/comment", elements, params, api.PutCommentById, w, r)

		case "DELETE":
			pathDyn("/user/:id", elements, map[string]interface{}{}, api.DeleteUserById, w, r)
			pathDyn("/subject/:id", elements, map[string]interface{}{}, api.DeleteSubjectById, w, r)
			pathDyn("/post/:id", elements, map[string]interface{}{}, api.DeletePostById, w, r)
			pathDyn("/comment/:id", elements, map[string]interface{}{}, api.DeleteCommentById, w, r)

		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
		}
	}
}

func pathDyn(path, actualPath string, params map[string]interface{}, f func(map[string]string, map[string]interface{}, http.ResponseWriter, *http.Request), w http.ResponseWriter, r *http.Request) bool {
	temp := strings.Split(path, "/")
	temp1 := strings.Split(actualPath, "/")
	if len(temp) != len(temp1) {
		return false
	}
	var m = map[string]string{}
	for i, elem := range temp {
		if elem != temp1[i] && !strings.Contains(elem, ":") {
			return false
		} else if strings.Contains(elem, ":") {
			m[strings.Replace(elem, ":", "", 1)] = temp1[i]
		}
	}
	f(m, params, w, r)
	return true
}
