package models

type Setting struct {
	BaseModel
	SchoolName string `gorm:"type:varchar(255);not null" json:"school_name"`
	// 学校名称 / 学校名

	PrincipalName string `gorm:"type:varchar(255);not null" json:"principal_name"`
	// 校长姓名 / 校長名

	EmailDomain string `gorm:"type:varchar(255);not null" json:"email_domain"`
	// 邮箱后缀（例：@school.jp）/ メールドメイン（例：@school.jp）

	AllowRegister bool `gorm:"default:true;not null" json:"allow_register"`
	// 是否允许注册 / 登録許可（true＝許可、false＝禁止）

	AllowChangeDevice bool `gorm:"default:true;not null" json:"allow_change_device"`
	// 是否允许更换登录设备 / デバイス変更許可（true＝許可、false＝禁止）

	RestrictClockInIP bool `gorm:"default:false;not null" json:"restrict_clockin_ip"`
	// 是否限制打卡IP / 打刻IP制限（true＝制限、false＝自由）
}
