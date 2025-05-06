package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"go.uber.org/zap"

	"Go/global"

)

const (
	SMTPHost = "sandbox.smtp.mailtrap.io"
	SMTPPort = "2525"
	SMTPUsername = "9c991bd0ac88cc"
	SMTPPassword = "8120a5cf073a43"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\"; \r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "Test"},
		To:      to,
		Subject: "Xác nhận đăng ký tài khoản",
		Body:    fmt.Sprintf("Mã xác nhận của bạn là: %s", otp),
	}

	messageEmail := BuildMessage(contentEmail)

	// send smtp
	fmt.Printf("Adds %s\n", SMTPHost+":"+SMTPPort)

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	
	// Send email
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Error sending email:", zap.Error(err))
		return err
	}

	return nil

}

func SendTemplateEmailOtp(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	htmlBody, err  := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
        return err
    }
	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTempalte map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("template-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTempalte)
	if err != nil {
        return "", err
    }
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "Test"},
		To:      to,
		Subject: "Xác nhận đăng ký tài khoản",
		Body:    htmlTemplate,
	}

	messageEmail := BuildMessage(contentEmail)

	// send smtp
	fmt.Printf("Adds %s\n", SMTPHost+":"+SMTPPort)

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	
	// Send email
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Error sending email:", zap.Error(err))
		return err
	}

	return nil
}