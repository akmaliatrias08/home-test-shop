package users

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"online-shop-home-test/modules/role"
	user "online-shop-home-test/modules/users/dto"
	users "online-shop-home-test/modules/users/dto"
	"online-shop-home-test/modules/users/models"
	"online-shop-home-test/modules/utils"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if err := utils.DB.Model(&user).Where("username", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func Register(createUserDTO user.CreateUserDTO) (models.User, error) {
	var user models.User

	//validate username is unique
	userByUsername, err := GetUserByUsername(createUserDTO.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}

	if strings.Compare(createUserDTO.Username, userByUsername.Username) == 0 {
		return user, errors.New("username already exist")
	}

	//validate if role is exist
	role, err := role.GetRoleById(createUserDTO.RoleID)
	if err != nil {
		return user, err
	}

	user.Name = createUserDTO.Name
	user.Username = createUserDTO.Username
	user.RoleID = role.ID
	user.Role = &role
	user.Salt = uuid.New().String()

	//hash password
	passwordHashed := pbkdf2.Key([]byte(createUserDTO.Password), []byte(user.Salt), 4096, 100, sha512.New)
	user.Password = hex.EncodeToString(passwordHashed)

	if err := utils.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func Login(loginDTO users.LoginDTO) (users.TokenResponse, error) {
	var data users.TokenResponse

	//validate is username exist
	user, err := GetUserByUsername(loginDTO.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return data, errors.New("username not found")
	}

	if err != nil {
		return data, err
	}

	//validate password
	calculatedPasswordHash := pbkdf2.Key([]byte(loginDTO.Password), []byte(user.Salt), 4096, 100, sha512.New)
	if hex.EncodeToString(calculatedPasswordHash) != user.Password {
		return data, errors.New("password is wrong")
	}

	accessToken, err := GenerateToken(user.ID.String())
	if err != nil {
		return data, err
	}

	data.AccessToken = accessToken

	return data, nil
}

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId

	//set expire at
	durationHour := 24

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(durationHour)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
