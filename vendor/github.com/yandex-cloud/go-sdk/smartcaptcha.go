package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/smartcaptcha"
)

const (
	SmartCaptchaServiceID Endpoint = "smart-captcha"
)

// SmartCaptcha returns SmartCaptcha object that is used to operate with captchas
func (sdk *SDK) SmartCaptcha() *smartcaptcha.SmartCaptcha {
	return smartcaptcha.NewSmartCaptcha(sdk.getConn(SmartCaptchaServiceID))
}
