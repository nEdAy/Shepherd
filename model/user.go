package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Mobile   string `json:"mobile"`
	Password string `json:"-"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Credit   int    `json:"credit"`
	Token    string `json:"token" gorm:"-"`
}

func AddUser(user *User) (err error) {
	err = db.Create(&user).Error
	return err
}

func GetUserById(id uint) (user *User, err error) {
	user = new(User)
	if err = db.First(&user, id).Error; err == nil {
		return user, nil
	}
	return nil, err
}

func GetUserByMobile(mobile string) (user *User, err error) {
	user = new(User)
	if err = db.Where("mobile = ?", mobile).Find(&user).Error; err == nil {
		return user, nil
	}
	return nil, err
}

func GetAllUser() (users []*User, err error) {
	if err = db.Order("id desc").Select("id,mobile,nickname").Find(&users).Error; err == nil {
		return users, nil
	}
	return nil, err
}

func UpdateUserById(user *User) (err error) {
	// ascertain id exists in the database
	if err = db.First(&user, user.ID).Error; err == nil {
		err = db.Save(user).Error
	}
	return err
}

func DeleteUser(id int) (err error) {
	user := User{}
	// ascertain id exists in the database
	if err = db.First(&user, id).Error; err == nil {
		err = db.Where("id = ?", id).Delete(user).Error
	}
	return err
}

func IsUserExist(mobile string) (exist bool, err error) {
	var count int
	if err = db.Model(&User{}).Where("mobile = ?", mobile).Count(&count).Error; err == nil {
		return count > 0, nil
	}
	return false, err
}
