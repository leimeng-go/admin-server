package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"admin-server/internal/errorx"
	"admin-server/internal/model"
	"admin-server/internal/svc"
	"admin-server/internal/types"
	"admin-server/internal/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 用户注册
func (l *UserLogic) Register(req *types.RegisterReq) error {
	// 检查用户名是否已存在
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err == nil {
		return fmt.Errorf("用户名已存在")
	}

	// 检查邮箱是否已存在
	if req.Email != "" {
		_, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
		if err == nil {
			return fmt.Errorf("邮箱已存在")
		}
	}

	// 检查手机号是否已存在
	if req.Mobile != "" {
		_, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
		if err == nil {
			return fmt.Errorf("手机号已存在")
		}
	}

	// 验证密码
	if req.Password != req.ConfirmPassword {
		return fmt.Errorf("两次输入的密码不一致")
	}

	// 生成密码哈希
	h := md5.New()
	h.Write([]byte(req.Password))
	passwordHash := hex.EncodeToString(h.Sum(nil))

	// 生成雪花ID
	id := l.svcCtx.Snowflake.NextID()

	// 创建用户
	user := &model.User{
		Id:       uint64(id),
		Username: req.Username,
		Password: passwordHash,
		Nickname: req.Nickname,
		Avatar:   "https://avatars.githubusercontent.com/u/1?v=4",
		Email:    req.Email,
		Mobile:   req.Mobile,
		Role:     "user",
		Status:   model.UserStatusNormal,
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}

	return nil
}

// Login 用户登录
func (l *UserLogic) Login(req *types.LoginReq) (*types.CommonResp, error) {
	// 查找用户
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, errorx.ErrUserNotFound
	}

	// 验证密码
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(req.Password))) {
		return nil, errorx.ErrPasswordError
	}

	// 检查用户状态
	if user.Status == 0 {
		return nil, errorx.ErrUserDisabled
	}

	// 生成 token
	// token, err := l.svcCtx.Auth.GenerateToken(user.Id, user.Username, user.Role)
	// if err != nil {
	// 	return nil, errorx.ErrServerError
	// }

	return &types.CommonResp{
		Code:    0,
		Message: "登录成功",
		Data: map[string]string{
			"token": "",
		},
	}, nil
}

// GetUserInfo 获取用户信息
func (l *UserLogic) GetUserInfo() (*types.UserInfo, error) {
	// 从context中获取用户ID
	userID, err := utils.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return nil, errorx.ErrUserNotFound
	}

	return &types.UserInfo{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Mobile:   user.Mobile,
		Role:     user.Role,
	}, nil
}

// UpdateUser 更新用户信息
func (l *UserLogic) UpdateUser(req *types.UpdateUserReq) error {
	// 从context中获取用户ID
	userID, err := utils.GetUserId(l.ctx)
	if err != nil {
		return err
	}

	// 查找用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return errorx.ErrUserNotFound
	}

	// 更新字段
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Mobile != "" {
		// 检查手机号是否已被使用
		existUser, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
		if err == nil && existUser.Id != user.Id {
			return errorx.ErrUserExist
		}
		user.Mobile = req.Mobile
	}
	if req.Role != "" {
		user.Role = req.Role
	}

	return l.svcCtx.UserModel.Update(l.ctx, user)
}

// DeleteUser 删除用户
func (l *UserLogic) DeleteUser() error {
	// 从context中获取用户ID
	userID, err := utils.GetUserId(l.ctx)
	if err != nil {
		return err
	}

	return l.svcCtx.UserModel.Delete(l.ctx, userID)
}

// ListUsers 获取用户列表
func (l *UserLogic) ListUsers(req *types.UserListReq) (*types.UserListResp, error) {
	users, total, err := l.svcCtx.UserModel.List(l.ctx, req.Page, req.PageSize, req.Keyword)
	if err != nil {
		return nil, errorx.ErrServerError
	}

	list := make([]types.UserInfo, 0, len(users))
	for _, user := range users {
		list = append(list, types.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Mobile:   user.Mobile,
			Role:     user.Role,
		})
	}

	return &types.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
