package datamodels

type Content struct {
	ID           int64  `bson:"id"`
	ContentId    int64  `bson:"content_id"`
	Class        int64  `bson:"class"`
	CommentId    int64  `bson:"comment_id"`
	PosterId     int64  `bson:"post_id"`
	TextContent  string `bson:"text_content"`
	PicContent   string `bson:"pic_content"`
	VideoContent string `bson:"video_content"`
}
