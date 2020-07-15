package denny

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/whatvn/denny/log"
	"net/http"
)

type controller interface {
	Handle(*Context)
	init()
	SetValidator(validator binding.StructValidator)
}

type Controller struct {
	binding.StructValidator
	*log.Log
}

func (c *Controller) init() {
	c.Log = log.New()
	c.StructValidator = binding.Validator
}

func (c *Controller) SetValidator(v binding.StructValidator) {
	c.StructValidator = v
}

type wrapController struct {
	Controller
	hf gin.HandlerFunc
}

func WrapGin(h gin.HandlerFunc) controller {
	return &wrapController{hf: h}
}

func (h *wrapController) Handle(ctx *Context) {
	h.hf(ctx)
}

func WrapH(h http.Handler) controller {
	return WrapGin(gin.WrapH(h))
}

func WrapF(h http.HandlerFunc) controller {
	return WrapGin(gin.WrapF(h))
}
