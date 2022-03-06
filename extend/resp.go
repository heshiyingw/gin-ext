package extend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Resp struct {
	Success bool        `json:"success" binding:"required" example:"true"` // 请求结果，失败:false，成功:true
	Msg     string      `json:"msg" binding:"required" example:"ok"`       // 请求结果的message
	Data    interface{} `json:"data,omitempty" binding:"required"`         // 返回的数据
	Code    int32       `json:"code"`                                      // code
}

func SendData(ctx *gin.Context, res Resp) {
	ctx.JSON(http.StatusOK, res)
}
func SendParamError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusBadRequest, Resp{
			Success: false,
			Msg:     err.Error(),
			Data:    nil,
			Code:    0,
		})
		return
	}
	errInfos := make([]string, 0, len(errs))
	for _, e := range errs {
		errInfo := e.Translate(Translator)
		errInfos = append(errInfos, fmt.Sprintf("%v:%v", e.Field(), errInfo))
	}
	ctx.JSON(http.StatusBadRequest, Resp{
		Success: false,
		Msg:     strings.Join(errInfos, ","),
		Data:    nil,
		Code:    0,
	})
}
