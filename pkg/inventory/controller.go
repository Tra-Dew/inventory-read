package inventory

import (
	"net/http"

	"github.com/d-leme/tradew-inventory-read/pkg/core"
	"github.com/gin-gonic/gin"
)

// Controller ...
type Controller struct {
	authenticate *core.Authenticate
	settings     *core.Settings
	repository   Repository
}

// NewController ...
func NewController(authenticate *core.Authenticate, repository Repository) Controller {
	return Controller{
		authenticate: authenticate,
		repository:   repository,
	}
}

// RegisterRoutes ...
func (c *Controller) RegisterRoutes(r *gin.RouterGroup) {
	inventory := r.Group("/inventory-read")
	{
		inventory.GET("", c.get)
		inventory.GET("my-items", c.authenticate.Middleware(), c.getMyItems)
		inventory.GET(":id", c.authenticate.Middleware(), c.getByID)
	}
}

func (c *Controller) get(ctx *gin.Context) {
	req := new(GetItemsRequest)

	if err := ctx.ShouldBindQuery(req); err != nil {
		core.HandleRestError(ctx, core.ErrMalformedJSON)
		return
	}

	res, err := c.repository.Get(ctx, req)

	if err != nil {
		core.HandleRestError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) getMyItems(ctx *gin.Context) {
	req := new(GetItemsRequest)
	userID := ctx.GetString("user_id")

	if err := ctx.ShouldBindQuery(req); err != nil {
		core.HandleRestError(ctx, core.ErrMalformedJSON)
		return
	}

	res, err := c.repository.GetUserItems(ctx, userID, req)

	if err != nil {
		core.HandleRestError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) getByID(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	id := ctx.Param("id")

	res, err := c.repository.GetByID(ctx, userID, id)

	if err != nil {
		core.HandleRestError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
