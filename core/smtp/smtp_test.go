package smtp

import (
	"log"
	"testing"
)

func TestUuid(t *testing.T) {
	s := NewSmtp("", "smtp.qq.com:587", "jobob@qq.com", "测试密码")
	from := "jobob@qq.com"
	to := []string{"hubinmdj@163.com", "243194995@qq.com"}
	subject := "Hi AiXinGe"
	content := "This is the email body"
	err := s.SendMail(from, to, subject, content)
	if err != nil {
		log.Fatal(err)
	}
}
