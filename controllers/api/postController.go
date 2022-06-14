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

//Function to get all the post
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

//Function to get a random id post
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

//Function to Search a post
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

//Function to get a post using an id
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

//Function to delete a post using an id
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

//Function to create a post on the website
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

//Function to modify a post using and id
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

//Function to get the last Post Created
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

//Function to like a post
func PostLike(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UserUUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	User, err := repository.GetUser("UUID", UserUUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Post, err := repository.GetPost("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	remove := false
	new := ""
	upvote := Post.ConvertUpVotes()
	for _, i := range upvote {
		if User.UUID == i {
			remove = true
		} else {
			new += "#" + i
		}
	}
	if !remove {
		new = ""
		for _, i := range Post.ConvertDownVotes() {
			if User.UUID == i {
				remove = true
			} else {
				new += "#" + i
			}
		}
		if remove {
			Post.DownVotes = new
		}
		Post.UpVotes += "#" + User.UUID
	} else {
		Post.UpVotes = new
	}
	repository.PostPost(*Post, "id", paramsURL["id"])
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Function to Hate a post
func PostHate(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UserUUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	User, err := repository.GetUser("UUID", UserUUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Post, err := repository.GetPost("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	remove := false
	new := ""
	Downvote := Post.ConvertDownVotes()
	for _, i := range Downvote {
		if User.UUID == i {
			remove = true
		} else {
			new += "#" + i
		}
	}
	if !remove {
		new = ""
		for _, i := range Post.ConvertUpVotes() {
			if User.UUID == i {
				remove = true
			} else {
				new += "#" + i
			}
		}
		if remove {
			Post.UpVotes = new
		}
		Post.DownVotes += "#" + User.UUID
	} else {
		Post.DownVotes = new
	}
	repository.PostPost(*Post, "id", paramsURL["id"])
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Function to count the like and the dislike on a post
func PostCount(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	Post, err := repository.GetPost("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Votes := models.Vote{UpVote: strings.Count(Post.UpVotes, "#"), DownVote: strings.Count(Post.DownVotes, "#")}
	result, err := json.Marshal(Votes)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(result)
}

//Function to tell if a use has already liked a post or not
func UserLikeOrHatePost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UserUUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	User, err := repository.GetUser("UUID", UserUUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Post, err := repository.GetPost("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	for _, i := range Post.ConvertUpVotes() {
		if i == User.UUID {
			w.Write([]byte("{\"downvote\":false,\"upvote\":true}"))
			return
		}
	}
	for _, i := range Post.ConvertDownVotes() {
		if i == User.UUID {
			w.Write([]byte("{\"downvote\":true,\"upvote\":false}"))
			return
		}
	}
	w.Write([]byte("{\"downvote\":false,\"upvote\":false}"))
}

func GetPostsBySubjectId(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	id := paramsURL["id"]
	allPosts, err := repository.GetPostsBySubjectId(id)
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

func GetPostsByUserId(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	id := paramsURL["id"]
	allPosts, err := repository.GetPostsByUserId(id)
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
