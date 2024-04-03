package main

import (
	"HTMX_GO/components"
	"HTMX_GO/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	models.InitData()
	data := models.Data[0:10]
	components.Index(data).Render(r.Context(), w)
}

func addConfiguration(w http.ResponseWriter, r *http.Request) {
	newConf := models.Configuration{}
	r.ParseForm()
	newConf.Name = r.Form.Get("Name")
	newConf.AppId = r.Form.Get("AppId")
	newConf.Expression = r.Form.Get("Expression")
	models.AddItem(newConf)
	components.Row(0, newConf).Render(r.Context(), w)
}

func removeConfiguration(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	removedIndex, err := strconv.Atoi(param)
	if err != nil {
		// If the conversion failed, return an error message to the client
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	models.DeleteItem(removedIndex - 1)
	w.WriteHeader(http.StatusOK)
}

func getConfigurations(w http.ResponseWriter, r *http.Request) {
	page := 2
	itemsPerPage := 10

	pageStr := r.URL.Query().Get("page")
	parsedPage, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	page = parsedPage
	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage

	data := models.Data[start:end]
	components.Rows(data).Render(r.Context(), w)
}
