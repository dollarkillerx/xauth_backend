package models

import (
	"time"
)

type CourseEnrollment struct {
	BaseModel
	StudentID uint64 `gorm:"index;not null" json:"student_id"`
	// 学生ID / 学生ID（Studentテーブル外部キー）

	CourseID uint64 `gorm:"index;not null" json:"course_id"`
	// 课程ID / 講義ID（Coursesテーブル外部キー）

	Term string `gorm:"type:varchar(50);not null" json:"term"`
	// 学期（例：2024春季学期、2024秋季学期）/ 学期（例：2024年春学期）

	Status string `gorm:"type:enum('enrolled','withdrawn','pending');default:'pending';not null" json:"status"`
	// 选课状态（enrolled=已报名，withdrawn=已退课，pending=待审核）/ 受講状況（enrolled=受講中、withdrawn=履修取消、pending=審査待ち）

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 报名时间 / 登録日時
}
