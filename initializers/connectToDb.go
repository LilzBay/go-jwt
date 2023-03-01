package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	dsn := os.Getenv("dsn") // 初始化时`godotenv`已经获得了`.env`中的环境变量配置
	// DB和err事先声明，此处仅需要赋值！
	// 此处绝对不能使用`:=`，会导致DB重新初始化，空指针nil引用
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// fmt.Println("database:", DB)

	if err != nil {
		log.Fatal(err)
	}
}
