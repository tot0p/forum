package riot

import (
	"encoding/json"
	"fmt"
	modelsriot "forum/tools/riot/modelsRiot"
	"io/ioutil"
	"net/http"
)

var (
	euw = "https://euw1.api.riotgames.com"
	API = riotAPI{}
)

type riotAPI struct {
	Key_api string
}

//Method to set the key for Riot api
func (r *riotAPI) SetKey(key string) {
	r.Key_api = key
}

//Method to get an user using an username from the Riot api
func (riot *riotAPI) GetUserByName(pseudo string) modelsriot.User {
	if pseudo == "" {
		return modelsriot.User{}
	}
	resp, err := http.Get(euw + "/lol/summoner/v4/summoners/by-name/" + pseudo + riot.sign())
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := modelsriot.User{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result

}

//Method to get an user using an id from the Riot api
func (r *riotAPI) GetUserById(id string) modelsriot.User {
	if id == "" {
		return modelsriot.User{}
	}
	resp, err := http.Get(euw + "/lol/summoner/v4/summoners/" + id + r.sign())
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := modelsriot.User{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

//Method to get a rank using an id from the Riot api
func (riot *riotAPI) GetRankById(id string) []modelsriot.Rank {
	if id == "" {
		return []modelsriot.Rank{}
	}
	resp, err := http.Get(euw + "/lol/league/v4/entries/by-summoner/" + id + riot.sign())
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := []modelsriot.Rank{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

//Method to add the api key to the url
func (riot *riotAPI) sign() string {
	return "?api_key=" + riot.Key_api
}
