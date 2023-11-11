package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todolist_app "todolist-app"
)

// @Summary      Create todo items
// @Security     ApiKeyAuth
// @Tags         items
// @Description  Create new items in a specific todo list
// @ID           create-items
// @Accept       json
// @Produce      json
// @Param        id     path    int                      true  "Todo List ID"
// @Param        input  body    todolist_app.TodoItem    true  "Item Info"
// @Success      200    {object} map[string]int          "ID of the created item"
// @Failure      400    {object} errorResponse           "Invalid list ID parameter or bad request data"
// @Failure      401    {object} errorResponse           "Authentication error"
// @Failure      404    {object} errorResponse           "Todo list not found"
// @Failure      500    {object} errorResponse           "Internal server error"
// @Router       /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todolist_app.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary      Get All Items
// @Security     ApiKeyAuth
// @Tags         items
// @Description  get all items for a specified list
// @ID           get-all-items
// @Accept       json
// @Produce      json
// @Param        listId  path  int  true  "List ID"
// @Success      200     {object}  []todolist_app.TodoItem  "List of Todo Items"
// @Failure      400     {object}  errorResponse            "Bad Request"
// @Failure      404     {object}  errorResponse            "Not Found"
// @Failure      500     {object}  errorResponse            "Internal Server Error"
// @Failure      default {object}  errorResponse            "Default Error"
// @Router       /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary      Get Items By Id
// @Security     ApiKeyAuth
// @Tags         items
// @Description  get item by id
// @ID           get-item-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  todolist_app.ListItem  "The requested Todo Item"
// @Failure      400  {object}  errorResponse           "Bad Request"
// @Failure      404  {object}  errorResponse           "Not Found"
// @Failure      500  {object}  errorResponse           "Internal Server Error"
// @Failure      default {object}  errorResponse        "Default Error"
// @Router       /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

// @Summary     Update Item
// @Security    ApiKeyAuth
// @Tags        items
// @Description update by id
// @ID          update-item-by-id
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Item ID"
// @Success     200  {object}  todolist_app.ListItem
// @Failure     400  {object}  errorResponse
// @Failure     404  {object}  errorResponse
// @Failure     500  {object}  errorResponse
// @Failure     default {object}  errorResponse
// @Router      /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todolist_app.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary      Delete Items
// @Security     ApiKeyAuth
// @Tags         items
// @Description  delete item by id
// @ID           delete-item-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  statusResponse          "Success Response"
// @Failure      400  {object}  errorResponse           "Bad Request"
// @Failure      404  {object}  errorResponse           "Not Found"
// @Failure      500  {object}  errorResponse           "Internal Server Error"
// @Failure      default {object}  errorResponse        "Default Error"
// @Router       /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
