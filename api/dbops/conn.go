package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"//仅需要引用
)

var (
	dbConn *sql.DB
	err error
)

//连接的标准方法
/*func openConn() *sql.DB {
	//此处的driver已经实现对应database/mysql内的interface
	//后续配置标准化

	//实质上open调用并不会直接连接，而是在内部实现
	dbConn, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/videos_info?charset=utf8")

	if err != nil {
		panic(err.Error())//注意此处的panic【对应内核的panic错误机制，所以需要加对应的日志和recover的重试机制】
	}

	return dbConn
}*/

//标准化当前包初始化，方便当前包内的其他方法的直接使用
func init(){
	//docker容器之间的连接自然需要宿主的docker-machine进行连接
	dbConn, err = sql.Open("mysql", "root:123456@tcp(172.17.0.4:3306)/videos_info?charset=utf8")

	if err != nil {
		panic(err.Error())
	}

	//当前变量就可以直接在包内使用
}
