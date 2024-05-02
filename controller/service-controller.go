package controller

import (
	"net/http"

	"github.com/fesTech-cloud/gin/entity"
	"github.com/fesTech-cloud/gin/service"
	"github.com/fesTech-cloud/gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	c.service.Save(video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

// func (c *controller) ShowAll(ctx *gin.Context) {
// 	videos := c.service.FindAll()

// 	data := gin.H{
// 		"title":  "Video Page",
// 		"videos": videos,
// 	}
// 	ctx.HTML(http.StatusOK, "index.html", data)
// }
