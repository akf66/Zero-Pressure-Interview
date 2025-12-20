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
			"../../sql",
		},
	}))

	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "../../sqlgen",
		ModelPkgPath:      "sqlgen",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    true,
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
	fieldopt := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField}

	user := g.GenerateModel("users", append(fieldopt, softDeleteField)...)

	g.ApplyBasic(user)
	g.Execute()
}
