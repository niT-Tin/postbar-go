package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"postbar/datamodels"
	err2 "postbar/err"
	"postbar/repositories"
)

type IUserService interface {
	IsPwsSuccess(account, pwd string) (user *datamodels.User, isOk bool)  //检查帐号密码是否正确
	Register(*datamodels.User) (int64, error)                             //注册
	Login(userName, pwd string) bool                                      //登陆
	Post(post *datamodels.SinglePost) (postId int64, err error)           //发帖
	PostComment(comment *datamodels.Comment) (commentId int64, err error) //评论
	Likes(commentId int64) (bool, error)                                  //评论点赞
	GetOneById(userId int64) *datamodels.User                             //根据id获得用户信息
	GetOneByAccount(account string) (*datamodels.User, error)             //根据帐号获得用户信息
	Update(*datamodels.User) (bool, error)                                //修改用户信息
	CreateBar(bar *datamodels.PostBar) (barId int64, err error)           //创建贴吧
	Delete(*datamodels.User) int64                                        //删除用户
	DeleteById(int64) int64                                               //根据id删除用户
	DeleteByAccount(account string) int64                                 //根据帐号删除用户
}

// 记录service层错误
func reciteErrorINService(err1 error) bool {
	if err1 != nil {
		s := err1.Error() + "on service"
		e := errors.New(s)
		err2.ReciteErr(&e)
		return true
	}
	return false
}

// GenerateHashedPassword 密码加密
func GenerateHashedPassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// NewUserService 创建user服务
func NewUserService(repo repositories.IUserRepository) IUserService {
	return &UserService{UserRepo: repo}
}

type UserService struct {
	UserRepo repositories.IUserRepository
}

func ValidateUser(userPassword, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

// IsPwsSuccess 判断用户帐号密码是否正确
func (u *UserService) IsPwsSuccess(account, pwd string) (user *datamodels.User, isOk bool) {
	var err error
	user1, err := u.UserRepo.GetUserByAccount(account)
	if reciteErrorINService(err) {
		return
	}
	ok, _ := ValidateUser(pwd, user1.Password)
	if !ok {
		return &datamodels.User{}, false
	}
	return user1, true
}

// Register 用户注册服务
func (u *UserService) Register(usr *datamodels.User) (int64, error) {
	err := u.UserRepo.Insert(usr)
	if reciteErrorINService(err) {
		return 0, nil
	}
	return usr.Userid, nil
}

// Login 用户登陆服务
func (u *UserService) Login(account, pwd string) bool {
	if _, ok := u.IsPwsSuccess(account, pwd); !ok {
		return false
	} else {
		return true
	}
}

// Post 用户发表帖子服务
func (u *UserService) Post(post *datamodels.SinglePost) (postId int64, err error) {
	err = u.UserRepo.PostCrud().Create(post)
	if err != nil {
		reciteErrorINService(err)
		return
	}
	return post.PosterId, nil
}

// PostComment 用户评论服务
func (u *UserService) PostComment(comment *datamodels.Comment) (commentId int64, err error) {
	_, err = u.UserRepo.CommentCrud().Insert(comment)
	if err != nil {
		reciteErrorINService(err)
		return
	}
	return comment.CommentId, nil
}

// Likes 插入点赞数服务
func (u *UserService) Likes(commentId int64) (bool, error) {
	one := u.UserRepo.CommentCrud().IncreaseLikeNumByOne(commentId)
	if one {
		return one, nil
	} else {
		var e = errors.New("点赞插入失败")
		reciteErrorINService(e)
		return false, e
	}
}

// GetOneById 根据用户id获取用户信息服务
func (u *UserService) GetOneById(userId int64) *datamodels.User {
	return u.UserRepo.GetOneById(userId)
}

// GetOneByAccount 根据用户帐号获取用户服务
func (u *UserService) GetOneByAccount(account string) (*datamodels.User, error) {
	byAccount, err := u.UserRepo.GetUserByAccount(account)
	if err != nil {
		reciteErrorINService(err)
		return &datamodels.User{}, err
	}
	return byAccount, nil
}

// Update 更新用户信息服务
func (u *UserService) Update(usr *datamodels.User) (bool, error) {
	err := u.UserRepo.Update(usr)
	if err != nil {
		reciteErrorINService(err)
		return false, err
	}
	return true, nil
}

// CreateBar 创建贴吧服务（只有用户达到那个一定等级才能创建）等级判断在父级实现
func (u *UserService) CreateBar(bar *datamodels.PostBar) (barId int64, err error) {
	err = u.UserRepo.PostBarCrud().Create(bar)
	if err != nil {
		reciteErrorINService(err)
		return
	}
	return bar.PostBarId, nil
}

// Delete 删除用户服务
func (u *UserService) Delete(usr *datamodels.User) int64 {
	if usr == nil {
		return 0
	}
	return u.DeleteById(usr.Userid)
}

// DeleteById 根据id删除用户
func (u *UserService) DeleteById(usrId int64) int64 {
	if usrId == 0 {
		return 0
	}
	err := u.UserRepo.Delete(usrId)
	if err != nil {
		reciteErrorINService(err)
		return 0
	}
	return usrId
}

// DeleteByAccount 根据帐号删除用户
func (u *UserService) DeleteByAccount(account string) int64 {
	if account == "" {
		return 0
	}
	byAccount, err := u.UserRepo.GetUserByAccount(account)
	if err != nil {
		reciteErrorINService(err)
		return 0
	}
	return u.Delete(byAccount)
}
