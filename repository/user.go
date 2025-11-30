package repository

import (
	"context"
	"fmt"
	"hotel-server/config"
	"hotel-server/model"
)

// UserRepository 用户数据访问层
type UserRepository struct{}

// GetUserByUsername 根据用户名查询用户
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*model.SysUser, error) {
	var user model.SysUser
	query := `SELECT id, username, password, role, create_time, update_time FROM sys_user WHERE username = $1`
	err := config.DB.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreateTime,
		&user.UpdateTime,
	)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	return &user, nil
}

// CreateAdmin 添加普通管理员
func (r *UserRepository) CreateAdmin(ctx context.Context, username, password string) error {
	query := `INSERT INTO sys_user (username, password, role) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(ctx, query, username, password, model.RoleAdmin)
	if err != nil {
		return fmt.Errorf("添加管理员失败: %v", err)
	}
	return nil
}

// ListAdmin 查询所有普通管理员
func (r *UserRepository) ListAdmin(ctx context.Context) ([]model.SysUser, error) {
	var users []model.SysUser
	query := `SELECT id, username, role, create_time, update_time FROM sys_user WHERE role = $1`
	rows, err := config.DB.Query(ctx, query, model.RoleAdmin)
	if err != nil {
		return nil, fmt.Errorf("查询管理员列表失败: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.SysUser
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
			&user.CreateTime,
			&user.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
