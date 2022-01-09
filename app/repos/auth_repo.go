package repos

import (
	"encoding/json"
	"fmt"
	"swift_typing_api/app/dbs"
	"swift_typing_api/app/models"
)

//const prefix = "_U_"

type IAuthRepo interface {
	Set(key string, authInfo *models.AuthInfo) error
	Get(key string, authInfo *models.AuthInfo) error
	GetKeys(keyStart string, keyEnd string) ([]string, error)
	Remove(key ...string) error
	SetCaptcha(key string, captcha string) error
	GetCaptcha(key string) (string, error)
}

type AuthRepo struct {
	gredis dbs.IRedis
}

func NewAuthRepo(gredis dbs.IRedis) IAuthRepo {
	return &AuthRepo{gredis: gredis}
}

func (authRepo *AuthRepo) Set(key string, authInfo *models.AuthInfo) error {
	//key = strconv.Itoa(authInfo.UserId) + prefix + key
	authInfoBytes, _ := json.Marshal(authInfo)
	err := authRepo.gredis.Set(key, authInfoBytes)
	return err
}

func (authRepo *AuthRepo) Get(key string, authInfo *models.AuthInfo) error {
	err := authRepo.gredis.Get(key, &authInfo)
	authRepo.gredis.Expire(key)

	return err
}

func (authRepo *AuthRepo) Remove(key ...string) error {
	err := authRepo.gredis.Remove(key...)
	return err
}

func (authRepo *AuthRepo) GetKeys(keyStart string, keyEnd string) ([]string, error) {
	key := fmt.Sprintf("*%s%s*", keyStart, keyEnd)
	keys, err := authRepo.gredis.Keys(key)
	return keys, err
}

func (authRepo *AuthRepo) SetCaptcha(key string, captcha string) error {
	//key = strconv.Itoa(authInfo.UserId) + prefix + key
	captchaBytes := []byte(captcha)
	err := authRepo.gredis.SetCaptcha(key, captchaBytes)
	return err
}

func (authRepo *AuthRepo) GetCaptcha(key string) (string, error) {
	captcha, err := authRepo.gredis.GetCaptcha(key)
	return captcha, err
}
