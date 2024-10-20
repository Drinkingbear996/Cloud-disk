package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"fmt"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	//1.Looking for user from database
	user := new(models.UserBasic)
	//打印
	fmt.Printf("MD5 password: %s\n", helper.Md5(req.Password))
	//Create helper.go MD 5 encode
	has, err := models.Engine.Where(" name = ? AND password = ?", req.Name, req.Password).Get(user)
	// 打印查询结果，检查是否找到了用户
	fmt.Printf("User found: %v, error: %v\n", has, err)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.New(" Username or password error")
	}
	//2.Generate token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	//测试token
	fmt.Println("Generated token: ", token)

	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	return
}
