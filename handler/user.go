package handler

import (
	"context"
	"go-user/domain/model"
	"go-user/domain/service"
	user "go-user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

//注册
func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest,
	userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName: userRegisterRequest.UserName,
		FirstName: userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_,err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "注册成功"
	return nil
}

//登录
func (u *User) Login(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

//查询用户信息
func (u *User) GetUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(userInfo)

	return nil
}

//类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID

	return response
}
