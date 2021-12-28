package util

import "github.com/gin-gonic/gin"

func InvalidParamsReponse(c *gin.Context) {
	c.JSON(int(InvalidParams), gin.H{
		"success": true,
		"msg":     GetMsg(int(InvalidParams)),
	})
	//c.Abort()
}
func ErrorAuthTokenReponse(c *gin.Context) {
	c.JSON(int(ErrorAuthToken), gin.H{
		"success": true,
		"msg":     GetMsg(int(ErrorAuthToken)),
	})
}

func InternalServerResponse(c *gin.Context) {
	c.JSON(int(ErrorInternalServer), gin.H{
		"success": true,
		"msg":     GetMsg(int(ErrorInternalServer)),
	})
	//c.Abort()
}
func SuccessResponse(c *gin.Context, result interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"rows":    result,
	})
}
func SuccessUpdateResponse(c *gin.Context, rowsCount int) {
	if rowsCount > 0 {
		c.JSON(200, gin.H{
			"success":      true,
			"rowsAffected": rowsCount,
		})
	} else {
		c.JSON(200, gin.H{
			"success": false,
		})
	}
}
func FailedResponse(c *gin.Context, message string, result interface{}) {
	c.JSON(200, gin.H{
		"success": false,
		"msg":     message,
		"rows":    result,
	})
}
