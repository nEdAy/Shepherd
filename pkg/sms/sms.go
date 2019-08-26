package sms

import (
	"github.com/nEdAy/Shepherd/pkg/config"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
)

var (
	InvalidJson       = errors.New("invalid json")
	AppkeyIsNull      = errors.New("AppKey为空")
	AppkeyIsIllegal   = errors.New("AppKey无效")
	PhoneOrCodeIsNull = errors.New("国家代码或手机号码为空")
	PhoneIsIllegal    = errors.New("手机号码格式错误")
	CodeIsNull        = errors.New("请求校验的验证码为空")
	IsFrequentChecks  = errors.New("请求校验验证码频繁（5分钟内同一个号码最多只能校验三次）")
	CodeIsError       = errors.New("验证码错误")
	ConfigIsError     = errors.New("没有打开服务端验证开关")
)

func VerifySMS(phone string, code string) error {
	parameter := "appkey=" + config.Mob.AppKey + "&phone=" + phone + "&zone=" + "86" + "&code=" + code
	response, err := resty.R().
		SetBody(parameter).
		Post("https://webapi.sms.mob.com/sms/verify")
	if err != nil {
		return err
	} else {
		if !gjson.Valid(response.String()) {
			return InvalidJson
		}
		status := gjson.Get(response.String(), "status").Int()
		switch status {
		case 200:
			return nil
		case 405:
			return AppkeyIsNull
		case 406:
			return AppkeyIsIllegal
		case 457:
			return PhoneOrCodeIsNull
		case 456:
			return PhoneIsIllegal
		case 466:
			return CodeIsNull
		case 467:
			return IsFrequentChecks
		case 468:
			return CodeIsError
		case 474:
			return ConfigIsError
		default:
			return errors.New("校验验证码其他错误")
		}
	}
}
