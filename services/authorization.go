package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var FACK_SALT = "qwer1234"

// 明文：password
var FACK_PASSWORD = "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"

// 使用前端传入的注册信息注册。
func Register(info models.RegisterInfo) {}

// 使用前端传入的登入信息登入。
// 通过比对登入信息中的密码与数据库中的密码是否一致来判断是否登入成功。
// 成功时返回令牌。
func Login(info models.LoginInfo) (token string, err error) {
	user, err := GetUserAuthorizateInfo(info)

	if err != nil {
		return "", err
	}

	err = ComparePassword(info.Password, user.Password, user.Salt)

	if err != nil {
		return "", err
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.UID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	var tokenKey = utils.GetENV(configs.ENV_TOKEN_KEY)

	return rawToken.SignedString([]byte(tokenKey))
}

// 使用密码（由前端传入）与盐值生成哈希化的密码。
// 如果以字符串方式传入盐值，则必须是经过 base64 编码的盐值。
func HashPassword[Salt []byte | string](password string, salt Salt) (hashedPassword []byte, err error) {
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

// 从数据库中获取用户
func GetUserAuthorizateInfo(loginInfo models.LoginInfo) (*models.UserAuthorizateInfo, error) {
	// FIXME: 实装从数据库中获取用户列表
	base64Password, base64Salt := getFackPassword()

	if base64Password == "" || base64Salt == "" {
		return nil, errors.New("生成假数据失败")
	}

	users := []models.UserAuthorizateInfo{
		{
			UID:      "1001",
			Username: "root",
			Password: base64Password,
			Salt:     base64Salt,
		},
	}

	var target *models.UserAuthorizateInfo

	for _, user := range users {

		if user.Username == loginInfo.Username {
			target = &user
			break
		}
	}

	if target == nil {
		return nil, errors.New(utils.Messages.AUTHORIZE_FAILED_NO_USER)
	}

	return target, nil
}

// 生成假密码
func getFackPassword() (string, string) {
	base64Salt := base64.StdEncoding.EncodeToString([]byte(FACK_SALT))
	hashed, err := HashPassword(FACK_PASSWORD, base64Salt)

	if err != nil {
		log.Println(err)
		return "", ""
	}

	base64Password := base64.StdEncoding.EncodeToString(hashed)

	return base64Password, base64Salt
}
