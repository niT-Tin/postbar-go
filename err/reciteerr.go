package err

import "log"

func Reciteerr(err *error) bool {
	if *err != nil {
		log.Printf("error opening database: %v", err)
		// TODO: 将错误信息写入RabbitMQ队列，进而将错误消息一个一个插入mysql数据库
		return true
	}
	return false
}
