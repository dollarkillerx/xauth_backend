package models

import "time"

type Attendance struct {
	BaseModel

	UserID uint64 `gorm:"index;not null" json:"user_id"`
	// 学生ID / 学生ID（Userテーブル外部キー）

	Type string `gorm:"type:varchar(20);not null" json:"type"`
	// 打卡类型（entry=入校、exit=退校、class=上课签到）/ 打刻タイプ（entry＝登校、exit＝下校、class＝授業出席）

	CourseID *string `gorm:"type:varchar(255);index" json:"course_id"`
	// 课程ID（只有上课打卡时填写，入校/退校为空）/ 講義ID（授業打刻の場合のみ）

	Status string `gorm:"type:varchar(20);not null" json:"status"`
	// 出勤状态（present=出席、absent=缺席、late=迟到）/ 出席状況（present＝出席、absent＝欠席、late＝遅刻）

	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
	// 打卡时间 / 打刻時間

	DeviceID string `gorm:"type:varchar(255)" json:"device_id"`
	// 打卡设备编号 / 打刻端末ID

	Location string `gorm:"type:varchar(255)" json:"location"`
	// 打卡地点描述（例：北门、教学楼A）/ 打刻場所（例：北門、講義棟A）

	BluetoothTerminalID string `gorm:"type:varchar(255)" json:"bluetooth_terminal_id"`
	// 蓝牙终端ID（如蓝牙门禁、信标设备ID）/ Bluetooth端末ID（例：ビーコン端末ID）
}
