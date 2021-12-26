package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tabinnorway/dive-arch/golang-server/api/auth"
	"github.com/tabinnorway/dive-arch/golang-server/api/models"
	"github.com/tabinnorway/dive-arch/golang-server/api/responses"
	"github.com/tabinnorway/dive-arch/golang-server/api/utils/formaterror"
)

func (server *Server) CreateDive(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dive := models.Dive{}
	err = json.Unmarshal(body, &dive)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dive.Prepare()
	err = dive.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != dive.DiverID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	diveCreated, err := dive.SaveDive(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, diveCreated.ID))
	responses.JSON(w, http.StatusCreated, diveCreated)
}

func (server *Server) GetDives(w http.ResponseWriter, r *http.Request) {

	dive := models.Dive{}

	dives, err := dive.FindAllDives(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dives)
}

func (server *Server) GetDive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	dive := models.Dive{}

	diveReceived, err := dive.FindDiveByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, diveReceived)
}

func (server *Server) UpdateDive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Check if the dive id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the dive exist
	dive := models.Dive{}
	err = server.DB.Model(models.Dive{}).Where("id = ?", pid).Take(&dive).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("dive not found"))
		return
	}

	// If a user attempt to update a post not belonging to him
	if uid != dive.DiverID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	diveUpdate := models.Dive{}
	err = json.Unmarshal(body, &diveUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != diveUpdate.DiverID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	diveUpdate.Prepare()
	err = diveUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	diveUpdate.ID = dive.ID //this is important to tell the model the post id to update, the other update field are set above

	diveUpdated, err := diveUpdate.UpdateADive(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, diveUpdated)
}

func (server *Server) DeleteDive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	post := models.Dive{}
	err = server.DB.Model(models.Dive{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this post?
	if uid != post.DiverID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = post.DeleteADive(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
