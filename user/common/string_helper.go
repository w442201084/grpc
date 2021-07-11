package common

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)


// 加密用户密码
func GeneratorPassword(userPassword string) ([]byte , error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword) , bcrypt.DefaultCost)
}

// 校验用户密码
func ValidateUserPassword(userPassword string , hashed string ) ( isOk bool , err error ){
	if err = bcrypt.CompareHashAndPassword([]byte(hashed) ,
		[]byte(userPassword)); nil != err {
		return false , errors.New("密码校验错误")
	}
	return true , nil
}