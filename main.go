package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	UserId   int `gorm:"primary_key"`
	Phone    string
	WxopenId string
	Tcreate  *time.Time
	Tprocess *time.Time
	Balance  int
	Src      string
	Level    int
}

type IpoBalance struct {
	ID               int             `gorm:"column:id;primary_key;auto_increment;not null"`
	S_INFO_WINDCODE  string          `gorm:"column:S_INFO_WINDCODE;unique_index:S_R"`     //Wind代码
	ANN_DT           string          `gorm:"column:ANN_DT"`                               //公告日期
	ACTUAL_ANN_DT    string          `gorm:"column:ACTUAL_ANN_DT"`                        //实际公告日期
	REPORT_PERIOD    string          `gorm:"column:REPORT_PERIOD;unique_index:S_R"`       //报告期
	STATEMENT_TYPE   string          `gorm:"column:STATEMENT_TYPE"`                       //报表类型
	MONETARY_CAP     sql.NullFloat64 `gorm:"column:MONETARY_CAP;type:decimal(20,4);"`     // 货币资金
	MONETARY_CAP_YOY sql.NullFloat64 `gorm:"column:MONETARY_CAP_YOY;type:decimal(20,4);"` // 货币资金
	OPDATE           int             `gorm:"column:OPDATE;type:int(10);"`
}

type Test struct {
	ID        int       `gorm:"column:id;primary_key;auto_increment;not null"`
	NAME      string    `gorm:"column:name;TYPE:VARCHAR(20);unique_index:S_R"`
	AGE       int       `gorm:"column:age;TYPE:int(10)"`
	SCORE     float32   `gorm:"column:score;TYPE:float(10,2)"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
}

type D struct {
	NAME  string  `gorm:"column:name;TYPE:VARCHAR(20);unique_index:S_R"`
	SCORE float32 `gorm:"column:score;TYPE:float(10,2)"`
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/test?charset=utf8&parseTime=true")
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&IpoBalance{}, &Test{})
	//data := &Test{
	//	NAME: "wangwu",
	//	AGE: 21,
	//	SCORE: 90.01,
	//	CreatedAt: time.Now(),
	//}
	//
	//if err := db.Model(&Test{}).Where("name=?","wangwu").Updates(&data).Error;err!=nil{
	//	fmt.Println("插入失败", err)
	//	return
	//}

	//temp := D{
	//	NAME: "zhangsan",
	//	SCORE: 90,
	//}
	//if err := db.Model(&Test{}).Where("name=?",temp.NAME).Updates(&temp).Error;err!=nil{
	//	fmt.Println("插入失败", err)
	//	return
	//}

	data1 := new(Test)

	if err := db.Model(&Test{}).Where("name=?", "zhangsan").Find(&data1).Error; err != nil {
		fmt.Println("获取失败", err)
		return
	}
	fmt.Println(*data1)

	data2 := new(D)
	if err := db.Raw("select name,score from test where name='zhangsan';").Scan(&data2).Error; err != nil {
		fmt.Println("获取失败", err)
		return
	}
	fmt.Println(*data2)

	var data3 []*D
	if err := db.Raw("select name,score from test;").Scan(&data3).Error; err != nil {
		fmt.Println("获取失败", err)
		return
	}

	for _, entity := range data3 {
		fmt.Println(*entity)
	}

}
