package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL ="https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
	
}
type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User  *User
	CreateAt time.Time `json:"create_at"`
	Body string
}
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}
func SearchIssues(terms []string)(*IssueSearchResult,error){
	q :=url.QueryEscape(strings.Join(terms," "))
	resp,error :=http.Get(IssuesURL +"?q="+q)
	if error !=nil{
        return nil,error
	}
    if resp.StatusCode !=http.StatusOK{
    	resp.Body.Close()
    	return  nil,fmt.Errorf("search query failed: %s",resp.Status)
	}
	var result IssueSearchResult
	if error :=json.NewDecoder(resp.Body).Decode(&result);error!=nil{
		resp.Body.Close()
		return nil,error
	}
	resp.Body.Close()
	return &result,nil
}
