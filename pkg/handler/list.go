package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todolist_app "todolist-app"
)

//	@Summary		Create todo list
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	create todolist_app list
//	@ID				create-list
//	@Accept			json
//	@Produce		json
//	@Param			input	body		todolist_app.TodoList	true	"list info"
//	@Success		200		{integer}	integer			1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists [post]

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todolist_app.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todolist_app.TodoList `json:"data"`
}

// @Summary		Get All Lists
// @Security		ApiKeyAuth
// @Tags			lists
// @Description	get all lists
// @ID				get-all-lists
// @Accept			json
// @Produce		json
// @Success		200		{object}	getAllListsResponse
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary		Get List By Id
// @Security		ApiKeyAuth
// @Tags			lists
// @Description	get list by id
// @ID				get-list-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	todolist_app.ListItem
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

// @Summary		Update List
// @Security		ApiKeyAuth
// @Tags			lists
// @Description		update by id
// @ID				update-list-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	todolist_app.ListItem
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists/:id [post]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todolist_app.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary		Delete List
// @Security		ApiKeyAuth
// @Tags			lists
// @Description		delete list by id
// @ID				delete-list-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	todolist_app.ListItem
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
