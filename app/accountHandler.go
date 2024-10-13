package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"server/dto"
	"server/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listingId := vars["listing_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.ListingId = listingId
		account, appError := a.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	listingId := vars["listing_id"]

	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.AccountId = listingId

		account, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
