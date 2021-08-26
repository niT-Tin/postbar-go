package datamodels

type SinglePost struct {
	ID           int64     `bson:"id"`
	SinglePostId int64     `bson:"single_post_id"`
	PosterId     int64     `bson:"poster_id"`
	Comments     []Comment `bson:"comments"`
}
