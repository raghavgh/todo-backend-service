package validators

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/todo-backend-service/dto"
)

func init() {
	valid.TagMap["not_empty"] = func(str string) bool {
		return len(str) != 0
	}
	valid.TagMap["password"] = func(password string) bool {
		if valid.IsNull(password) {
			return false
		}
		if len(password) < 8 {
			return false
		}
		if !valid.StringMatches(password, `[A-Z]`) {
			return false
		}
		if !valid.StringMatches(password, `[a-z]`) {
			return false
		}
		if !valid.StringMatches(password, `[0-9]`) {
			return false
		}
		return true
	}
}

func IsValidSignupRequest(signupRequest *dto.SignupRequest) (bool, error) {
	isValid, err := valid.ValidateStruct(signupRequest)
	if err != nil {
		return false, nil
	}
	return isValid, nil
}

func IsValidLoginRequest(loginRequest *dto.LoginRequest) (bool, error) {
	isValid, err := valid.ValidateStruct(loginRequest)
	if err != nil {
		return false, nil
	}
	return isValid, nil
}
