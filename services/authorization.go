package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"golang.org/x/crypto/bcrypt"
)

// 使用前端传入的登入信息登入。
// 通过比对登入信息中的密码与数据库中的密码是否一致来判断是否登入成功。
func Login(info models.LoginInfo) error {
	user, err := GetUserInfo(info)

	if err != nil {
		return err
	}

	err = ComparePassword(info.Password, user.Password, user.Salt)

	return err
}

// 使用密码（由前端传入）与盐值生成哈希化的密码。
// 如果以字符串方式传入盐值，则必须是经过 base64 编码的盐值。
func HashPassword[Salt []byte | string](password string, salt Salt) ([]byte, error) {
	var byteSalt []byte

	switch any(salt).(type) {
	case string:
		decoded, err := base64.StdEncoding.DecodeString(string(salt))

		if err != nil {
			return nil, err
		}

		byteSalt = decoded

	case []byte:
		byteSalt = []byte(salt)
	}

	appended := append([]byte(password), byteSalt...)

	return bcrypt.GenerateFromPassword(appended, configs.BCRYPT_COST)
}

// 比对密码。
// 将传入的密码同数据库中的盐值进行哈希化，与数据库中的哈希化后的密码及盐值进行比对。
func ComparePassword(password string, base64Hashed string, base64Salt string) error {
	// 解码密码
	hashed, err := base64.StdEncoding.DecodeString(base64Hashed)

	if err != nil {
		return err
	}

	// 解码盐值
	salt, err := base64.StdEncoding.DecodeString(base64Salt)

	if err != nil {
		return err
	}

	// 拼接传入密码与盐值
	payload := append([]byte(password), salt...)

	// 比对拼接结果与哈希后密码
	err = bcrypt.CompareHashAndPassword(hashed, payload)

	if err != nil {
		log.Println(err)
		err = errors.New("密码不一致")
	}

	return err
}

// 生成盐值
func GenerateBase64Salt() (string, error) {
	salt := make([]byte, configs.SALT_LENGTH)

	// 填充随机字节
	_, err := rand.Read(salt)

	if err != nil {
		return "", nil
	}

	// 将盐值编码为字符串以便于储存和传输
	result := base64.StdEncoding.EncodeToString(salt)

	return result, nil
}
