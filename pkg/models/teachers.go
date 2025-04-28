package models

type Teachers struct {
	BaseModel
	TeacherNumber string `gorm:"type:varchar(255);uniqueIndex;not null" json:"teacher_number"`
	// 教师编号
	Introduction string `gorm:"type:text;not null" json:"introduction"`
	// 简介（教师个人介绍）/ 自己紹介・プロフィール
}
