package models

import (
	"time"
)

type Notification struct {
	BaseModel
	Title string `gorm:"type:varchar(255);not null" json:"title"`
	// 通知标题 / 通知タイトル

	Content string `gorm:"type:text;not null" json:"content"`
	// 通知内容 / 通知内容

	TargetType string `gorm:"type:enum('broadcast','unicast');not null" json:"target_type"`
	// 发送类型（broadcast=广播，unicast=单播）/ 送信種別（broadcast＝全体、unicast＝個別）

	TargetUserID *uint64 `gorm:"index" json:"target_user_id"`
	// 单播时目标用户ID（广播时为空）/ 個別送信時の対象ユーザーID（全体送信ならNULL）

	SentAt time.Time `gorm:"autoCreateTime" json:"sent_at"`
	// 发送时间 / 送信日時
}
