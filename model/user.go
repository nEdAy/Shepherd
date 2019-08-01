package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Password string `gorm:"column:password" json:"-"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Credit   uint   `gorm:"column:credit" json:"credit"`
	Token    string `gorm:"-" json:"token"`
}

// TableName 返回user表名称
func (User) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "user")
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(user *User) (err error) {
	err = DB.Create(&user).Error
	return err
}

// GetUserById retrieves User by Id. Returns error if Id doesn't exist
func GetUserById(id uint) (user *User, err error) {
	user = new(User)
	if err = DB.First(&user, id).Error; err == nil {
		return user, nil
	}
	return nil, err
}

// GetUserByMobile retrieves User by mobile. Returns error if Id doesn't exist
func GetUserByMobile(mobile string) (user *User, err error) {
	user = new(User)
	if err = DB.Where("mobile = ?", mobile).Find(&user).Error; err == nil {
		return user, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if no records exist
func GetAllUser() (users []*User, err error) {
	if err = DB.Order("id desc").Select("id,mobile,nickname").Find(&users).Error; err == nil {
		return users, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if the record to be updated doesn't exist
func UpdateUserById(user *User) (err error) {
	// ascertain id exists in the database
	if err = DB.First(&user, user.ID).Error; err == nil {
		err = DB.Save(user).Error
	}
	return err
}

// DeleteUser deletes User by Id and returns error if the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	user := User{}
	// ascertain id exists in the database
	if err = DB.First(&user, id).Error; err == nil {
		err = DB.Where("id = ?", id).Delete(user).Error
	}
	return err
}

func IsUserExist(mobile string) (exist bool, err error) {
	var count int
	if err = DB.Model(&User{}).Where("mobile = ?", mobile).Count(&count).Error; err == nil {
		return count > 0, nil
	}
	return false, err
}
