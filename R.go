package R

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author yNsLuHan
 * @Description: Response构造体
 */
type Response struct {
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

/**
 * @Author yNsLuHan
 * @Description: res整合
 * @Time 2021-06-08 15:32:40
 * @param c
 * @param code
 * @param message
 * @param data
 */
func R(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

/**
 * @Author: yNsLuHan
 * @Description: res整合
 * @Time: 2021-06-09 22:31:31
 * @param: c
 * @param: code
 * @param: message
 * @param: data
 */
func Res(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	return
}

/**
 * @Author: yNsLuHan
 * @Description: err
 * @Time: 2021-06-09 22:31:39
 * @param: c
 * @param: a
 */
func Error(c *gin.Context, a ...interface{}) {
	if len(a) == 0 {
		c.JSON(http.StatusOK, Response{
			Code:    0,
			Message: "error",
			Data:    nil,
		})
		return
	} else if len(a) == 1 {
		c.JSON(http.StatusOK, Response{
			Code:    0,
			Message: a[0],
			Data:    nil,
		})
		return
	} else if len(a) == 2 {
		c.JSON(http.StatusOK, Response{
			Code:    0,
			Message: a[1],
			Data:    a[2],
		})
		return
	}
}

/**
 * @Author: yNsLuHan
 * @Description: ok
 * @Time: 2021-06-09 22:31:43
 * @param: c
 * @param: a
 */
func Ok(c *gin.Context, a ...interface{}) {
	if len(a) == 0 {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "success",
			Data:    nil,
		})
		return
	} else if len(a) == 1 {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "success",
			Data:    a[0],
		})
		return
	} else if len(a) == 2 {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: a[0],
			Data:    a[1],
		})
		return
	} else if len(a) == 3 {
		c.JSON(http.StatusOK, Response{
			Code:    a[0],
			Message: a[1],
			Data:    a[2],
		})
		return
	}
	return
}
