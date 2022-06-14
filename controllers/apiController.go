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

/*
The controller for the API, contains the path definition functions linked to the api
And Parse variable Post or Get and Put
*/
func (a *APIController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	elements := strings.Replace(key, "/api", "", 1)
	repository.AddConnection(r.RemoteAddr)
	if elements == "/" {
		w.Write([]byte("Doc"))
	} else {
		w.Header().Set("content-Type", "application/json")
		fmt.Println("Method :", r.Method, "At :", r.URL.Path, "By :", r.RemoteAddr)
		des := false
		switch r.Method {
		//Function to get some docs with the API
		case "GET":
			var params = map[string]interface{}{}
			for k, elem := range r.URL.Query() {
				params[k] = strings.Join(elem, ",")
			}
			//docs
			//TODO Api docs
			//user

			des = des || pathDyn("/users", elements, params, api.GetAllUsers, w, r)
			des = des || pathDyn("/user", elements, params, api.GetUserMe, w, r)
			des = des || pathDyn("/username/:id", elements, params, api.GetUserUsername, w, r)
			des = des || pathDyn("/user/:id", elements, params, api.GetUserById, w, r)
			des = des || pathDyn("/user/search/:word", elements, params, api.SearchUser, w, r)
			des = des || pathDyn("/user/by-username/:username", elements, params, api.GetUserByUsername, w, r)
			//subject
			des = des || pathDyn("/subjects", elements, params, api.GetAllSubject, w, r)
			des = des || pathDyn("/subject/:id", elements, params, api.GetSubjectById, w, r)
			des = des || pathDyn("/subject/:id/upvote", elements, params, api.SubjectLike, w, r)
			des = des || pathDyn("/subject/:id/count", elements, params, api.SubjectCount, w, r)
			des = des || pathDyn("/subject/:id/downvote", elements, params, api.SubjectHate, w, r)
			des = des || pathDyn("/subject/:id/vote", elements, params, api.UserLikeOrHateSubject, w, r)
			des = des || pathDyn("/subject/search/:word", elements, params, api.SearchSubject, w, r)
			des = des || pathDyn("/subject/GetNBSubject/:nb", elements, params, api.GetNbRandomSubject, w, r)
			des = des || pathDyn("/subject/GetLastSubjectUpdate/:nb", elements, params, api.GetSubjectLastUpdate, w, r)
			des = des || pathDyn("/subject/GetSubjectsByUser/:id", elements, params, api.GetSubjectByUser, w, r)

			//post
			des = des || pathDyn("/posts", elements, params, api.GetAllPost, w, r)
			des = des || pathDyn("/post/:id", elements, params, api.GetPostById, w, r)
			des = des || pathDyn("/post/:id/upvote", elements, params, api.PostLike, w, r)
			des = des || pathDyn("/post/:id/count", elements, params, api.PostCount, w, r)
			des = des || pathDyn("/post/:id/downvote", elements, params, api.PostHate, w, r)
			des = des || pathDyn("/post/:id/vote", elements, params, api.UserLikeOrHatePost, w, r)
			des = des || pathDyn("/post/search/:word", elements, params, api.SearchPost, w, r)
			des = des || pathDyn("/post/GetNBPost/:nb", elements, params, api.GetNbRandomPost, w, r)
			des = des || pathDyn("/post/GetLastPost/:nb", elements, params, api.GetLastPost, w, r)
			des = des || pathDyn("/post/GetPostsBySubject/:id", elements, params, api.GetPostsBySubjectId, w, r)
			des = des || pathDyn("/post/GetPostsByUser/:id", elements, params, api.GetPostsByUserId, w, r)

			//comment
			des = des || pathDyn("/comments", elements, params, api.GetAllComment, w, r)
			des = des || pathDyn("/comment/:id", elements, params, api.GetCommentById, w, r)
			des = des || pathDyn("/comment/:id/upvote", elements, params, api.CommentLike, w, r)
			des = des || pathDyn("/comment/:id/count", elements, params, api.CommentCount, w, r)
			des = des || pathDyn("/comment/:id/downvote", elements, params, api.CommentHate, w, r)
			des = des || pathDyn("/comment/:id/vote", elements, params, api.UserLikeOrHateComment, w, r)
			des = des || pathDyn("/comment/GetCommentByPost/:id", elements, params, api.GetCommentsByPostId, w, r)

			//count
			des = des || pathDyn("/count", elements, params, api.GetCount, w, r)
			des = des || pathDyn("/count/user", elements, params, api.GetNBUser, w, r)
			des = des || pathDyn("/count/post", elements, params, api.GetNBPost, w, r)
			des = des || pathDyn("/count/subject", elements, params, api.GetNBSubject, w, r)
			des = des || pathDyn("/count/session", elements, params, api.GetNBUserConnected, w, r)

		case "POST":
			//Function to post some info into the api
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
			des = des || pathDyn("/user", elements, params, api.CreateUser, w, r)
			des = des || pathDyn("/subject", elements, params, api.CreateSubject, w, r)
			des = des || pathDyn("/post", elements, params, api.CreatePost, w, r)
			des = des || pathDyn("/comment", elements, params, api.CreateComment, w, r)
			des = des || pathDyn("/login", elements, params, api.UserLogin, w, r)

		case "PUT":
			//Function to modify some info in the api
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
			des = des || pathDyn("/user", elements, params, api.PutUserById, w, r)
			des = des || pathDyn("/subject", elements, params, api.PutSubjectsById, w, r)
			des = des || pathDyn("/post", elements, params, api.PutPostById, w, r)
			des = des || pathDyn("/comment", elements, params, api.PutCommentById, w, r)

		case "DELETE":
			//Function to delete some info in the api
			des = des || pathDyn("/signout", elements, map[string]interface{}{}, api.DeleteSession, w, r)
			des = des || pathDyn("/user/:id", elements, map[string]interface{}{}, api.DeleteUserById, w, r)
			des = des || pathDyn("/subject/:id", elements, map[string]interface{}{}, api.DeleteSubjectById, w, r)
			des = des || pathDyn("/post/:id", elements, map[string]interface{}{}, api.DeletePostById, w, r)
			des = des || pathDyn("/comment/:id", elements, map[string]interface{}{}, api.DeleteCommentById, w, r)

		}
		if !des {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"msg": "not found"}`))
		}
	}
}

/*
Dynamic Path
We use a fonction to create an id that will be stored and reused to generate a new path
*/
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
