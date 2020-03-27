package dbops

import (
	"fmt"
	"testing"
)

//对数据库的单元测试：
//dblogin
//truncate table
//run tests
//clear data

func TestMain(m *testing.M){
	clearData()
	fmt.Println("user test start")
	m.Run()
	fmt.Println("user test end")
	clearData()
}

//进行sub test工作流组织
func TestUserWorkFolw(t *testing.T) {
	t.Run("add", testAddUserCredential);
	t.Run("select", testGetUserCredential)
}

func clearData(){
	dbConn.Exec("TRUNCATE TABLE user")
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("test", "test_pwd")

	if err != nil {
		t.Errorf("add user is err:%v", err)
	}
}

func testGetUserCredential(t *testing.T) {

	pwd, err := GetUserCredential("test")

	if err != nil || pwd == "" {
		t.Errorf("select user is err: %v", err)
		t.Errorf("select user value is err: %s", pwd)
	}
}