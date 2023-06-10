package service

import (
	"Chess/database"
	"Chess/model"
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"
)

func Register(username, password string) (err error) {
	err, user := database.FindUserByUsername(username)
	if err != nil {
		return err
	} else if user != (model.User{}) {
		return errors.New("用户名已注册")
	}
	user.Name = username
	//TODO:加密
	user.Password = password
	err = database.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(username, password string) (token, refreshToken string, err error) {
	var user model.User
	user, err = database.Login(username)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	//TODO:解密
	if user.Password != password {
		return "", "", errors.New("密码错误")
	}
	return createTokenAndRefreshToken(strconv.Itoa(user.Id))
}

func CheckRefreshTokenAndReturnToken(refreshToken string) (token, refreshTokenNew string, err error) {
	mapClaims, err := parseJWT(refreshToken, refreshTokenSecret)
	if err != nil {
		log.Println("jwt解密错误,1")
		return "", "", errors.New("Error parsing JWT:" + err.Error())
	}
	expStr := (*mapClaims)["exp"].(string)
	re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d+) ([+-]\d{4}) ([A-Z]+) m=(.*)`)
	match := re.FindStringSubmatch(expStr)
	if match == nil {
		log.Println("解析错误")
		return "", "", errors.New("解析时间戳错误")
	}
	exp, err := time.Parse("2006-01-02 15:04:05.9999999 -0700 MST", match[1]+" "+match[2]+" "+match[3])
	if err != nil {
		return "", "", err
	}
	if time.Now().After(exp) {
		return "", "", errors.New("刷新token过期")
	}
	id := (*mapClaims)["id"].(string)
	return createTokenAndRefreshToken(id)
}
