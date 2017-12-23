package controller

import "github.com/gin-gonic/gin"

type DocumentationController struct {
	r *gin.Engine
}

func DefaultDocumentationController(r *gin.Engine) *DocumentationController {
	dc := &DocumentationController{
		r: r,
	}

	dc.Default()
	return dc
}

func (dc *DocumentationController) Default() {
	dc.r.Static("/docs", "./content/docs/")
}
