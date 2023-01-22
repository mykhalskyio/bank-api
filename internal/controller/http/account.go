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

// @Tags        account
// @Description create account
// @Accept      json
// @Produce     json
// @Param       accountParams body entity.CreateAccountParams true "account info"
// @Success     201
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/account/create [post]
func (c Controller) createAccount(ctx *gin.Context) {
	var accountParams *entity.CreateAccountParams
	if err := ctx.BindJSON(&accountParams); err != nil {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidRequestBodyError)
		return
	}
	if accountParams.OwnerEmail == "" || accountParams.Currency == "" {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidRequestBodyError)
		return
	}
	accountParams.OwnerEmail = strings.ToLower(accountParams.OwnerEmail)
	if !utils.IsValidEmail(accountParams.OwnerEmail) {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidEmailError)
		return
	}
	accountParams.Currency = strings.ToUpper(accountParams.Currency)
	if !utils.IsValidCurrency(accountParams.Currency) {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidCurrencyError)
		return
	}
	err := c.service.InsertAccount(accountParams)
	if err != nil {
		if err == entity.ClientNotFoundError {
			errorResponse(ctx, http.StatusBadRequest, err)
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("account created for %s", accountParams.OwnerEmail)})
}

// @Tags        account
// @Description get account
// @Produce     json
// @Param       id query int true "account id"
// @Success     200
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /api/account [get]
func (c Controller) getAccount(ctx *gin.Context) {
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
	account, err := c.service.GetAccount(idNum)
	if err != nil {
		if err == entity.AccountNotFoundError {
			errorResponse(ctx, http.StatusNotFound, err)
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// @Tags        account
// @Description get accounts list
// @Produce     json
// @Param       user_email query string false "user email"
// @Success     200
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/account/list [get]
func (c Controller) getAccounts(ctx *gin.Context) {
	email := ctx.Query("user_email")
	if email != "" {
		if !utils.IsValidEmail(email) {
			errorResponse(ctx, http.StatusBadRequest, entity.InvalidEmailError)
			return
		}
	}
	email = strings.ToLower(email)
	accounts, err := c.service.GetAccounts(email)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
