package main

import(
	"net/http"
	"html/template"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"

)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":1111", nil)
}


type Post struct {
	UserId int `json:"userId"`
	Id     int `json:"Id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Posts []Post

type customData2 struct{
		Title string
		Posts []Post
	}

func index(w http.ResponseWriter, r *http.Request){
	
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

    if err != nil {
        log.Fatal(err)
    }

    content, _ := ioutil.ReadAll(resp.Body)

    var _posts Posts

    err = json.Unmarshal(content, &_posts)

    if err != nil {
        log.Fatal(err)
    }

    cd:= customData2{
		Title: "Home",
		Posts: _posts,
	}


    //fmt.Println(posts[0].Body)

	tpl.ExecuteTemplate(w, "index.gohtml", cd)
}

func about(w http.ResponseWriter, r *http.Request){
	type customData struct{
		Title string
		Members []string
	}

	cd:= customData{
		Title: "The Team",
		Members: []string{"Bond", "Pennywise", "Batman"},
	}


	tpl.ExecuteTemplate(w, "about.gohtml", cd)
}

