package controller

import (
	"hotel-server/repository"
	"hotel-server/util"
	"log" // 导入日志包
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// JWT密钥（生产环境需配置到环境变量）
var jwtSecret = []byte("hotel_jwt_secret_2025")

// UserController 用户控制器
type UserController struct {
	repo *repository.UserRepository
}

// NewUserController 初始化用户控制器
func NewUserController() *UserController {
	return &UserController{
		repo: &repository.UserRepository{},
	}
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	// 1. 接收前端参数
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Fail(ctx, 400, "参数错误："+err.Error())
		return
	}

	// 打印前端传入的参数（排查参数是否正确接收）
	log.Printf("前端传入的登录参数：username=%s, password=%s", req.Username, req.Password)

	// 2. 查询用户
	user, err := c.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		// 打印真实的数据库查询错误（关键！）
		log.Printf("查询用户失败：%v，用户名：%s", err, req.Username)
		util.Fail(ctx, 500, "用户不存在或查询失败："+err.Error()) // 返回真实错误提示
		return
	}

	// 打印查询到的用户（排查是否查询到正确数据）
	log.Printf("查询到的用户：%+v", user)

	// 3. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		// 打印密码验证错误（关键！）
		log.Printf("密码验证失败：%v，用户名：%s", err, req.Username)
		util.Fail(ctx, 400, "密码错误："+err.Error())
		return
	}

	// 4. 生成JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 24小时过期
	})
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("生成JWT令牌失败：%v", err)
		util.Fail(ctx, 500, "生成令牌失败："+err.Error())
		return
	}

	// 5. 返回响应
	util.Success(ctx, gin.H{
		"token": tokenStr,
		"userInfo": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	}, "登录成功")
}

// AddAdmin 添加普通管理员
func (c *UserController) AddAdmin(ctx *gin.Context) {
	// 1. 接收参数
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.Fail(ctx, 400, "参数错误："+err.Error())
		return
	}

	// 2. 密码加密
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		util.Fail(ctx, 500, "密码加密失败")
		return
	}

	// 3. 添加管理员
	err = c.repo.CreateAdmin(ctx, req.Username, string(hashPwd))
	if err != nil {
		util.Fail(ctx, 500, err.Error())
		return
	}

	util.Success(ctx, nil, "添加管理员成功")
}

// ListAdmin 查询管理员列表
func (c *UserController) ListAdmin(ctx *gin.Context) {
	users, err := c.repo.ListAdmin(ctx)
	if err != nil {
		util.Fail(ctx, 500, err.Error())
		return
	}
	// 统一返回 {list: 数组} 的结构，与房间/预订接口对齐
	util.Success(ctx, gin.H{"list": users}, "查询成功")
}
