package models

import (
	"time"
)

// Comment 评论
type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt struct {
		Time  time.Time
		Valid bool // Valid is true if Time is not NULL
	} `gorm:"index"`
	Content string `gorm:"type:text;comment:评论内容"` // 评论内容
	Score   int    `gorm:"comment:评论评分"`           // 评论评分
	DishID  uint   `gorm:"comment:评论所属菜品id"`       // 评论所属菜品id
}
