package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todolist_app "todolist-app"
)

type getAllListsResponse struct {
	Data []todolist_app.TodoList `json:"data"`
}

// @Summary      Create todo list
// @Security     ApiKeyAuth
// @Tags         lists
// @Description  Create a new todo list
// @ID           create-list
// @Accept       json
// @Produce      json
// @Param        input  body      todolist_app.TodoList  true  "List Info"
// @Success      200    {object}  map[string]int        "id"
// @Failure      400    {object}  errorResponse         "Invalid input"
// @Failure      404    {object}  errorResponse         "Not Found"
// @Failure      500    {object}  errorResponse         "Internal Server Error"
// @Failure      default {object}  errorResponse         "Unexpected error"
// @Router       /api/lists [post]
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

// @Summary      Get All Lists
// @Security     ApiKeyAuth
// @Tags         lists
// @Description  Retrieve all todo lists created by the authenticated user
// @ID           get-all-lists
// @Accept       json
// @Produce      json
// @Success      200 {object} getAllListsResponse "List of all todo lists"
// @Failure      401 {object} errorResponse      "Authentication error"
// @Failure      500 {object} errorResponse      "Internal server error"
// @Router       /api/lists [get]
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

// @Summary      Get List By Id
// @Security     ApiKeyAuth
// @Tags         lists
// @Description  Retrieve a specific todo list by its ID
// @ID           get-list-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo List ID"
// @Success      200  {object}  todolist_app.ListItem "Details of the specific todo list"
// @Failure      400  {object}  errorResponse         "Invalid ID parameter"
// @Failure      401  {object}  errorResponse         "Authentication error"
// @Failure      404  {object}  errorResponse         "Todo list not found"
// @Failure      500  {object}  errorResponse         "Internal server error"
// @Router       /api/lists/{id} [get]
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

// @Summary      Update List
// @Security     ApiKeyAuth
// @Tags         lists
// @Description  Update a specific todo list by its ID
// @ID           update-list-by-id
// @Accept       json
// @Produce      json
// @Param        id    path      int                        true "Todo List ID"
// @Param        input body      todolist_app.UpdateListInput true "Update data"
// @Success      200   {object}  statusResponse             "List updated successfully"
// @Failure      400   {object}  errorResponse              "Invalid request parameters"
// @Failure      401   {object}  errorResponse              "Authentication error"
// @Failure      404   {object}  errorResponse              "Todo list not found"
// @Failure      500   {object}  errorResponse              "Internal server error"
// @Router       /api/lists/{id} [put]
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

// @Summary      Delete List
// @Security     ApiKeyAuth
// @Tags         lists
// @Description  Delete a specific todo list by its ID
// @ID           delete-list-by-id
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "Todo List ID"
// @Success      200   {object}  statusResponse "List deleted successfully"
// @Failure      400   {object}  errorResponse  "Invalid ID parameter"
// @Failure      401   {object}  errorResponse  "Authentication error"
// @Failure      404   {object}  errorResponse  "Todo list not found"
// @Failure      500   {object}  errorResponse  "Internal server error"
// @Router       /api/lists/{id} [delete]
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
