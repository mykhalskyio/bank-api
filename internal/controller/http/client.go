package http

import (
	"bank-api/internal/entity"
	"bank-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @Tags        client
// @Description create client
// @Accept      json
// @Produce     json
// @Param       clientParams body entity.CreateClientParams true "client info"
// @Success     201
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/client/create [post]
func (c Controller) createClient(ctx *gin.Context) {
	var clientParams *entity.CreateClientParams
	if err := ctx.BindJSON(&clientParams); err != nil {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidRequestBodyError)
		return
	}
	if clientParams.UserEmail == "" || clientParams.Username == "" {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidRequestBodyError)
		return
	}
	clientParams.UserEmail = strings.ToLower(clientParams.UserEmail)
	fmt.Println(utils.IsValidEmail(clientParams.UserEmail))
	if !utils.IsValidEmail(clientParams.UserEmail) {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidEmailError)
		return
	}
	err := c.service.InsertClient(clientParams)
	if err != nil {
		if err == entity.EmailInUseError {
			errorResponse(ctx, http.StatusBadRequest, err)
			return
		} else {
			fmt.Println(err)
			errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
			return
		}
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "client created"})
}

// @Tags        client
// @Description get client
// @Produce     json
// @Param       user_email query string true "user email"
// @Success     200
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /api/client [get]
func (c Controller) getClient(ctx *gin.Context) {
	email := ctx.Query("user_email")
	if !utils.IsValidEmail(email) {
		errorResponse(ctx, http.StatusBadRequest, entity.InvalidEmailError)
		return
	}
	email = strings.ToLower(email)
	client, err := c.service.GetClient(email)
	if err != nil {
		if err == entity.ClientNotFoundError {
			errorResponse(ctx, http.StatusNotFound, err)
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusOK, client)
}

// @Tags        client
// @Description get clients list
// @Produce     json
// @Success     200
// @Failure     500 {object} response
// @Router      /api/client/list [get]
func (c Controller) getClients(ctx *gin.Context) {
	clients, err := c.service.GetClients()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, entity.DatabaseError)
		return
	}
	ctx.JSON(http.StatusOK, clients)
}
