package models

import (
	"time"
)

type User struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	// 姓名 / 名前
	NameKana string `gorm:"type:varchar(255);not null" json:"name_kana"`
	// 假名（全角片假名表记）/ フリガナ（全角カタカナ表記）
	Avatar string `gorm:"type:varchar(255)" json:"avatar"`
	// 头像URL / アバター画像URL
	Email string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	// 邮箱（登录用，唯一）/ メールアドレス（ログイン用、一意）
	Phone string `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	// 手机号（登录用，唯一）/ 携帯電話番号（ログイン用、一意）
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	// 加密后的密码 / 暗号化されたパスワード
	Role string `gorm:"type:varchar(20);not null" json:"role"`
	// 角色类型（admin=管理员，student=学生，teacher=教师，guardian=监护人）/ 役割タイプ（admin=管理者、student=学生、teacher=教師、guardian=保護者）
	AuditStatus string `gorm:"type:varchar(20);default:'pending';not null" json:"audit_status"`
	// 审核状态（pending=待审核，pass=通过，reject=拒绝，disabled=禁用）/ 審査ステータス（pending=審査待ち、pass=承認、reject=却下、disabled=無効化）
	LastLoginAt time.Time `json:"last_login_at"`
	// 最后登录时间 / 最終ログイン日時
	LastLoginDeviceID string `gorm:"type:varchar(255)" json:"last_login_device_id"`
	// 最后登录设备ID / 最後にログインしたデバイスID
	ForbidChangeDevice bool `gorm:"default:true" json:"forbid_change_device"`
	// 禁止换设备登录（true=禁止，false=允许）/ デバイス変更禁止（true=禁止、false=許可）

}

type LoginLog struct {
	BaseModel
	UserID string `gorm:"index;not null" json:"user_id"`
	// 哪个用户登录 / ログインしたユーザーID
	DeviceID string `gorm:"type:varchar(255);not null" json:"device_id"`
	// 登录设备ID / ログインデバイスID
	IP string `gorm:"type:varchar(45)" json:"ip"`
	// 登录IP地址（支持IPv4/IPv6）/ ログインIPアドレス（IPv4/IPv6対応）
	UserAgent string `gorm:"type:varchar(512)" json:"user_agent"`
	// 浏览器或App的User-Agent信息 / ブラウザまたはアプリのUser-Agent情報
	LoginAt time.Time `gorm:"autoCreateTime" json:"login_at"`
	// 登录时间 / ログイン日時
}
