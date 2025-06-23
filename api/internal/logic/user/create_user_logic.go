package user

import (
	"admin-server/api/internal/errorx"
	"admin-server/api/internal/model/user"
	"context"

	"golang.org/x/crypto/bcrypt"

	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// 1. 检查用户是否存在
	_, err = l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName)
	if err == nil {
		return nil, errorx.ErrUserExist
	}

	// 2. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 3. 构造用户实体
	newUser := user.User{
		UserName: req.UserName,
		Password: string(hashedPassword),
		Email:    req.Email,
		Mobile:   req.Mobile,
		NickName: req.Nickname,
		Avatar:   req.Avatar,
		RoleId:   req.RoleID,
		Status:   user.UserStatusNormal, // 使用预定义常量
	}
	// 4. 插入数据库
	_, err = l.svcCtx.UserModel.Insert(l.ctx, nil, &newUser)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateUserResp{}
	return resp, nil
}
