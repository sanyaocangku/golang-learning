// 使用命令: go mod init testmod

package testmod

import "fmt"

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
