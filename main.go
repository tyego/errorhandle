package main

import (
	"errors"
	"fmt"
	"github.com/tyego/errorhandle/dao"
)

/*
作业：
	我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
	为什么，应该怎么做请写出代码？

分析：
	sql.ErrNoRows 本质上就是数据库这种没有这条记录。
	至于要不要 Wrap 这个 error，抛给上层，取决于具体业务场景，比如：
	1，假设我要查询某个用户的订单列表，那查不到并不算是错误，而是用户本身就没下过单，这时候就应该自己处理，不需要上抛；
	2，假设我要查询某个用户的身份信息，比如手机号，那查不到可能是因为你参数传错了，那这时候我就要网上抛，并且携带一些额外信息。
	总之，这个答案并不是唯一的。

以下实现两种场景的代码：

*/

func main() {
	testNoWarp()
	testWarp()
}

func testNoWarp() {
	result, err := dao.QueryOrders(2)
	if err != nil {
		fmt.Printf("QueryOrders fail, err = %v\n", err)
		return
	}
	fmt.Printf("QueryOrders success, result = %v\n", result)
}

func testWarp() {
	result, err := dao.QueryUserInformation(2)
	if err != nil {
		fmt.Printf("QueryUserInformation fail, err = %v\n", err)
		fmt.Printf("Unwrap err = %v\n", errors.Unwrap(err))
		return
	}
	fmt.Printf("QueryUserInformation success, result = %v\n", result)
}
