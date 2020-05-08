package customerrors

import "github.com/gin-gonic/gin"

type cusError struct {
	err                error
	errorData          []map[string]interface{}
	errorCode          string
	subCode            string
	logData            map[string]interface{}
	errorType          string
	ctx                *gin.Context
}

func NewCusError(ctx *gin.Context, code string, err error) *cusError {
	cusErr := &cusError{
		err: 			err,
		errorCode:		code,
		logData:		map[string]interface{}{},
		ctx:			ctx,
	}
	return  cusErr
}