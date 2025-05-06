package initialize

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"Go/global"
	"Go/internal/po"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
        panic(err)
    }
    
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password,m.Host, m.Port, m.Dbname) // trả về chuỗi bằng Sprintf
  	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err,"InitMySql initialization error")
	global.Logger.Info("MySQL initialization sucess")
	global.Mdb = db

	//SetPool
	SetPool()
	// migrateTables()
	
	// genTableDAO()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns)) // khi có nhiều kết nối nhàn rỗi thì thực hiện kết nối mới nhanh hơn
	sqlDb.SetMaxOpenConns(m.MaxOpenConns) // giới hạn kết nối số lượng tối đa, tránh tình trạng quá tải
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) // sau khi kết nối tồn tại thì sẽ bị đóng
}

func migrateTables() {
	global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
    global.Logger.Info("Migration success")

}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	})
	
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	g.GenerateModel("go_crm_user")
	
	//   // Generate basic type-safe DAO API for struct `model.User` following conventions
	//   g.ApplyBasic(model.User{})
	
	//   // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
	
	  // Generate the code
	  g.Execute()
}