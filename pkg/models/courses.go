package models

import (
	"time"
)

type Courses struct {
	BaseModel
	CourseCode string `gorm:"type:varchar(50);uniqueIndex;not null" json:"course_code"`
	// 课程ID（课程编号）/ 講義ID（コード）

	CourseName string `gorm:"type:varchar(255);not null" json:"course_name"`
	// 课程名称 / 講義名

	TeacherID uint64 `gorm:"not null" json:"teacher_id"`
	// 授课教师ID（Teachers表关联）/ 担当教員ID（Teachersテーブル外部キー）

	TeacherName string `gorm:"type:varchar(255);not null" json:"teacher_name"`
	// 授课教师姓名（冗余存储）/ 担当教員名（冗長保存）

	StartTime time.Time `gorm:"not null" json:"start_time"`
	// 上课开始时间（时分，比如14:00）/ 授業開始時間（時刻）

	EndTime time.Time `gorm:"not null" json:"end_time"`
	// 上课结束时间（时分，比如15:00）/ 授業終了時間（時刻）

	Period string `gorm:"type:varchar(20);not null" json:"period"`
	// 上课时间分类（如：1限、2限）/ 時限区分（例：1限、2限）

	Weekday int `gorm:"not null" json:"weekday"`
	// 星期几（1=周一，7=周日）/ 曜日（1＝月曜日、7＝日曜日）

	Credit int `gorm:"default:2;not null" json:"credit"`
	// 学分数（例：2学分）/ 単位数（例：2単位）

	Capacity int `gorm:"default:50;not null" json:"capacity"`
	// 最大人数限制 / 定員（上限人数）

	Location string `gorm:"type:varchar(255)" json:"location"`
	// 上课地点（教室名称）/ 授業場所（教室名）

	Description string `gorm:"type:text" json:"description"`
	// 课程简介 / 講義説明

	Remark string `gorm:"type:text" json:"remark"`
	// 备注（其他补充信息，可选）/ 備考（自由記述欄、任意）
}
