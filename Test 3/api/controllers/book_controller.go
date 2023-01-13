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


func (server *Server) CreateBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	book.Prepare()
	err = book.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bookCreated, err := book.StoreBook(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, bookCreated.ID))
	responses.JSON(w, http.StatusCreated, struct {
			Message string `json:"message"`
			Data 		models.Book `json:"data"`
		}{
			Message: "Success Created an Book",
			Data: *bookCreated,
		})
}

func (server *Server) GetBooks(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}

	books, err := book.FindAllBook(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		
		return
	}
	responses.JSON(w, http.StatusOK, books)
}

func (server *Server) GetBookPublisher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	book := models.Book{}
	publisher := models.Publisher{}
	bookReceived, err := book.FindBookByPublisher(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	publisherReceived, err := publisher.FindPublisherByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
			Publisher string `json:"publisher"`
			Data 		[]models.Book `json:"data"`
		}{
			Publisher: publisherReceived.Name,
			Data: *bookReceived,
		})
}

func (server *Server) GetBookAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	book := models.Book{}
	author := models.Author{}
	bookReceived, err := book.FindBookByAuthor(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	authorReceived, err := author.FindAuthorByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
			Author string `json:"author"`
			Data 		[]models.Book `json:"data"`
		}{
			Author: authorReceived.Fullname,
			Data: *bookReceived,
		})
}

func (server *Server) GetBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	book := models.Book{}

	bookReceived, err := book.FindBookByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, bookReceived)
}

func (server *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {
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
	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	book.Prepare()
	err = book.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bookUpdated, err := book.UpdateBook(server.DB, aid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
			Data 		models.Book `json:"data"`
		}{
			Message: "Success Updated an Book",
			Data: *bookUpdated,
		})
}


func (server *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	book := models.Book{}
	err = server.DB.Debug().Model(models.Book{}).Where("id = ?", pid).Take(&book).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = book.DeleteBook(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "Success Deleted an Book",
		})
}