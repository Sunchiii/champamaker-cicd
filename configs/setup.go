package configs

import (
	"log"

	"github.com/Sunchiii/champamker-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var AdminDB *gorm.DB
var MemberDB *gorm.DB

func ConnectDatebase() {
	admindb, err := gorm.Open(sqlite.Open("databases/admin.db"), &gorm.Config{})
	memberdb, memerr := gorm.Open(sqlite.Open("databases/member.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to admin database!")
	}
	if memerr != nil {
		log.Fatal("can't connect to member database!")
	}
	admindb.AutoMigrate(&models.Admin{})
	memberdb.AutoMigrate(&models.Member{})

	AdminDB = admindb
	MemberDB = memberdb
}
