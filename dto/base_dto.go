package dto

import (
	"encoding/json"
	"fmt"
)

type BaseResponse struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Operation  string        `json:"operation"`
	Errors     []interface{} `json:"errors"`
	Data       interface{}   `json:"data"`
	//Validate   map[string]string `json:"-"`
}

func NewBaseResponse(statusCode int) BaseResponse {
	return BaseResponse{
		StatusCode: statusCode,
	}
}

func (r BaseResponse) WithMessage(message string) BaseResponse {
	r.Message = message
	return r
}

func (r BaseResponse) WithData(data interface{}) BaseResponse {
	r.Data = data
	return r
}

func (r BaseResponse) WithValidate(validate map[string]string) BaseResponse {
	for k, v := range validate {
		//fmt.Printf("key[%s] value[%s]\n", k, v)
		r.Errors = append(r.Errors, k+" "+v)
		//r.Errors = k + " " + v
	}
	return r
}
func (r BaseResponse) WithError(errs interface{}) BaseResponse {
	//r.Errors = errs
	return r
}
func (r BaseResponse) WithOperation(op string) BaseResponse {
	r.Operation = op
	return r
}

func (r BaseResponse) ToJsonWithByte() []byte {
	strJson, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}
	return strJson
}
func (r BaseResponse) ToJson() string {
	strJson, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return r.Message
	}
	return string(strJson)
}
