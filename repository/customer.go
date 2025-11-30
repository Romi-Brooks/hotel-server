package repository

import (
	"context"
	"fmt"
	"hotel-server/config"
	"hotel-server/model"
)

// CustomerRepository 客户数据访问层
type CustomerRepository struct{}

// GetCustomerList 获取所有客户列表
func (r *CustomerRepository) GetCustomerList(ctx context.Context) ([]model.Customer, error) {
	var customers []model.Customer
	sql := `SELECT customer_id, name, phone, id_card_or_passport, email, create_time, update_time FROM customer;`

	rows, err := config.DB.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("查询客户列表失败: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Customer
		err := rows.Scan(
			&c.CustomerId, &c.Name, &c.Phone,
			&c.IdCardOrPassport, &c.Email, &c.CreateTime, &c.UpdateTime,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描客户数据失败: %w", err)
		}
		customers = append(customers, c)
	}
	return customers, nil
}
