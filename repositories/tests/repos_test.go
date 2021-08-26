package tests

import (
	"errors"
	"fmt"
	"log"
	"postbar/datamodels"
	"postbar/db"
	"postbar/db/mongodb"
	"postbar/repositories"
	"testing"
	"time"
)

var (
	c  repositories.IContent
	cm repositories.IComment
	sp repositories.ISinglePostRepository
	pb repositories.IPostBar
	u  repositories.IUserRepository

	data    datamodels.Content
	comment datamodels.Comment
	sigp    datamodels.SinglePost
	pstb    datamodels.PostBar
	usr     datamodels.User
)

func getc() {
	client := mongodb.GetClient()
	c = repositories.NewContentRepository(db.MongoDBName, db.ContentCollectionName, client)
	cm = repositories.NewCommentRepository(db.MongoDBName, db.CommentCollectionName, client)
	sp = repositories.NewSinglePostRepository(db.MongoDBName, db.SinglePostCollectionName, client)
	pb = repositories.NewPostBarRepository(db.MongoDBName, db.PostBarCollectionName, client)
	u = repositories.NewUserRepository(db.MongoDBName, db.UserCollectionName, client)
	fmt.Println("c = ", c)
}

func printErr(t *testing.T, err error, ts string) {
	if err != nil {
		t.Errorf(ts+": %v", err)
	}
}

func TestMain(m *testing.M) {
	getc()
	data = datamodels.Content{
		Class:        0,
		CommentId:    13,
		PosterId:     10,
		TextContent:  "This is a content",
		PicContent:   "http://test_pic_url",
		VideoContent: "http://test_video_url",
	}

	comment = datamodels.Comment{
		PublisherId: 1,
		SubId:       0,
		ParentId:    0,
		PublishDate: time.Now(),
		ReplyNum:    100,
		Contents: datamodels.Content{
			Class:        0,
			CommentId:    comment.CommentId,
			PosterId:     10,
			TextContent:  "This is a content",
			PicContent:   "http://test_pic_url",
			VideoContent: "http://test_video_url",
		},
		LikeNums: 0,
	}

	sigp = datamodels.SinglePost{
		PosterId: 1,
		Comments: []datamodels.Comment{
			comment,
			comment,
			comment,
			comment,
		},
	}

	pstb = datamodels.PostBar{
		PCount:     1,
		MainId:     1,
		NewInc:     1,
		PostAvatar: "http://post_avatar",
	}

	usr = datamodels.User{
		NickName:         "nickname",
		AvatarUrl:        "http://head_url",
		Account:          "test",
		Password:         "123456",
		Userid:           0,
		Subscribers:      []int64{1, 2, 3},
		SubscribePB:      []int64{12, 13},
		LikeNum:          120,
		CollectedPostIds: []int64{1, 2, 3, 4},
		PostCount:        14,
		Age:              3.3,
		Grade:            5,
		Exp:              1234,
		Ps:               []int64{1, 2, 3},
	}

	m.Run()
}

func TestContentWorkFlow(t *testing.T) {
	t.Run("create", testContentRepository_Create)
	t.Run("get_all", testContentRepository_GetAll)
	t.Run("update", testContentRepository_Update)
	t.Run("get_one_by_id", testContentRepository_GetOneById)
	t.Run("delete", testContentRepository_Delete)
}

func TestCommentWorkFlow(t *testing.T) {
	t.Run("insert", testCommentRepository_Insert)
	t.Run("get_one_by_id", testCommentRepository_GetOneById)
	t.Run("update", testCommentRepository_Update)
	t.Run("get_content_in_comment_by_id", testCommentRepository_GetContentInCommentById)
	t.Run("increase_like_num_by_one", testCommentRepository_IncreaseLikeNumByOne)
	t.Run("increase_like_num", testCommentRepository_IncreaseLikeNum)
	t.Run("delete", testCommentRepository_Delete)
}

func TestSinglePostWorkFlow(t *testing.T) {
	t.Run("create", testSinglePostRepository_Create)
	t.Run("get_one_by_id", testSinglePostRepository_GetOneById)
	t.Run("get_comments_in_post_by_post_id", testSinglePostRepository_GetCommentsInPostByPostId)
	t.Run("update", testSinglePostRepository_Update)
	t.Run("dalete", testSinglePostRepository_Delete)
}

func TestPostBarWorkFlow(t *testing.T) {
	t.Run("create", testPostBarRepository_Create)
	t.Run("get_one_by_id", testPostBarRepository_GetOneById)
	t.Run("get_all_post_bar", testPostBarRepository_GetAllPostBar)
	t.Run("update", testPostBarRepository_Update)
	t.Run("delete", testPostBarRepository_Delete)
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("insert", testUserRepository_Insert)
	t.Run("get_one_by_id", testUserRepository_GetOneById)
	t.Run("get_all_users", testUserRepository_GetAllUsers)
	t.Run("update", testUserRepository_Update)
	t.Run("get_user_by_account", testUserRepository_GetUserByAccount)
	t.Run("delete", testUserRepository_Delete)
}

func testCommentRepository_Insert(t *testing.T) {
	tp, err := cm.Insert(&comment)
	log.Printf("%+v", tp)
	printErr(t, err, "insertComment")
}

func testCommentRepository_GetOneById(t *testing.T) {
	tp, err := cm.GetOneById(comment.CommentId)
	log.Printf("%+v", tp)
	printErr(t, err, "get_one_by_idComment")
}

func testCommentRepository_Update(t *testing.T) {
	comment.ReplyNum = 2333
	tp, err := cm.Update(&comment)
	log.Printf("%+v", tp)
	printErr(t, err, "updateComment")
}

func testCommentRepository_GetContentInCommentById(t *testing.T) {
	tp, err := cm.GetContentInCommentById(comment.CommentId)
	log.Printf("%+v", tp)
	printErr(t, err, "get_content_in_comment_by_idComment")
}

func testCommentRepository_IncreaseLikeNumByOne(t *testing.T) {
	tp := cm.IncreaseLikeNumByOne(1)
	if !tp {
		printErr(t, errors.New("incorrect increase_like_num_by_one"), "increase_like_num_by_oneComment")
	}
}

func testCommentRepository_IncreaseLikeNum(t *testing.T) {
	tp := cm.IncreaseLikeNum(1, 325)
	if !tp {
		printErr(t, errors.New("incorrect increase_like_num_by_one"), "increase_like_numComment")
	}
}

func testCommentRepository_Delete(t *testing.T) {
	err := cm.Delete(&comment)
	printErr(t, err, "deleteComment")
}

func testSinglePostRepository_Create(t *testing.T) {
	err := sp.Create(&sigp)
	printErr(t, err, "createSinglePost")
}

func testSinglePostRepository_GetOneById(t *testing.T) {
	tp, err := sp.GetOneById(sigp.SinglePostId)
	log.Printf("%+v", tp)
	printErr(t, err, "get_one_by_idSinglePost")
}

func testSinglePostRepository_GetCommentsInPostByPostId(t *testing.T) {
	tp, err := sp.GetCommentsInPostByPostId(sigp.SinglePostId)
	log.Printf("%+v", tp)
	printErr(t, err, "get_comments_in_post_by_post_idSinglePost")
}

func testSinglePostRepository_Update(t *testing.T) {
	sigp.ID = 2333
	err := sp.Update(&sigp)
	printErr(t, err, "updateSinglePost")
}

func testSinglePostRepository_Delete(t *testing.T) {
	err := sp.Delete(sigp.SinglePostId)
	printErr(t, err, "deleteSinglePost")
}

func testPostBarRepository_Create(t *testing.T) {
	err := pb.Create(&pstb)
	printErr(t, err, "createPostBar")
}

func testPostBarRepository_GetOneById(t *testing.T) {
	tp, err := pb.GetOneById(pstb.PostBarId)
	log.Printf("%+v", tp)
	printErr(t, err, "get_one_by_idPostBar")
}

func testPostBarRepository_GetAllPostBar(t *testing.T) {
	bars, err := pb.GetAllPostBar()
	log.Printf("%+v", bars)
	printErr(t, err, "get_all_post_barPostBar")
}

func testPostBarRepository_Update(t *testing.T) {
	pstb.ID = 1234567
	err := pb.Update(&pstb)
	printErr(t, err, "updatePostBar")
}

func testPostBarRepository_Delete(t *testing.T) {
	err := pb.Delete(pstb.PostBarId)
	printErr(t, err, "deletePostBar")
}

func testUserRepository_Insert(t *testing.T) {
	err := u.Insert(&usr)
	printErr(t, err, "insertUser")
}

func testUserRepository_GetOneById(t *testing.T) {
	id := u.GetOneById(usr.Userid)
	log.Printf("%+v", id)
	if id == nil {
		printErr(t, errors.New("error in get_one_by_idUser"), "error in get_one_by_idUser")
	}
}

func testUserRepository_GetAllUsers(t *testing.T) {
	users := u.GetAllUsers()
	log.Printf("%+v", users)
	if users == nil {
		printErr(t, errors.New("error in get_all_users"), "get_all_usersUser")
	}
}

func testUserRepository_Update(t *testing.T) {
	usr.ID = 9876
	err := u.Update(&usr)
	printErr(t, err, "updateUser")
}

func testUserRepository_GetUserByAccount(t *testing.T) {
	tp, err := u.GetUserByAccount(usr.Account)
	log.Printf("%+v", tp)
	printErr(t, err, "get_user_by_accountUser")
}

func testUserRepository_Delete(t *testing.T) {
	err := u.Delete(usr.Userid)
	printErr(t, err, "deleteUser")
}

func testContentRepository_Create(t *testing.T) {
	create, err := c.Create(&data)
	log.Println(create)
	printErr(t, err, "createContent")
}

func testContentRepository_Delete(t *testing.T) {
	result, err := c.Delete(&data)
	log.Println(result)
	printErr(t, err, "deleteContent")
}

func testContentRepository_GetAll(t *testing.T) {
	all, err := c.GetAll()
	log.Println(all)
	printErr(t, err, "get_allContent")

}

func testContentRepository_Update(t *testing.T) {
	dataNew := datamodels.Content{
		ContentId:    1,
		Class:        1,
		CommentId:    13,
		PosterId:     11,
		TextContent:  "This is a content Update",
		PicContent:   "http://_pic_url",
		VideoContent: "http://_video_url",
	}
	update, err := c.Update(&dataNew)
	log.Println(update)
	printErr(t, err, "updateContent")
}

func testContentRepository_GetOneById(t *testing.T) {
	id, err := c.GetOneById(13)
	log.Println(id)
	printErr(t, err, "get_one_by_idContent")
}
