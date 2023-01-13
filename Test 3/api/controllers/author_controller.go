package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zhafranammar/rest-api/api/models"
	"github.com/zhafranammar/rest-api/api/responses"
	"github.com/zhafranammar/rest-api/api/utils/formaterror"
)


func (server *Server) CreateAuthor(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	author := models.Author{}
	err = json.Unmarshal(body, &author)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	author.Prepare()
	err = author.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	authorCreated, err := author.StoreAuthor(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, authorCreated.ID))
	responses.JSON(w, http.StatusCreated, struct {
			Message string `json:"message"`
			Data 		models.Author `json:"data"`
		}{
			Message: "Success Created an Author",
			Data: *authorCreated,
		})
}

func (server *Server) GetAuthors(w http.ResponseWriter, r *http.Request) {

	author := models.Author{}

	authors, err := author.FindAllAuthor(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		
		return
	}
	responses.JSON(w, http.StatusOK, authors)
}

func (server *Server) GetAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	author := models.Author{}

	authorReceived, err := author.FindAuthorByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, authorReceived)
}

func (server *Server) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	aid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	author := models.Author{}
	err = json.Unmarshal(body, &author)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	author.Prepare()
	err = author.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	authorUpdated, err := author.UpdateAuthor(server.DB, aid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
			Data 		models.Author `json:"data"`
		}{
			Message: "Success Updated an Author",
			Data: *authorUpdated,
		})
}


func (server *Server) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	author := models.Author{}
	err = server.DB.Debug().Model(models.Author{}).Where("id = ?", pid).Take(&author).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = author.DeleteAuthor(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "Success Deleted an Author",
		})
}