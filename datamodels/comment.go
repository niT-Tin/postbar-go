package datamodels

import "time"

type Comment struct {
	ID          int64     `bson:"id"`
	CommentId   int64     `bson:"comment_id"`
	PublisherId int64     `bson:"publisher_id"`
	SubId       int64     `bson:"sub_id"`
	ParentId    int64     `bson:"parent_id"`
	PublishDate time.Time `bson:"publish_date"`
	ReplyNUm    int64     `bson:"reply_num"`
	Contents    Content   `bson:"contents"`
	LikeNums    int64     `bson:"like_num"`
}
