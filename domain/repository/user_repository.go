package repository

import (
	"go-user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(name string) (*model.User, error)
	FindUserByID(id int64) (*model.User, error)
	CreatUser(user *model.User) (int64, error)
	DeleteUserByID(id int64) error
	UpdateUser(user *model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDBb: db}
}

type UserRepository struct {
	mysqlDBb *gorm.DB
}

//初始化表结构
func (r *UserRepository) InitTable() error {
	return r.mysqlDBb.AutoMigrate(&model.User{})
}

//根据用户名查找用户
func (r *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, r.mysqlDBb.Where("user_name = ?", name).Find(user).Error
}

//根据id查找用户
func (r *UserRepository) FindUserByID(id int64) (*model.User, error) {
	user := &model.User{}
	return user, r.mysqlDBb.First(user, id).Error
}

//创建用户
func (r *UserRepository) CreatUser(user *model.User) (int64, error) {
	return user.ID, r.mysqlDBb.Create(user).Error
}

//根据id删除用户
func (r *UserRepository) DeleteUserByID(id int64) error {
	return r.mysqlDBb.Where("id = ?", id).Delete(&model.User{}).Error
}

//更新用户信息
func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.mysqlDBb.Model(user).Updates(&user).Error
}

//查找所有
func (r *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, r.mysqlDBb.Find(&userAll).Error
}