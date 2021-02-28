package service

import (
	"errors"
	"go-user/domain/model"
	"go-user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(id int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(name string) (*model.User, error)
	CheckPwd(name string, pwd string) (isOk bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

//加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//验证密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

//添加用户
func (uds *UserDataService) AddUser(user *model.User) (id int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return uds.UserRepository.CreatUser(user)
}

//删除用户
func (uds *UserDataService) DeleteUser(id int64) error {
	return uds.UserRepository.DeleteUserByID(id)
}

//更新用户信息
func (uds *UserDataService) UpdateUser(user *model.User, isChangePwd bool) error {
	//如果用户更改了密码，先加密
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return uds.UserRepository.UpdateUser(user)
}

//根据用户名查找用户
func (uds *UserDataService) FindUserByName(name string) (*model.User, error) {
	return uds.UserRepository.FindUserByName(name)
}

//对比账号密码是否正确
func (uds *UserDataService) CheckPwd(name string, pwd string) (isOk bool, err error) {
	user, err := uds.UserRepository.FindUserByName(name)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}