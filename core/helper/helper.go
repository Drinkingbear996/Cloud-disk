package helper

//一些加密函数
import (
	"cloud-disk/core/define"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// Md5 encode
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	// id, identity, name 用于生成 JWT Token
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	// 使用 HS256 对称加密算法生成 token
	//之前使用ES256 加密报错了，
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey)) // 使用对称密钥
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
