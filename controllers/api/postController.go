package api

import (
	"encoding/hex"
	"encoding/json"
	"forum/models"
	"forum/repository"
	"forum/tools"
	"forum/tools/authorization"
	"forum/tools/session"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Check if a post is containted in an array of post
func ContainsPost(AllPosts []models.Post, post models.Post) bool {
	for i := range AllPosts {
		if AllPosts[i].Id == post.Id {
			return true
		}
	}
	return false
}

//Get all post
func GetAllPost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allPosts, err := repository.GetAllPost()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	posts, err := json.Marshal(allPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(posts)
}

//Get random post id
func GetNbRandomPost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allPosts, err := repository.GetAllPost()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	nb, err := strconv.Atoi(paramsURL["nb"])
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Need number in url \"}"))
		return
	}
	if nb > len(*allPosts) {
		posts, err := json.Marshal(allPosts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
		w.Write(posts)
		return
	}
	result := []models.Post{}
	for i := 0; i < nb; i++ {
		temp := rand.Intn(len(*allPosts))
		for ContainsPost(result, (*allPosts)[temp]) {
			temp = rand.Intn(len(*allPosts))
		}
		result = append(result, (*allPosts)[temp])
	}
	posts, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(posts)
}

//Search a post
func SearchPost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allPosts, err := repository.GetAllPost()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	searchPosts := []models.Post{}
	word, err := url.QueryUnescape(paramsURL["word"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	for i := range *allPosts {
		if strings.Contains(strings.ToLower((*allPosts)[i].Title), strings.ToLower(word)) {
			searchPosts = append(searchPosts, (*allPosts)[i])
		}
	}
	posts, err := json.Marshal(searchPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(posts)
}

//Get a post by Id
func GetPostById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	postById, err := repository.GetPost("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	post, err := json.Marshal(postById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(post)
}

//Delete a post by Id
func DeletePostById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	user, err := repository.GetUser("UUID", UUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if user.Role != "admin" {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Forbidden\"}"))
		return
	}
	err = repository.DeletePost("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Create a post
func CreatePost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	var tags = []string{}
	for _, i := range params["tags"].([]interface{}) {
		tags = append(tags, i.(string))
	}
	var imageData []byte
	if params["image"].(string) == "" {
		imageData = tools.CreateImg()
	} else {
		imageData, err = hex.DecodeString(params["image"].(string))
		if err != nil {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	err = repository.CreatePost(
		params["title"].(string),
		params["description"].(string),
		sess.SessionID(),
		params["parent"].(string),
		imageData,
		int(params["nsfw"].(float64)),
		tags,
	)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Modify a post by his Id
func PutPostById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	postById, err := repository.GetPost("id", params["id"].(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	user, err := repository.GetUser("UUID", UUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if user.Role != "admin" {
		if postById.Owner != UUID.(string) {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"You are not the owner of this post\"}"))
			return
		}
	}
	if params["title"] != nil {
		postById.Title = params["title"].(string)
	}
	if params["description"] != nil {
		postById.Description = params["description"].(string)
	}
	if params["nsfw"] != nil {
		postById.NSFW = int(params["nsfw"].(float64))
	}
	if params["image"] != nil {
		if params["image"].(string) != "" {
			imageData, err := hex.DecodeString(params["image"].(string))
			if err != nil {
				w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
				return
			}
			postById.Image = imageData
		}
	}
	if params["tags"] != nil {
		postById.Tags = params["tags"].(string)
	}
	if params["upvotes"] != nil {
		postById.UpVotes = params["upvotes"].(string)
	}
	if params["downvotes"] != nil {
		postById.DownVotes = params["downvotes"].(string)
	}
	if params["publishdate"] != nil {
		postById.PublishDate = params["publishdate"].(string)
	}
	if params["comments"] != nil {
		postById.Comments = params["comments"].(string)
	}
	err = repository.PostPost(*postById, "id", postById.Id)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	post, err := json.Marshal(postById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if post != nil {
		w.Write(post)
	}
}

//Get the last created post
func GetLastPost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allPosts, err := repository.GetLastPost()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	nb, err := strconv.Atoi(paramsURL["nb"])
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Need number in url \"}"))
		return
	}
	if nb > len(*allPosts) {
		posts, err := json.Marshal(allPosts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
		w.Write(posts)
		return
	}
	*allPosts = (*allPosts)[:nb]

	posts, err := json.Marshal(allPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(posts)
}
