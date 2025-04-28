package models

import "time"

type Student struct {
	BaseModel
	StudentNumber string `gorm:"type:varchar(50);uniqueIndex;not null" json:"student_number"`
	// 学籍番号 / 学籍番号（がくせきばんごう）

	Department string `gorm:"type:varchar(100);not null" json:"department"`
	// 学系 / 学部（がくぶ）

	Major string `gorm:"type:varchar(100);not null" json:"major"`
	// 学科 / 学科（がっか）

	EnrollmentDate time.Time `gorm:"not null" json:"enrollment_date"`
	// 利用开始日期（入学日期）/ 利用開始日（入学日）

	GraduationDate *time.Time `json:"graduation_date"`
	// 利用结束日期（毕业日期，可为空）/ 利用終了日（卒業日、NULL可）

	Birthday time.Time `gorm:"not null" json:"birthday"`
	// 出生日期 / 生年月日（せいねんがっぴ）

	Gender string `gorm:"type:enum('male','female','other');not null" json:"gender"`
	// 性别（male=男，female=女，other=其他）/ 性別（male=男性、female=女性、other=その他）

	BirthPlace string `gorm:"type:varchar(255);not null" json:"birth_place"`
	// 出生地 / 出生地（しゅっしょうち）

	Address string `gorm:"type:varchar(255);not null" json:"address"`
	// 现住地址 / 現住所（げんじゅうしょ）

	EmergencyContactName string `gorm:"type:varchar(255);not null" json:"emergency_contact_name"`
	// 紧急联系人姓名 / 緊急連絡先氏名（きんきゅうれんらくさき しめい）

	EmergencyContactPhone string `gorm:"type:varchar(20);not null" json:"emergency_contact_phone"`
	// 紧急联系人电话 / 緊急連絡先電話番号（きんきゅうれんらくさき でんわばんごう）

	EmergencyContactRelation string `gorm:"type:varchar(100);not null" json:"emergency_contact_relation"`
	// 紧急联系人关系（如：父亲、母亲）/ 緊急連絡先との関係（例：父、母）

	EmergencyContactAddress string `gorm:"type:varchar(255);not null" json:"emergency_contact_address"`
	// 紧急联系人地址 / 緊急連絡先住所（きんきゅうれんらくさき じゅうしょ）
}
