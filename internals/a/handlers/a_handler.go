package handlers

import (
	"microservice/internals/a/adapters/inputs"
	"microservice/internals/a/adapters/outputs"
	"microservice/internals/a/usecases"
	"microservice/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Get(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	usecase usecases.IUsecase
}

func NewHandler(usecase usecases.IUsecase) IHandler {
	return &Handler{usecase: usecase}
}

// @summary Get
// @description Get user
// @tags a
// // @Security ApiKeyAuth
// @id Get
// // @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @response 200 {object} responses.SuccessOKSwagger "Get successfully"
// @response 201 {object} responses.SuccessCreatedSwagger "Create Successfully"
// @response 202 {object} responses.SuccessAcceptedSwagger "Accepted successfully"
// @response 400 {object} responses.ErrorBadRequestedSwagger "Bad Request"
// @response 401 {object} responses.ErrorUnauthorizedSwagger "Unauthorized"
// @response 400 {object} responses.ErrorNotFoundSwagger "Not Found"
// @response 422 {object} responses.ErrorValidatedSwagger "Unprocessable Entity"
// @response 500 {object} responses.ErrorInternalServerErrorSwagger "Internal Server Error"
// @router /api/v1/a/get [GET]
func (h *Handler) Get(c *gin.Context) {
	output, err := h.usecase.Get()
	if err != nil {
		c.JSON(responses.ErrorInternalServer(err, responses.NoDetail{}))
		return
	}

	c.JSON(responses.SuccessOK(output))
}

// @summary Add
// @description Add user
// @tags a
// // @Security ApiKeyAuth
// @id Add
// // @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param body body inputs.AddInput true "user data"
// @response 200 {object} responses.SuccessOKSwagger "Get successfully"
// @response 201 {object} responses.SuccessCreatedSwagger "Create Successfully"
// @response 202 {object} responses.SuccessAcceptedSwagger "Accepted successfully"
// @response 400 {object} responses.ErrorBadRequestedSwagger "Bad Request"
// @response 401 {object} responses.ErrorUnauthorizedSwagger "Unauthorized"
// @response 400 {object} responses.ErrorNotFoundSwagger "Not Found"
// @response 422 {object} responses.ErrorValidatedSwagger "Unprocessable Entity"
// @response 500 {object} responses.ErrorInternalServerErrorSwagger "Internal Server Error"
// @router /api/v1/a/add [post]
func (h *Handler) Add(c *gin.Context) {
	var input inputs.AddInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(responses.ErrorBadRequested(err, responses.NoDetail{}))
		return
	}
	if err := h.usecase.Add(input); err != nil {
		c.JSON(responses.ErrorInternalServer(err, responses.NoDetail{}))
		return
	}
	c.JSON(responses.SuccessCreated(responses.NoData{}))

}

// @summary Update
// @description Update user
// @tags a
// // @Security ApiKeyAuth
// @id Update
// // @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "id"
// @response 200 {object} responses.SuccessOKSwagger "Get successfully"
// @response 201 {object} responses.SuccessCreatedSwagger "Create Successfully"
// @response 202 {object} responses.SuccessAcceptedSwagger "Accepted successfully"
// @response 400 {object} responses.ErrorBadRequestedSwagger "Bad Request"
// @response 401 {object} responses.ErrorUnauthorizedSwagger "Unauthorized"
// @response 400 {object} responses.ErrorNotFoundSwagger "Not Found"
// @response 422 {object} responses.ErrorValidatedSwagger "Unprocessable Entity"
// @response 500 {object} responses.ErrorInternalServerErrorSwagger "Internal Server Error"
// @router /api/v1/a/update/{id} [PUT]
func (h *Handler) Update(c *gin.Context) {
	if c.Param("id") == "" {
		output := []outputs.ValidateOutput{}
		output = append(output, outputs.ValidateOutput{
			Field:   "id",
			Message: "โปรดระบุ ID",
		})
		c.JSON(responses.ErrorValidated(output))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(responses.ErrorInternalServer(err, responses.NoDetail{}))
		return
	}

	input := inputs.UpdateInput{
		ID: id,
	}

	if err := h.usecase.Update(input); err != nil {
		c.JSON(responses.ErrorInternalServer(err, responses.NoDetail{}))
		return
	}

	c.JSON(responses.SuccessAccepted(responses.NoData{}))
}

// @summary Delete
// @description Delete user
// @tags a
// // @Security ApiKeyAuth
// @id Delete
// // @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "id"
// @response 200 {object} responses.SuccessOKSwagger "Get successfully"
// @response 201 {object} responses.SuccessCreatedSwagger "Create Successfully"
// @response 202 {object} responses.SuccessAcceptedSwagger "Accepted successfully"
// @response 400 {object} responses.ErrorBadRequestedSwagger "Bad Request"
// @response 401 {object} responses.ErrorUnauthorizedSwagger "Unauthorized"
// @response 400 {object} responses.ErrorNotFoundSwagger "Not Found"
// @response 422 {object} responses.ErrorValidatedSwagger "Unprocessable Entity"
// @response 500 {object} responses.ErrorInternalServerErrorSwagger "Internal Server Error"
// @router /api/v1/a/delete/{id} [DELETE]
func (h *Handler) Delete(c *gin.Context) {

	if c.Param("id") == "" {
		output := []outputs.ValidateOutput{}
		output = append(output, outputs.ValidateOutput{
			Field:   "id",
			Message: "โปรดระบุ ID",
		})
		c.JSON(responses.ErrorValidated(output))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(responses.ErrorInternalServer(err, responses.NoDetail{}))
		return
	}

	input := inputs.DeleteInput{
		ID: id,
	}
	input.ID = id

	if err := h.usecase.Delete(input); err != nil {
		c.JSON(responses.ErrorInternalServer(err, responses.NoDetail{}))
		return
	}

	c.JSON(responses.SuccessAccepted(responses.NoData{}))
}
