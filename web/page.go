package web

import (
	"github.com/dchest/captcha"
)

// 页面
type Page struct {
	// 标题
	Title string
	// 验证码字符串
	Cid string
	// 是否是管理员
	IsManager bool
}

// 新建页面对象
func PageNew(title string, isManager bool) (p Page) {
	p.Title = title
	p.Cid = captcha.NewLen(4)
	p.IsManager = isManager
	return
}
