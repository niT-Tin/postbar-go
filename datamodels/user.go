package datamodels

type User struct {
	ID          int64   `bson:"db_id"`        //数据库id
	NickName    string  `bson:"nick_name"`    //用户昵称
	AvatarUrl   string  `bson:"avatar_url"`   //头像url
	Account     string  `bson:"account"`      //帐号
	Password    string  `bson:"password"`     //密码
	Userid      int64   `bson:"userid"`       //用户id
	Subscribers []int64 `bson:"subscribers"`  //关注的人数id
	SubscribePB []int64 `bson:"subscribe_pb"` //关注的帖子
	LikeNum     int64   `bson:"like_num"`     //获得的点赞数

	CollectedPostIds []int64 `bson:"collected_post_ids"` //收藏的贴吧id
	PostCount        int64   `bson:"post_count"`         //发出的帖子总数
	Age              float32 `bson:"age"`                //吧龄
	Grade            int64   `bson:"grade"`              //等级
	Exp              int64   `bson:"exp"`                //当前经验
	Ps               []int64 `bson:"ps"`                 //身为吧主对应贴吧的id
}
