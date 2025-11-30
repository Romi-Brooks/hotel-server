package model

import "time"

// SysUser 用户/管理员模型（与数据库表sys_user映射）
type SysUser struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"-"` // 不返回密码
	Role       string    `json:"role"`
	CreateTime time.Time `json:"createTime"` // 小驼峰标签
	UpdateTime time.Time `json:"updateTime"`
}

// 角色常量
const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
)
