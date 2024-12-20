package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var auth model.AuthRequest
	err := validation.BindJsonAndValidate(c, &auth)
	if err != nil {
		return
	}

	err = h.userService.Register(c, auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: err.Error(),
				Code:    httpcommon.ErrorResponseCode.InternalServerError,
			},
		))
		return
	}
	c.AbortWithStatus(204)
}

func (h *UserHandler) Login(c *gin.Context) {
	var auth model.AuthRequest
	err := validation.BindJsonAndValidate(c, &auth)
	if err != nil {
		return
	}

	user, err := h.userService.Login(c, auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: err.Error(),
				Code:    httpcommon.ErrorResponseCode.InvalidRequest,
			},
		))
		return
	}
	c.JSON(200, httpcommon.NewSuccessResponse[entity.User](&entity.User{
		Username: user.Username,
	}))
}

func (h *UserHandler) WhoAmI(c *gin.Context) {
	userId, _ := c.Get("userId")

	c.JSON(200, gin.H{"userId": userId})
}
