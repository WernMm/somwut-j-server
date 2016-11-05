package main

import (
	"backend/portfolio"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

var projectMgr = portfolio.NewProjectReposity()

func getPortfolios(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Hello, world!")
	var res struct {
		Projects []*portfolio.Project
	}
	res.Projects = projectMgr.All()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		// log.Println(err)
	}
}

func getPortfolioById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	idInt, err := strconv.Atoi(param.ByName("id"))

	if project, ok := projectMgr.GetProjectById(idInt); ok {
		var res struct {
			Project *portfolio.Project
		}
		res.Project = project

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "oops", http.StatusInternalServerError)
			// log.Println(err)
		}
	}

	// if !ok {
	// }

	// strconv.Ato ps.ByName("id")

}
