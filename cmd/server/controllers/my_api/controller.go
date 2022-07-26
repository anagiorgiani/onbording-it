package controller

import (
	"fmt"
	"net/http"

	"github.com/anagiorgiani/onbording-it/internal/my_api/domain"
	"github.com/anagiorgiani/onbording-it/pkg/web"
	"github.com/gin-gonic/gin"
)

type apiController struct {
	apiService domain.ApiService
}

func NewController(s domain.ApiService) *apiController {
	return &apiController{
		apiService: s,
	}
}

func (c *apiController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		data := ctx.Query("data")
		fmt.Println("Dados", data)
		if data == "" {
			obj, err := c.apiService.GetCurrencies("usd-brl", "brl-usd", "ars-brl", "eur-brl")
			if err != nil {
				ctx.JSON(http.StatusPartialContent, nil)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"response": obj,
			})
			return
		}
		returnData, err := c.apiService.ProcessData(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, returnData, ""))
	}
}
