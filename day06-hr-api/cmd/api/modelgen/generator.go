package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres123 dbname=hr_db_dev port=5432 sslmode=disable search_path=hr"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "internal/domain/query",
		ModelPkgPath:      "internal/models",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	g.WithTableNameStrategy(func(tableName string) (targetTableName string){
		if tableName == "schema_migrations" {
			return ""
		}
		return tableName
	})

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
