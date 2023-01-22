package http

import (
	"bank-api/internal/entity"
	"bank-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// @Tags        transaction
// @Description create transaction
// @Accept      json
// @Produce     json
// @Param       transactionParams body entity.CreateTransactionParams true "transaction info"
// @Success     201
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/transaction/create [post]
func (c Controller) createTransaction(ctx *gin.Context) {
	var transactionParams *entity.CreateTransactionParams
	if err := ctx.BindJSON(&transactionParams); err != nil {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidRequestBodyError)
		return
	}
	if transactionParams.Type == "" || transactionParams.ToAccountId == 0 || transactionParams.Amount == 0 {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidRequestBodyError)
		return
	}
	transactionParams.Type = strings.ToLower(transactionParams.Type)
	if !utils.IsValidTransactionType(transactionParams.Type) {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidTransactionTypeError)
		return
	}
	if transactionParams.Amount < 1 {
		errorResponse(ctx, http.StatusBadRequest, entity.TooSmallMoneyError)
		return
	}
	if transactionParams.ToAccountId < 1 {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidIdError)
		return
	}
	if transactionParams.Type == "transfer" {
		if transactionParams.FromAccountId < 1 {
			errorResponse(ctx, http.StatusBadRequest, entity.InvalidIdError)
			return
		} else if transactionParams.FromAccountId == transactionParams.ToAccountId {
			errorResponse(ctx, http.StatusBadRequest, entity.SameIdError)
			return
		}
	}
	err := c.service.InsertTransaction(transactionParams)
	if err != nil {
		if err == entity.ToClientNotFoundError || err == entity.FromClientNotFoundError || err == entity.DifferentCurrencyError || err == entity.NotEnoughMoneyError {
			errorResponse(ctx, http.StatusBadRequest, err)
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "transaction created"})
}

// @Tags        transaction
// @Description get transaction
// @Produce     json
// @Param       id query string true "transaction id"
// @Success     200
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/transaction [get]
func (c Controller) getTransaction(ctx *gin.Context) {
	id := ctx.Query("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidIdError)
		return
	}
	if idNum < 1 {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidIdError)
		return
	}
	transaction, err := c.service.GetTransaction(idNum)
	if err != nil {
		if err == entity.TransactionNotFoundError {
			errorResponse(ctx, http.StatusBadRequest, err)
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

// @Tags        transaction
// @Description get transactions list
// @Produce     json
// @Param       account_id query string false "account id"
// @Success     200
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/transaction/list [get]
func (c Controller) getTransactions(ctx *gin.Context) {
	var (
		idNum int64
		err   error
	)
	id := ctx.Query("account_id")
	if id != "" {
		idNum, err = strconv.ParseInt(id, 10, 64)
	}
	transactions, err := c.service.GetTransactions(idNum)
	if err != nil {
		if err == entity.AccountNotFoundError {
			errorResponse(ctx, http.StatusBadRequest, err)
			return
		}
		fmt.Println(err)
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}
