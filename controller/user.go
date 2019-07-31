package controller

import (
	"Shepherd/model"
	"Shepherd/pkg/jwt"
	"Shepherd/pkg/response"
	"Shepherd/pkg/scrypt"
	"Shepherd/pkg/sms"
	"github.com/gin-gonic/gin"
)

// Binding from RegisterOrLogin JSON
type register struct {
	Mobile     string `json:"mobile" binding:"required"`
	SmsCode    string `json:"smsCode"`
	Password   string `json:"password"`
	InviteCode string `json:"inviteCode"`
}

// @Summary 用户注册 用户登录(密码) 用户登录（短信验证码）
// @Description register or login user by mobile,password,smsCode
// @Accept  json
// @Produce  json
// @Param mobile query string true "Mobile"
// @Param password query string false "Password"
// @Param smsCode query string false "SmsCode"
// @Param inviteCode query string false "InviteCode"
// @Success 200 {string} json "{"time": 1561513181, "code": 200, "msg": "成功", "data" : {}}"
// @Failure 400 {string} json "{"time": 1561513181, "code": 400, "msg": "msg"}"
// @Router /v1/registerOrLogin/ [post]
func RegisterOrLogin(c *gin.Context) {
	var param register
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	// 查询该手机号是否已经注册
	user, err := model.GetUserByMobile(param.Mobile)
	if err != nil {
		// 进行注册
		var scryptPassword string
		if len(param.Password) > 0 {
			// 如果存在密码，则加密存储
			scryptPassword = scrypt.GetScryptPasswordBase64(param.Password, param.Mobile)
		}
		// Verify SmsCode
		err = sms.VerifySMS(param.Mobile, param.SmsCode)
		if err != nil {
			response.ErrorWithMsg(c, err.Error())
			return
		}
		// Creat User DB
		user := new(model.User)
		user.Mobile = param.Mobile
		user.Password = scryptPassword
		if err := model.AddUser(user); err == nil {
			responseUserWithToken(c, user)
		} else {
			response.ErrorWithMsg(c, err.Error())
		}
	} else {
		// 进行登录
		if len(param.Password) > 0 {
			// login via password
			if user.Password != scrypt.GetScryptPasswordBase64(param.Password, param.Mobile) {
				response.ErrorWithMsg(c, "账户或密码错误")
				return
			}
			responseUserWithToken(c, user)
		} else if len(param.SmsCode) > 0 { // login via smsCode
			err = sms.VerifySMS(param.Mobile, param.SmsCode)
			if err != nil {
				response.ErrorWithMsg(c, err.Error())
				return
			}
			responseUserWithToken(c, user)
		} else {
			response.ErrorWithMsg(c, "参数缺少密码或短信验证码，无法登录")
		}
	}
}

func responseUserWithToken(c *gin.Context, user *model.User) {
	token, err := jwt.CreateToken(user.Id)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	user.Token = token
	response.JsonWithData(c, user)
}

// @Summary 获取用户
func GetUser(c *gin.Context) {
	userId, ok := c.Get(jwt.KeyUserId)
	if ok {
		user, err := model.GetUserById(userId.(uint))
		if err != nil {
			response.ErrorWithMsg(c, err.Error())
		} else {
			response.JsonWithData(c, user)
		}
	} else {
		response.ErrorWithMsg(c, "Token异常，用户不存在")
	}
}

// @Summary 获取用户
//func GetUsers(c *gin.Context) {
//	list := make([]*model.User, 0)
//	list, err := model.GetAllUser()
//	if err != nil {
//		helper.ErrorWithMsg(c, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, list)
//}

// @Summary 删除用户
//func DelUser(c *gin.Context) {
//	id := c.Param("id")
//	intId, err := strconv.Atoi(id)
//	if err != nil {
//		helper.ErrorWithMsg(c, "输入删除用户id非法")
//		return
//	}
//	err = model.DeleteUser(intId)
//	if err != nil {
//		helper.ErrorWithMsg(c, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, "ok")
//}
