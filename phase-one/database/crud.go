package database

import (
	"database/sql"
	"fmt"
)

// insert 插入数据
func Insert(db *sql.DB) {
	// 一条sql，插入2行记录
	res, err := db.Exec("insert into student (name,province,city,enrollment) values ('小明', '广东省', '深圳市', '2026-09-01'),('小红','上海市','浦东新区','2026-09-01')")
	CheckError(err)
	lastId, err := res.LastInsertId() // ID自增，用过的id（即使对应的行已delete）不会重复使用。如果使用单个INSERT语句将多行插入到表中，则LastInsertId是第一条数据使用的id
	CheckError(err)
	fmt.Printf("after insert last id %d\n", lastId)
	rows, err := res.LastInsertId() // ID自增，用过的id（即使对应的行已delete）不会重复使用。如果使用单个INSERT语句将多行插入到表中，则LastInsertId是第一条数据使用的id
	CheckError(err)
	fmt.Printf("after insert last id %d\n", lastId)
	rows, err = res.RowsAffected() // 插入2行，所以影响了2行
	CheckError(err)
	fmt.Printf("insert affect %d row\n", rows)
}

// replace 插入（覆盖）数据
func Replace(db *sql.DB) {
	// 由于name字段上有唯一索引，insert重复的name会报错。而使用replace会先删除，再插入
	res, err := db.Exec("replace into student (name,province,city,enrollment) values ('小明','深圳','深圳','2026-09-01'),('小红','上海','上海','2026-09-01')")
	CheckError(err)
	lastId, err := res.LastInsertId() // ID自增，用过的id（即使对应的行已delete）不会重复使用
	CheckError(err)
	fmt.Printf("after insert last id %d\n", lastId)
	rows, err := res.RowsAffected() // 先删除，后插入，影响了4行
	CheckError(err)
	fmt.Printf("insert affect %d row\n", rows)
}
