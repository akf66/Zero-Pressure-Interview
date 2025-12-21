package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// EmailSender 邮件发送器
type EmailSender struct {
	smtpHost string
	smtpPort int
	username string
	password string
	fromName string
}

// NewEmailSender 创建邮件发送器
func NewEmailSender(smtpHost string, smtpPort int, username, password, fromName string) *EmailSender {
	return &EmailSender{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
		fromName: fromName,
	}
}

// SendVerifyCode 发送验证码邮件
func (s *EmailSender) SendVerifyCode(to, code, purpose string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(s.username, s.fromName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", s.getSubject(purpose))
	m.SetBody("text/html", s.getHTMLBody(code, purpose))

	d := gomail.NewDialer(s.smtpHost, s.smtpPort, s.username, s.password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// getSubject 根据用途获取邮件主题
func (s *EmailSender) getSubject(purpose string) string {
	switch purpose {
	case "1": // REGISTER
		return "【零压面试】注册验证码"
	case "2": // LOGIN
		return "【零压面试】登录验证码"
	case "3": // RESET_PASSWORD
		return "【零压面试】重置密码验证码"
	case "4": // CHANGE_PHONE
		return "【零压面试】修改手机号验证码"
	case "5": // CHANGE_EMAIL
		return "【零压面试】修改邮箱验证码"
	default:
		return "【零压面试】验证码"
	}
}

// getHTMLBody 根据用途获取邮件HTML内容
func (s *EmailSender) getHTMLBody(code, purpose string) string {
	purposeText := s.getPurposeText(purpose)

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f9f9f9;
        }
        .header {
            background-color: #4CAF50;
            color: white;
            padding: 20px;
            text-align: center;
            border-radius: 5px 5px 0 0;
        }
        .content {
            background-color: white;
            padding: 30px;
            border-radius: 0 0 5px 5px;
        }
        .code {
            font-size: 32px;
            font-weight: bold;
            color: #4CAF50;
            text-align: center;
            padding: 20px;
            background-color: #f0f0f0;
            border-radius: 5px;
            margin: 20px 0;letter-spacing: 5px;
        }
        .footer {
            text-align: center;
            margin-top: 20px;
            color: #666;
            font-size: 12px;
        }
        .warning {
            color: #ff6b6b;
            font-size: 14px;
            margin-top: 15px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>零压面试</h1>
        </div>
        <div class="content">
            <h2>%s</h2>
            <p>您好！</p>
            <p>您正在进行%s操作，您的验证码是：</p>
            <div class="code">%s</div>
            <p>验证码有效期为 <strong>5分钟</strong>，请尽快使用。</p>
            <p class="warning">⚠️ 如果这不是您本人的操作，请忽略此邮件。</p>
        </div>
        <div class="footer">
            <p>此邮件由系统自动发送，请勿回复。</p><p>© 2024 零压面试 - Zero Pressure Interview</p>
        </div>
    </div>
</body>
</html>
`, purposeText, purposeText, code)
}

// getPurposeText 根据用途获取文本描述
func (s *EmailSender) getPurposeText(purpose string) string {
	switch purpose {
	case "1":
		return "账号注册"
	case "2":
		return "账号登录"
	case "3":
		return "重置密码"
	case "4":
		return "修改手机号"
	case "5":
		return "修改邮箱"
	default:
		return "身份验证"
	}
}
