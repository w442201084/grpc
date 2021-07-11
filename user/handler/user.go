package handler

import (
	"context"
	"github.com/w442201084/grpc/domain/model"
	"github.com/w442201084/grpc/domain/service"
	pb "github.com/w442201084/grpc/proto/user"
	"strconv"
)

type Users struct {
	UserService service.IUserService
}

func (this *Users) Register(ctx context.Context,
	req *pb.UserRegisterRequest, resp *pb.UserRegisterResponse) error{
	userRegister := &model.User{
		UserName: req.UserName ,
		FirstName: req.FirstName ,
		Pwd: req.Pwd,
	}
	_, err := this.UserService.AddUser(userRegister)
	if nil != err {
		return err
	}
	resp.Message = "reg success ..."
	return nil
}

func (this *Users) Login(ctx context.Context,
	req *pb.UserLoginRequest, resp *pb.UserLoginResponse) error {
	isOk ,err := this.UserService.CheckPwd(req.UserName , req.Pwd)
	if nil != err {
		return err
	}

	resp.Message = "bool" + strconv.FormatBool(isOk)
	return nil
}

func (this *Users) GetUserInfo( ctx context.Context,
	req *pb.GetUserInfoRequest, resp *pb.GetUserInfoResponse) error {
	user , err := this.UserService.FindUserByUserName(req.UserName)
	if nil != err {
		return err
	}
	resp.UserName = user.UserName
	resp.FirstName = user.FirstName
	resp.UserId = user.Id
	return nil
}

