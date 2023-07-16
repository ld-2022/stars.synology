package action

import "encoding/json"

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *R) Ok(data interface{}) *R {
	r.Code = 0
	r.Data = data
	return r
}

func (r *R) OkMsg(message string, data interface{}) *R {
	r.Code = 0
	r.Message = message
	r.Data = data
	return r
}

func (r *R) Error(message string) *R {
	r.Code = 500
	r.Message = message
	return r
}

func (r *R) NotLogin() *R {
	r.Code = 1002
	r.Message = "未登录"
	return r
}

func (r *R) ToJsonBytes() []byte {
	bytes, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return bytes
}
