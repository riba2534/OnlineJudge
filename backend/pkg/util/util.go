package util

import "github.com/riba2534/OnlineJudge/backend/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
