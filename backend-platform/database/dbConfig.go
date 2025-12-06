package database

import (
	"lietcode/logic/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Config string
	Entity *[]interface{}
}

var DatabaseInstance *gorm.DB

func (db *Database) ConnectDatabase() *gorm.DB {
	dnsDatabase := db.Config
	databaseInstance, err := gorm.Open(mysql.Open(dnsDatabase), &gorm.Config{})

	if err != nil {
		log.Printf("❌ Cannot connect to MySQL: %v", err) // ⬅ STOP NGAY
	}

	log.Println("Connected to MySQL successfully")
	return databaseInstance
}
func (db *Database) MigrateDatabase(Database *gorm.DB) {
	if Database == nil {
		log.Printf("❌ Cannot run migrations: database connection is nil")
	}

	for _, entity := range *db.Entity {
		if err := Database.AutoMigrate(entity); err != nil {
			log.Printf("❌ AutoMigrate failed for entity %v: %v", entity, err)
		}
	}

	log.Println("✅ Database migrated successfully")
}
func InitDatabase(config string) *Database {
	database := Database{
		Config: config,
		Entity: &[]interface{}{
			&entity.User{},
			&entity.Submission{},
			&entity.Problem{},
			&entity.CodeTemplate{},
			&entity.ListProblem{},
			&entity.TestCase{},
		},
	}

	DatabaseInstance = database.ConnectDatabase()
	database.MigrateDatabase(DatabaseInstance)
	return &database
}
