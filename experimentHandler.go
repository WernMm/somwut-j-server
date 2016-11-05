package main

import (
	"encoding/json"
	// "fmt"
	fb "github.com/huandu/facebook"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func experiment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var globalApp = fb.New("180978795693925", "fade77647e20b926279608da39cdf9d2")
	// globalApp.Version = "v2.2"
	// globalApp.RedirectUri = "http://your.site/canvas/url/"
	context := appengine.NewContext(r)
	session := globalApp.Session(globalApp.AppAccessToken())
	session.HttpClient = urlfetch.Client(context)
	res, err := session.Get("/425641614195264/feed", nil)
	// res, err := fb.Get("/425641614195264", fb.Params{
	// "fields":       "name",
	// "access_token": "EAACEdEose0cBAJGZCZAy8ZAIH1xZCGRJTSCCDq6lSp7KxYGr3OUpW25a3Bl5XLyFjKgHUYdjbMDMgqoKAEhieN43DjCg3OyV0b7xPnqBqCW0hIzmRNci7QUmq1q6XP4qBY3ZA0Ks8kYE7v3MamRQ4P6eACv5VZCEtpFIb5u7kZCdAZDZD",
	// })
	if err != nil {
		log.Infof(context, err.Error())
	}
	// log.Println(res)
	// log.Println(err)
	// fmt.Println("here is my facebook first name:", res["first_name"])
	json.NewEncoder(w).Encode(res)
}
