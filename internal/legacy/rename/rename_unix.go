
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有（c）2013-2014 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

//+建设！窗户，！计划9

package rename

import (
	"os"
)

//原子提供原子文件重命名。如果新路径
//已经存在。
func Atomic(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}
