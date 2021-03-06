/*
 * @Description:
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-04 14:46:00
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-01-04 18:40:45
 */

package eagletunnel

import (
	"../eaglelib/src"
)

// Handler 请求处理者
type Handler interface {
	Handle(request Request, tunnel *eaglelib.Tunnel) error
}
