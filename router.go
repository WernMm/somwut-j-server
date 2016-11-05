package main

import (
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// // This function's name is a must. App Engine uses it to drive the requests properly.
func init() {
	router := httprouter.New()
	router.GET("/rest/portfolios", getPortfolios)
	router.GET("/rest/portfolios/:id", getPortfolioById)
	// router.GET("/rest/football/:team", getFootball)
	router.GET("/rest/experiment", experiment)
	http.Handle("/", router)
}

func main() {

}
