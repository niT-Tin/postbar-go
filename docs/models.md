## models

**User** 使用mongodb

- ID 数据库内id
- nickname 昵称
- avatarurl 头像url
- account 帐号
- password 密码
- userid 用户id
- subscribers 关注的用户
- subscribePB 关注的贴吧
- likeNume 获得点赞数
- collectedPostids 收藏的贴子id（数组对象）
- postcount 发帖数
- age 吧龄
- grade 当前等级
- exp 当前经验值
- ps 身为吧主对应贴吧的id（数组）

**PostBar** 使用mongodb

- ID 数据库内id
- postid 贴吧id
- pcount 贴子数量
- mainid 吧主id
- newInc 每日新增贴数
- postAvatar 贴吧图表url

**SinglePost** 使用mongodb

- ID数据库id
- singlepostid 帖子id
- posterid 楼主id
- comments 评论**（数组对象）**

**Comment** 使用mongodb

- ID 数据库内id
- commentid 评论id
- publisherid 发布者id 指向user
- subid 子评论id
- parentid 父评论id
- publishdate 发布时间
- replynum 回复数
- content 内容 **（对象）**
- likenums 点赞数

**Content** 使用mongodb

- ID 数据库id
- contentid 内容id
- class 内容类别（属于评论类型，还是私信类型）
- commentid 父评论id
- posterid 发布者id
- textContent 文字内容
- picCOntent 图片url
- videoContent 视频内容

**Errors** 使用mysql

- 错误内容
- 错误时间
  - 使用rebbitmq将错误写入消息队列，后一个一个进入数据库

**Routers** 使用mysql

- 路由id
- 路由字符串



## 前台服务，

用户：注册，登陆，发帖，评论，点赞，查看自己详细信息，修改密码，昵称，头像等信息，关注贴吧，收藏帖子，关注其他用户，被回复提醒，被点赞提醒，显示吧龄，关注数，显示等级，经验值，拥有的贴吧（等级到达一定时才能申请吧主），创建贴吧（申请吧主成功后才能创建），

## 后台管理服务

对贴吧的增删改查， 对具体帖子的增删改查， 对具体帖子评论的增删改查， 设置吧主， 察看贴吧日在线人数，察看贴吧总吧数， 察看贴吧总贴数， 察看所有评论数量