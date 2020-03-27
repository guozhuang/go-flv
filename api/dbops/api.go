package dbops

import (
	"database/sql"
	"fmt"
	"time"
)

func AddUserCredential(loginName, pwdVal string) error {
	//sql:insert into user (name, pwd) values(?, ?)
	stmIns, err := dbConn.Prepare("insert into user (`name`, `pwd`,  `create_time`) values(?, ?, ?)")

	fmt.Println("123")
	if err != nil {
		return err
	}

	_, err = stmIns.Exec(loginName, pwdVal, Time())//todo：将标准函数标准化
	if err != nil {
		return err
	}

	defer stmIns.Close()//将前面有异常的

	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmSel , err := dbConn.Prepare("SELECT pwd FROM user WHERE name = ?")

	if err != nil {
		return "", err
	}

	//将对应的结构标准化:select查询的标准化，需要将结构体传入方法中
	var pwd string
	err = stmSel.QueryRow(loginName).Scan(&pwd)
	//对结果健壮化处理
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmSel.Close()

	return pwd, err
}

func Time() int64 {
	return time.Now().Unix()
}
