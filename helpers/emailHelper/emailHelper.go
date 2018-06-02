package emailHelper

import (
	"bytes"
	"fmt"
	"echo-boilerplate/conf"
	"echo-boilerplate/models"
	"html/template"
	"net/smtp"
)

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles("templates/" + templateFileName + ".html")
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendEmail(templateFileName string, sendTo string, subject string, data interface{}) (bool, error) {
	smtpInfo := conf.Conf.SMTP

	addr := smtpInfo.Server + ":" + smtpInfo.Port
	auth := smtp.PlainAuth("", smtpInfo.From, smtpInfo.Pwd, smtpInfo.Server)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := fmt.Sprintf("From: %s\n", smtpInfo.From)
	to := fmt.Sprintf("To: %s\n", sendTo)
	subject = fmt.Sprintf("Subject: %s\n", subject)

	body, _ := parseTemplate(templateFileName, data)
	msg := []byte(from + to + subject + mime + "\n" + body)

	if err := smtp.SendMail(addr, auth, smtpInfo.From, []string{sendTo}, msg); err != nil {
		return false, err
	}
	return true, nil
}

// AuthCodeEmailSend : 인증용 메일 발송
func AuthCodeEmailSend(user *models.User, email string) (bool, error) {

	server := conf.Conf.Server
	url := fmt.Sprintf("%s://%s/api/auth/email_auth/%s",
		server.Protocol,
		server.Domain,
		user.UserCode)

	data := struct {
		Code string
		URL  string
	}{
		Code: user.UserCode,
		URL:  url,
	}

	return SendEmail("emailAuthTemplate", email, "GGFighter 이메일 인증", data)
}
