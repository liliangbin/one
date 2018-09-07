# one
自动爬取one的数据，并且同时存储到数据库，同时能够展示到我的主页，
### 使用的技术
- shell  脚本， 定时任务 crontab  Python scrapy
- go 主要是用到了数据库的存取，http服务器的建立

#### 启动
- go run insert.go   //数据的存取
- go build Web.go   ./Web   //http服务器的开启
- 
