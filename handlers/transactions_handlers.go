package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-final-homework/models"
	"github.com/go-final-homework/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	repository models.TransactionRepository
}

func NewTransactionHandler(repository models.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repository: repository}
}

func (th *TransactionHandler) InitTransactionRoutes(router *mux.Router) {
	router.HandleFunc("/transactions", th.Add).Methods("POST")
	router.HandleFunc("/transactions", th.GetAll).Methods("GET")
	router.HandleFunc("/transactions/{id}", th.GetById).Methods("GET")
	router.HandleFunc("/transactions/{id}", th.Update).Methods("PUT")
	router.HandleFunc("/transactions/{id}", th.Delete).Methods("DELETE")
}

func (th *TransactionHandler) Add(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := utils.ParseJSON(r, &transaction); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(transaction); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid transaction: %v", errors))
		return
	}
	id, err := th.repository.AddTransaction(transaction)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]int{"id": id})
}

func (th *TransactionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	transactions, err := th.repository.GetAllTransactions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, transactions)
}

func (th *TransactionHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	transaction, err := th.repository.GetTransactionById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, transaction)
}

func (th *TransactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var transaction models.Transaction
	if err := utils.ParseJSON(r, &transaction); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(transaction); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid transaction: %v", errors))
		return
	}
	if err := th.repository.UpdateTransaction(id, transaction); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}

func (th *TransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := th.repository.DeleteTransaction(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
