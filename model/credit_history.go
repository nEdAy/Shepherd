package model

import "github.com/jinzhu/gorm"

type CreditHistory struct {
	gorm.Model
	UserId  uint
	Change  int
	Credit  int
	Message string
}

func GetCreditHistoriesByUserId(userId uint) (creditHistories []*CreditHistory, err error) {
	if err = db.Where("user_id = ?", userId).Order("create_time desc").Find(&creditHistories).Error; err == nil {
		return creditHistories, nil
	}
	return nil, err
}

func ModifyCreditHistory(userId uint, change int, message string) (*CreditHistory, error) {
	tx := db.Begin()
	// 根据userId（主键）查找
	user := User{}
	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 更新User.credit
	credit := user.Credit + change
	if err := tx.Model(&user).Select("credit").Updates(map[string]interface{}{"credit": credit}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 创建积分历史记录
	creditHistory := CreditHistory{UserId: userId, Change: change, Credit: credit, Message: message}
	if err := tx.Create(creditHistory).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &creditHistory, nil
}
