package repository

import "github.com/w442201084/grpc/domain/model"
import "github.com/jinzhu/gorm"

type IUserRepository interface {

	InitTable () error

	// 查询用户信息
	FindUserByName (string)	( *model.User , error )

	// 根据ID查询用户信息
	FindUserById	(int64)	( *model.User , error )

	// 创建用户
	CreateUser (*model.User)	(int64 , error)

	// 删除用户
	DeleteUserById(int64)	(error)

	// 更新用户
	UpdateUser(*model.User) (error)
}


type UserRepository struct {
	mysqlDb *gorm.DB
}

func (this *UserRepository) InitTable() error {
	return this.mysqlDb.CreateTable(&model.User{}).Error
}

func (this *UserRepository) FindUserByName(userName string) ( *model.User , error ) {
	user := &model.User{}
	return user , this.mysqlDb.Where("user_name=?" , userName).Find(user).Error
}

func (this *UserRepository) FindUserById(id int64) (*model.User , error) {
	user := &model.User{}
	return user , this.mysqlDb.Where("id=?" , id).Find(user).Error
}

func (this *UserRepository) CreateUser(user *model.User) (int64 , error) {
	return user.Id , this.mysqlDb.Create(user).Error
}

func (this *UserRepository) DeleteUserById(id int64) (error) {
	return this.mysqlDb.Where("id = ?" , id).Delete(&model.User{}).Error
}

func (this *UserRepository) UpdateUser(user *model.User) (error) {
	return this.mysqlDb.Model(user).Update(&user).Error
}


func NewUserRepository(db *gorm.DB) IUserRepository{
	return &UserRepository{
		mysqlDb: db,
	}
}