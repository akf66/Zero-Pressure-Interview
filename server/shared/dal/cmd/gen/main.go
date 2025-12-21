package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

func main() {

	db, err := gorm.Open(rawsql.New(rawsql.Config{
		FilePath: []string{
			"./server/shared/dal/sql/users.sql",
			"./server/shared/dal/sql/interviews.sql",
			"./server/shared/dal/sql/questions.sql",
			"./server/shared/dal/sql/storage.sql",
		},
	}))

	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./server/shared/dal/sqlfunc",
		ModelPkgPath:      "./server/shared/dal/sqlentity",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	autoUpdateTimeField := gen.FieldGORMTag("updated_at", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "updated_at")
		tag.Set("type", "datetime")
		tag.Set("autoUpdateTime", "")
		return tag
	})
	autoCreateTimeField := gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "created_at")
		tag.Set("type", "datetime")
		tag.Set("autoCreateTime", "")
		return tag
	})

	softDeleteField := gen.FieldType("deleted_at", "gorm.DeletedAt")
	fieldopt := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}

	// 生成所有表
	tables := []string{"users", "interviews", "questions", "storage"}
	for _, tableName := range tables {
		model := g.GenerateModel(tableName, fieldopt...)
		g.ApplyBasic(model)
	}

	g.Execute()
}
