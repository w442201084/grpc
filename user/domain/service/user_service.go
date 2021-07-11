package service

import (
	"github.com/w442201084/grpc/common"
	"github.com/w442201084/grpc/domain/model"
	"github.com/w442201084/grpc/domain/repository"
)

type IUserService interface {

	AddUser (user *model.User) (int64 , error)

	DeleteUser(int64)(error)

	UpdateUser(user *model.User , isChangePwd bool) (err error)

	FindUserByUserName( userName string ) ( *model.User , error )

	CheckPwd(userName string , pwd string) (isOk bool , err error)

}

type UserService struct {
	userRepository repository.IUserRepository
}


func (this *UserService) CheckPwd(userName string ,
	pwd string) (isOk bool , err error) {
	user , err := this.FindUserByUserName(userName)
	if nil != err {
		return false , err
	}
	return common.ValidateUserPassword(pwd , user.Pwd)
}

func (this *UserService) FindUserByUserName( userName string ) ( *model.User , error ) {
	return this.userRepository.FindUserByName(userName)
}

func (this *UserService) UpdateUser(user *model.User ,
	isChangePwd bool) ( err error) {
	if isChangePwd {
		pwdByte ,err := common.GeneratorPassword(user.Pwd)
		if nil != err {
			return err
		}
		user.Pwd = string(pwdByte)
	}
	return  this.userRepository.UpdateUser(user)
}


func (this *UserService) DeleteUser(userId int64) error {
	return this.userRepository.DeleteUserById(userId)
}

func (this *UserService) AddUser (user *model.User) (int64 , error) {
	pwdByte , err := common.GeneratorPassword(user.Pwd)
	if nil != err {
		return user.Id , err
	}

	user.Pwd = string(pwdByte)
	return this.userRepository.CreateUser(user)
}



func NewUserService(userRepository repository.IUserRepository) *UserService{
	return &UserService{
		userRepository: userRepository ,
	}
}