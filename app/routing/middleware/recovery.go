package middleware

import (
	"github.com/EventRegistration/app/customerrors"
	"github.com/EventRegistration/app/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Recovery(ctx *gin.Context){
	defer func(ctx *gin.Context) {
		if rec := recover(); rec != nil {
			err := "Error unhandled route being hit" //util
			errorCode := "UNHANDLED_ROUTE"

			cusErr := customerrors.NewCusError(ctx,errorCode,err)
			response := models.Base{}
			response.Success = false
			response.Error =
			ctx.JSON(http.StatusInternalServerError, )
		}
	}(ctx)

	ctx.Next()
}