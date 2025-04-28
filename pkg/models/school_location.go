package models

type SchoolLocation struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	// 地点名称（如 校门口、教学楼）/ 場所名（例：正門、講義棟）

	Latitude float64 `gorm:"not null" json:"latitude"`
	// 纬度 / 緯度

	Longitude float64 `gorm:"not null" json:"longitude"`
	// 经度 / 経度

	Radius int `gorm:"not null" json:"radius"`
	// 允许打卡的半径范围（米）/ 打刻許可半径（メートル単位）
}
