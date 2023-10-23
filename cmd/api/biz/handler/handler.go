package handler

import (
	"github.com/Yra-A/Fusion_Go/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func BadResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})

}
func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, data)
}
