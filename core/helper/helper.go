package helper

//Some encoding functions
import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
)

// Md5 encode
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {

	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	// 使用 HS256 对称加密算法生成 token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey)) // 使用对称密钥
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Email verification code
func MailCode(mail, code string) error {
	// todo: add your logic here and delete this line
	e := email.NewEmail()
	e.From = "Get <Drinkingbear996@gmail.com>"
	e.To = []string{mail}
	e.Subject = "Verification code"
	e.HTML = []byte("Your code is<h1>" + code + "</h1>")

	// Ensure you are using the correct app-specific password for Gmail
	err := e.SendWithTLS(
		"smtp.gmail.com:465",
		smtp.PlainAuth("", "Drinkingbear996@gmail.com", define.Gmailpwd, "smtp.gmail.com"),
		&tls.Config{ServerName: "smtp.gmail.com"},
	)

	if err != nil {
		return err
	}
	return nil
}

// Random Ver Code
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomVerCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

//Get UUid (Identity)

func GetUUID() string {
	return uuid.NewV4().String()
}
