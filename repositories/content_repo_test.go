package repositories

import (
	"fmt"
	"log"
	"postbar/datamodels"
	"postbar/db"
	"postbar/db/mongodb"
	"testing"
)

var (
	c    IContent
	data datamodels.Content
)

func getc() {
	client := mongodb.GetClient()
	c = NewContentRepository(db.MongoDBName, "content", client)
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
		ContentId:    1,
		Class:        0,
		CommentId:    13,
		PosterId:     10,
		TextContent:  "This is a content",
		PicContent:   "http://test_pic_url",
		VideoContent: "http://test_video_url",
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

func testContentRepository_Create(t *testing.T) {
	create, err := c.Create(&data)
	log.Println(create)
	printErr(t, err, "create")
}

func testContentRepository_Delete(t *testing.T) {
	result, err := c.Delete(&data)
	log.Println(result)
	printErr(t, err, "delete")
}

func testContentRepository_GetAll(t *testing.T) {
	all, err := c.GetAll()
	log.Println(all)
	printErr(t, err, "get_all")

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
	printErr(t, err, "update")
}

func testContentRepository_GetOneById(t *testing.T) {
	id, err := c.GetOneById(13)
	fmt.Println("T-------------------------------T")
	log.Println(id)
	fmt.Println("B-------------------------------B")
	printErr(t, err, "get_one_by_id")
}
