package dao

import (
	"database/sql"
	"errors"
	"fmt"
)

func QueryOrders(id int) ([]string, error) {
	result, err := simulationQueryMySql(id)
	if err != nil {
		if err == sql.ErrNoRows {
			// 该场景错误信息不用处理，返回空集合即可
			return []string{}, nil
		}
		// 其他错误正常往上抛，此处根据具体业务场景可以选择包装或不包装
		return nil, err
	}
	// 没问题，正常返回
	return result, nil
}

func QueryUserInformation(id int) ([]string, error) {
	result, err := simulationQueryMySql(id)
	if err != nil {
		if err == sql.ErrNoRows {
			// 这里需要将错误网上抛，并且可以 warp 一些额外信息 (包装也可使用 github.com/pkg/errors.Wrap 方法)
			// 使用 "%w" 包装的错误，使用系统的 errors.Unwrap(err) 即可解出
			return nil, fmt.Errorf("this is a invalid user, detail err is [%w]", err)
		}
		// 其他错误正常往上抛，此处根据具体业务场景可以选择包装或不包装
		return nil, err
	}
	// 没问题，正常返回
	return result, nil
}

// 模拟一个外部的MySQL的查询
// 合法id参数必须大于0，且数据库中只有一条id为1的数据
func simulationQueryMySql(id int) ([]string, error) {
	if id <= 0 {
		return nil, errors.New("invalid param")
	}
	if id == 1 {
		return []string{"item1", "item2"}, nil
	}
	return nil, sql.ErrNoRows
}
