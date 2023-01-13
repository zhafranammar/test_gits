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


func (server *Server) CreatePublisher(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	publisher := models.Publisher{}
	err = json.Unmarshal(body, &publisher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	publisher.Prepare()
	err = publisher.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	publisherCreated, err := publisher.StorePublisher(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, publisherCreated.ID))
	responses.JSON(w, http.StatusCreated, struct {
			Message string `json:"message"`
			Data 		models.Publisher `json:"data"`
		}{
			Message: "Success Created an Publisher",
			Data: *publisherCreated,
		})
}

func (server *Server) GetPublishers(w http.ResponseWriter, r *http.Request) {

	publisher := models.Publisher{}

	publishers, err := publisher.FindAllPublisher(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		
		return
	}
	responses.JSON(w, http.StatusOK, publishers)
}

func (server *Server) GetPublisher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	publisher := models.Publisher{}

	publisherReceived, err := publisher.FindPublisherByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, publisherReceived)
}

func (server *Server) UpdatePublisher(w http.ResponseWriter, r *http.Request) {
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
	publisher := models.Publisher{}
	err = json.Unmarshal(body, &publisher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	publisher.Prepare()
	err = publisher.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	publisherUpdated, err := publisher.UpdatePublisher(server.DB, aid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
			Data 		models.Publisher `json:"data"`
		}{
			Message: "Success Updated an Publisher",
			Data: *publisherUpdated,
		})
}


func (server *Server) DeletePublisher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	publisher := models.Publisher{}
	err = server.DB.Debug().Model(models.Publisher{}).Where("id = ?", pid).Take(&publisher).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = publisher.DeletePublisher(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "Success Deleted an Publisher",
		})
}