package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoService service.TodoService
	userService service.UserService
}

func NewTodoHandler(todoService service.TodoService, userService service.UserService) *TodoHandler {
	return &TodoHandler{todoService: todoService, userService: userService}
}

func (handler *TodoHandler) Add(c *gin.Context) {
	var req *model.TodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Invalid token", Field: "userId", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	req.UserId = userId.(int64)
	newTodo, err := handler.todoService.AddNewTodo(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[*entity.Todo](&newTodo))
}

func (handler *TodoHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Missing param id", Field: "id", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	todoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "id", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	var req *model.TodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Invalid token", Field: "userId", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	req.UserId = userId.(int64)
	updatedUser, err := handler.todoService.UpdateTodo(c, req, todoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[*entity.Todo](&updatedUser))
}

func (handler *TodoHandler) GetList(c *gin.Context) {
	filter := c.Query("filter")
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "Invalid token", Field: "userId", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}

	todoList, err := handler.todoService.GetListTodo(c, userId.(uint64), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[[]entity.Todo](&todoList))
}
