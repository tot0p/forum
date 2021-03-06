package controllers

import (
	"fmt"
	"forum/controllers/client"
	adminpage "forum/controllers/client/adminPage"
	datapage "forum/controllers/client/dataPage"
	"net/http"
	"strings"
)

type HandlerDyn interface {
	ServeHTTP(http.ResponseWriter, *http.Request, map[string]string)
}

type ClientController struct {
	Page404 HandlerDyn
}

/*
The controller for the API, contains the path definition functions linked to the client
*/
func (a *ClientController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	var des = false
	fmt.Println("Method :", r.Method, "At :", r.URL.Path, "By :", r.RemoteAddr)
	//To add the path for standard page
	des = des || pathDynPage("/", key, w, r, &client.Page{Path: "index.html"})
	des = des || pathDynPage("/signout", key, w, r, &client.SignOutPage{})
	des = des || pathDynPage("/profile", key, w, r, &client.ProfilePage{Path: "profile.html"})
	des = des || pathDynPage("/user", key, w, r, &client.ProfilePage{Path: "profile.html"})
	des = des || pathDynPage("/user/:name", key, w, r, &client.UserPage{Path: "user.html"})
	des = des || pathDynPage("/register", key, w, r, &client.RegisterPage{Path: "register.html"})
	des = des || pathDynPage("/login", key, w, r, &client.LoginPage{Path: "login.html"})
	des = des || pathDynPage("/subject", key, w, r, &client.CreateSubjectPage{Path: "create_subject.html"})
	des = des || pathDynPage("/subject/:id", key, w, r, &client.SubjectPage{Path: "subject.html"})
	des = des || pathDynPage("/data", key, w, r, &datapage.UserDataPage{Path: "data.html"})
	des = des || pathDynPage("/admin", key, w, r, &adminpage.AllUser{Path: "all_user.html"})
	des = des || pathDynPage("/admin/subjects", key, w, r, &adminpage.AllSubject{Path: "all_subject.html"})
	des = des || pathDynPage("/admin/posts", key, w, r, &adminpage.AllPost{Path: "all_post.html"})
	des = des || pathDynPage("/post", key, w, r, &client.CreatePostPage{Path: "create_post.html"})
	des = des || pathDynPage("/post/:id", key, w, r, &client.PostPage{Path: "post.html"})
	des = des || pathDynPage("/update-profile", key, w, r, &client.UpdateProfilePage{Path: "edit_profile.html"})
	des = des || pathDynPage("/update-subject/:id", key, w, r, &client.UpdateSubjectPage{Path: "edit_subject.html"})
	des = des || pathDynPage("/update-post/:id", key, w, r, &client.UpdatePostPage{Path: "edit_post.html"})
	des = des || pathDynPage("/explorer", key, w, r, &client.SubjectExplorerPage{Path: "explorer.html"})
	des = des || pathDynPage("/GTU", key, w, r, &client.StaticPage{Path: "GTU.html"})
	des = des || pathDynPage("/privacy", key, w, r, &client.StaticPage{Path: "privacy.html"})
	des = des || pathDynPage("/team", key, w, r, &client.TeamPage{Path: "team.html"})
	des = des || pathDynPage("/license", key, w, r, &client.StaticPage{Path: "license.html"})
	des = des || pathDynPage("/contact", key, w, r, &client.StaticPage{Path: "Contact.html"})
	des = des || pathDynPage("/search", key, w, r, &client.PageSearch{Path: "search.html"})
	des = des || pathDynPage("/tag/:type/:tag", key, w, r, &client.TagPage{Path: "tags.html"})

	if !des {
		pathDynPage(key, key, w, r, &client.Page404{Path: "404.html"})
	}
}

/*
Dynamic Path
We use a fonction to create an id that will be stored and reused to generate a new path
*/
func pathDynPage(path, actualPath string, w http.ResponseWriter, r *http.Request, page HandlerDyn) bool {
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
	page.ServeHTTP(w, r, m)
	return true
}
