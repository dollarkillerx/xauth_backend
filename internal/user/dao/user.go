package dao

import (
	"github.com/dollarkillerx/xauth_backend/api/user"
	"time"
)

type RegisterStudentRequest struct {
	Name                     string    `json:"name" validate:"required"`                           // 姓名 / 氏名（しめい）
	NameKana                 string    `json:"name_kana" validate:"required"`                      // 姓名(カナ) / 氏名（カナ）
	Avatar                   string    `json:"avatar" validate:"omitempty,url"`                    // 头像 / アバター（URL）
	Email                    string    `json:"email" validate:"required,email"`                    // 邮箱 / メールアドレス
	Phone                    string    `json:"phone" validate:"required,e164"`                     // 手机号 / 電話番号（国際形式）
	Password                 string    `json:"password" validate:"required,min=6"`                 // 密码 / パスワード（6文字以上）
	DeviceID                 string    `json:"device_id" validate:"omitempty"`                     // 设备ID / デバイスID
	StudentNumber            string    `json:"student_number" validate:"required"`                 // 学号 / 学籍番号（がくせきばんごう）
	Department               string    `json:"department" validate:"omitempty"`                    // 学系 / 学部（がくぶ）
	Major                    string    `json:"major" validate:"omitempty"`                         // 学科 / 学科（がっか）
	EnrollmentDate           time.Time `json:"enrollment_date" validate:"omitempty"`               // 入学日期 / 入学日（にゅうがくび）
	GraduationDate           time.Time `json:"graduation_date" validate:"omitempty"`               // 毕业日期 / 卒業日（そつぎょうび）
	Birthday                 time.Time `json:"birthday" validate:"omitempty"`                      // 生日 / 生年月日（せいねんがっぴ）
	Gender                   string    `json:"gender" validate:"required,oneof=male female other"` // 性别 / 性別（せいべつ）：male / female / other
	BirthPlace               string    `json:"birth_place" validate:"omitempty"`                   // 出生地 / 出生地（しゅっしょうち）
	Address                  string    `json:"address" validate:"omitempty"`                       // 地址 / 住所（じゅうしょ）
	EmergencyContactName     string    `json:"emergency_contact_name" validate:"omitempty"`        // 紧急联系人姓名 / 緊急連絡人の氏名（きんきゅうれんらくにん）
	EmergencyContactPhone    string    `json:"emergency_contact_phone" validate:"omitempty,e164"`  // 紧急联系人电话 / 緊急連絡先電話番号
	EmergencyContactRelation string    `json:"emergency_contact_relation" validate:"omitempty"`    // 与紧急联系人的关系 / 緊急連絡人との関係（かんけい）
	EmergencyContactAddress  string    `json:"emergency_contact_address" validate:"omitempty"`     // 紧急联系人地址 / 緊急連絡先の住所
}

func NewRegisterStudentRequestFromGRPC(req *user.RegisterStudentRequest) *RegisterStudentRequest {
	enrollmentDatate := time.Unix(req.GetEnrollmentDate(), 0)
	graduationDate := time.Unix(req.GetGraduationDate(), 0)
	birthday := time.Unix(req.GetBirthday(), 0)

	return &RegisterStudentRequest{
		Name:                     req.GetName(),
		NameKana:                 req.GetNameKana(),
		Avatar:                   req.GetAvatar(),
		Email:                    req.GetEmail(),
		Phone:                    req.GetPhone(),
		Password:                 req.GetPassword(),
		DeviceID:                 req.GetDeviceId(),
		StudentNumber:            req.GetStudentNumber(),
		Department:               req.GetDepartment(),
		Major:                    req.GetMajor(),
		EnrollmentDate:           enrollmentDatate,
		GraduationDate:           graduationDate,
		Birthday:                 birthday,
		Gender:                   req.GetGender(),
		BirthPlace:               req.GetBirthPlace(),
		Address:                  req.GetAddress(),
		EmergencyContactName:     req.GetEmergencyContactName(),
		EmergencyContactPhone:    req.GetEmergencyContactPhone(),
		EmergencyContactRelation: req.GetEmergencyContactRelation(),
		EmergencyContactAddress:  req.GetEmergencyContactAddress(),
	}
}
