package controllers

import "reflect"

type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type DataListResp struct {
	BaseResp
	Data interface{} `json:"data"`
}

type DataObjectResp struct {
	BaseResp
	Data interface{} `json:"data"`
}

func NewBaseResp(code int, message string) *BaseResp {
	var b BaseResp
	b.Code = code
	b.Message = message
	return &b
}
func (b *BaseResp) Set(code int, message string) {
	b.Code = code
	b.Message = message
}
func NewObjectResp(code int, message string, data interface{}) *DataObjectResp {
	var dor DataObjectResp
	dor.Code = code
	dor.Message = message

	if isNotNull(data) {
		dor.Data = data
	} else {
		dor.Data = map[string]interface{}{}
	}
	return &dor
}

func (d *DataObjectResp) Set(code int, message string, data interface{}) {
	d.Code = code
	d.Message = message

	if isNotNull(data) {
		d.Data = data
	} else {
		d.Data = map[string]interface{}{}
	}
}

func NewListResp(code int, message string, data interface{}) *DataListResp {
	var dlr DataListResp
	dlr.Code = code
	dlr.Message = message

	if isNotNull(data) {
		dlr.Data = data
	} else {

		dlr.Data = []interface{}{}
	}
	return &dlr
}

func isNotNull(i interface{}) bool {
	if i == nil {
		return false
	}
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Slice:
		if v.Len() > 0 {
			return true
		}
	case reflect.Map:
		if len(v.MapKeys()) > 0 {
			return true
		}
	default:
		return true
	}
	return false
}
