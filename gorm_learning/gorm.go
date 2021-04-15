package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// gorm连接数据库
	// db变量为*gorm.DB对象
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:mzbmzbmzb@/mzb1?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	// db.Xx

	// 操作
	type User struct {
		ID   int64
		Name string `gorm:"default:'小王子'"`
		Age  int64
	}
	user := User{Name: "q1mi", Age: 18}

	// C
	// 使用使用NewRecord()查询主键是否存在，主键为空使用Create()创建记录
	//fmt.Println(db.NewRecord(user)) // 主键为空返回`true`
	//db.Create(&user)   // 创建user  INSERT INTO users(name, age) values('q1mi', 18);  注意表名+s 是users
	//fmt.Println(db.NewRecord(user)) // 创建`user`后返回`false`

	// R
	// 查询所有的记录
	db.Debug().Find(&user, "name = ?", "miaozhibin")
	fmt.Println(user)  //
	// SELECT * FROM users where name='miaozhibin';

	var users []User
	db.Debug().Find(&users)
	fmt.Println(users)
	// select * from users;

	// U
	db.Debug().Model(User{}).Where("id = ?", 5).Updates(User{Name: "hello", Age: 18})
	// UPDATE `users` SET `age` = 18, `name` = 'hello'  WHERE (id = 5)
	db.Debug().Find(&users)
	fmt.Println(users)

	// D
	// 警告 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录
	db.Debug().Where("name = ?", "q1mi").Delete(User{})
	db.Debug().Find(&users)
	fmt.Println(users)

}