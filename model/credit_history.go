package model

import "github.com/jinzhu/gorm"

// CreditHistory表
type CreditHistory struct {
	Model
	User    User   `gorm:"ForeignKey:UserId;AssociationForeignKey:Id"`
	UserId  int    `gorm:"column:user_id" json:"user_id"`
	Value   int    `gorm:"column:value" json:"value"`
	Message string `gorm:"column:message" json:"message"`
}

// TableName 返回credit_history表名称
func (CreditHistory) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "credit_history")
}

// GetAllCreditHistory retrieves all CreditHistory matches certain condition. Returns empty list if no records exist
func GetAllCreditHistory() (creditHistories []*CreditHistory, err error) {
	if err = DB.Order("create_time desc").Find(&creditHistories).Error; err == nil {
		return creditHistories, nil
	}
	return nil, err
}
