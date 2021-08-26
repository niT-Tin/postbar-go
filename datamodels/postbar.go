package datamodels

type PostBar struct {
	ID         int64  `bson:"id"`
	PostBarId  int64  `bson:"postbar_id"`
	PCount     int64  `bson:"pcount"`
	MainId     int64  `bson:"main_id"`
	NewInc     int64  `bson:"new_inc"`
	PostAvatar string `bson:"post_avatar"`
}
