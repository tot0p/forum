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

func (r *riotAPI) SetKey(key string) {
	r.Key_api = key
}

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

func (riot *riotAPI) sign() string {
	return "?api_key=" + riot.Key_api
}
