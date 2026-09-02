package database

import (
	"fmt"

	"time"

	gsb "github.com/huandu/go-sqlbuilder"
)

func SqlInsert() {
	insertBuilder := gsb.NewInsertBuilder()
	insertBuilder = insertBuilder.InsertInto("student").Cols("name", "province", "city", "enrollment").Values("小明", "河南", "郑州", "2015-01-01") //除InsertInto外也支持ReplaceInto
	for i := 0; i < 3; i++ {
		insertBuilder = insertBuilder.Values(RandStringRunes(4), "河南", "郑州", time.Now().Add(time.Hour*24*time.Duration(i)).Format("2006-01-02"))
	}
	sql, args := insertBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}

func SqlDelete() {
	deleteBuilder := gsb.NewDeleteBuilder()
	deleteBuilder = deleteBuilder.DeleteFrom("student").Where(deleteBuilder.Equal("city", "郑州"))
	sql, args := deleteBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}

func SqlRead() {
	selectBuilder := gsb.NewSelectBuilder()
	selectBuilder.SetFlavor(gsb.MySQL)
	selectBuilder = selectBuilder.Select("name", "province").From("student")
	selectBuilder = selectBuilder.Where("score<80")
	selectBuilder = selectBuilder.Where(
		selectBuilder.Or(
			selectBuilder.Equal("province", "河南"),
			selectBuilder.GE("enrollment", "2015-01-01"),
		),
		selectBuilder.In("city", "郑州", "开封", "洛阳"),
	)
	selectBuilder = selectBuilder.OrderByDesc("name")
	selectBuilder = selectBuilder.Offset(10).Limit(3)
	sql, args := selectBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}

func SqlUpdate() {
	updateBuilder := gsb.NewUpdateBuilder()
	updateBuilder = updateBuilder.Update("student").Set("name=libai", "score=100").Where("score=0")
	sql, args := updateBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}
