package model

import "time"

// Customer 客户表模型
type Customer struct {
	CustomerId       string    `json:"customerId"`       // 客户ID（主键）
	Name             string    `json:"name"`             // 姓名
	Phone            string    `json:"phone"`            // 电话
	IdCardOrPassport string    `json:"idCardOrPassport"` // 身份证/护照
	Email            string    `json:"email"`            // 邮箱
	CreateTime       time.Time `json:"createTime"`       // 创建时间
	UpdateTime       time.Time `json:"updateTime"`       // 更新时间
}
