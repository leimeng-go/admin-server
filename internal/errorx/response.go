package errorx

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Body 统一响应结构
type Body struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 错误信息
	Data    interface{} `json:"data"`    // 数据
}

// WriteSuccess 成功响应
func WriteSuccess(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, &Body{
		Code:    int(Success),
		Message: "success",
		Data:    data,
	})
}

// WriteError 错误响应
func WriteError(w http.ResponseWriter, err error) {
	if e, ok := err.(*Error); ok {
		httpx.WriteJson(w, e.GetHTTPCode(), &Body{
			Code:    int(e.Code),
			Message: e.Message,
		})
		return
	}

	// 处理其他错误
	httpx.WriteJson(w, http.StatusInternalServerError, &Body{
		Code:    int(ServerError),
		Message: err.Error(),
	})
}
