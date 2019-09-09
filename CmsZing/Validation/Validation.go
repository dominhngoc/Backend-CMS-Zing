package Validation

import (
	"backend-cms-zing/CmsZing/conf"
	"github.com/astaxie/beego/validation"
)

func RebuildValidate(v interface{}) (bool, interface{}) {
	valid := validation.Validation{}
	if b, _ := valid.Valid(v); !b {

		// valid doesn't pass
		var detailErrorCode = make(map[string]int)
		for _, err := range valid.Errors {
			if err.Name == "Required" {
				detailErrorCode[err.Field] = conf.FIELD_REQUIRED
			}
			if err.Name == "MinSize" {
				detailErrorCode[err.Field] = conf.MIN_SIZE
			}
			if err.Name == "MaxSize" {
				detailErrorCode[err.Field] = conf.MAX_SIZE
			}
		}
		return false, detailErrorCode
	}
	return true, nil
}
