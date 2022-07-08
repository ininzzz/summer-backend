package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

// MailConfig 邮箱发送配置
type MailConfig struct {
	Account  string //xxx@qq.com
	Password string //需开通qq邮箱的smtp服务（免费）后获取
	Port     int    // QQ：POP/SMTP 587
	Host     string // QQ：smtp.qq.com
}

var mail_conf *MailConfig

// SendMail 发送邮件
// from 发送者别名，mailTo 发送对象，subject主题，body 内容
func SendMail(from string, mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	//这种方式可以添加别名，即“XX官方”
	m.SetHeader("From", m.FormatAddress(mail_conf.Account, from))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文
	d := gomail.NewDialer(mail_conf.Host, mail_conf.Port, mail_conf.Account, mail_conf.Password)
	err := d.DialAndSend(m)
	return err
}

//初始化邮箱设定
func Init_Email_Conf() {
	mail_conf = &MailConfig{os.Getenv("EMAIL_SMTP"), os.Getenv("EMAIL_PWD"), 587, "smtp.qq.com"}
}
