package models

import (
	"gorm.io/gorm"
)

// Comment 评论
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;comment:评论内容" json:"content"  form:"content" binding:"required"`  // 评论内容
	Score   int    `gorm:"comment:评论评分" json:"score,omitempty" form:"score" binding:"required"`       // 评论评分
	DishID  uint   `gorm:"comment:评论所属菜品id" json:"dishID,omitempty" form:"dishID" binding:"required"` // 评论所属菜品id
}
