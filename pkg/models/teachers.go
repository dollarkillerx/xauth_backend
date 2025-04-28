package models

type Teachers struct {
	BaseModel
	Introduction string `gorm:"type:text;not null" json:"introduction"`
	// 简介（教师个人介绍）/ 自己紹介・プロフィール
}
