package datamodels

import "gorm.io/gorm"

type IdVar struct {
	gorm.Model
	ContentIdInc    int64 `gorm:"content_id_inc"`
	CommentIdInc    int64 `gorm:"comment_id_inc"`
	SinglePostIdInc int64 `gorm:"single_post_id_inc"`
	PostBarIdInc    int64 `gorm:"post_bar_id_inc"`
	UserIdInc       int64 `gorm:"user_id_inc"`
}
