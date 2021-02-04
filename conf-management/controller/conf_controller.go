package controller

import (
	"config-center/conf-core/exception"
	"config-center/conf-management/controller/model"
	"config-center/conf-management/controller/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func init() {

}

type ProfileController struct {
	storageService service.ProfileService
}

func (c *ProfileController) Router(engine *gin.Engine) {
	confEngine := engine.Group("/v1/conf")
	{
		confEngine.GET("/:id", c.getConf)
		confEngine.POST("", c.addConf)
		confEngine.PUT("", c.updateConf)
		confEngine.DELETE("/:id", c.deleteConf)
		confEngine.GET("/query", c.queryConf)
	}
}

func (c *ProfileController) getConf(ctx *gin.Context) {
	id := ctx.Param("id")
	p, err := c.storageService.GetProfileById(id)
	if err != nil {
		ctx.JSON(err.(interface{}).(exception.ErrorCode).GetStatus(), gin.H{
			"success": false,
		})

		return
	}

	ctx.JSON(200, model.ProfileDto{}.From(p))
}

func (c *ProfileController) addConf(ctx *gin.Context) {

	var requestBody model.AddProfileBody
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(400, exception.INTERVAL_ERROR)
		return
	}

	p := requestBody.ToProfile()
	p.SetId(uuid.NewV4().String())
	err = c.storageService.AddProfile(*p, "admin")
	if err != nil {
		ctx.JSON(err.(interface{}).(exception.ErrorCode).GetStatus(), err)

		return
	}

	p, err = c.storageService.GetProfileById(p.GetId())
	if err != nil {
		ctx.JSON(err.(interface{}).(exception.ErrorCode).GetStatus(), err)

		return
	}

	ctx.JSON(201, model.ProfileDto{}.From(p))
}

func (c *ProfileController) updateConf(ctx *gin.Context) {


}

func (c *ProfileController) deleteConf(ctx *gin.Context) {

}

func (c *ProfileController) queryConf(ctx *gin.Context) {

}
