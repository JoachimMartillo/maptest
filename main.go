package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"os"
	"time"
)
// This stuff all initializes the Beego Orm database subsystem.
func init() {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// registerModels()

	if _, err := orm.GetDB(); err != nil {
		// Connect default DB
		driverName := os.Getenv("DB_DRIVER")
		dataSource := os.Getenv("DB_SOURCE")
		if driverName == "" {
			driverName = beego.AppConfig.String("dbDriver")
		}
		if dataSource == "" {
			dataSource = beego.AppConfig.String("dbSource")
		}
		maxIdle, _ := beego.AppConfig.Int("dbMaxIdle")
		maxConn, _ := beego.AppConfig.Int("dbMacConn")
		orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxConn)

		orm.DefaultTimeLoc = time.UTC
	}

	orm.Debug, _ = beego.AppConfig.Bool("ormDebug")
}


//var (
//	fileExcludedMap = make(map[string]int) // keeping track of excluded files
//	// it's useful for making log less noisy & for (Prometheus) monitoring
//)

func main() {
	//var count int;
	//count = fileExcludedMap["FirstFile.text"]
	//println(fmt.Sprintf("count is %d", count))
	////fileExcludedMap["FirstFile.text"] = 0;
	//println(fmt.Sprintf("value is %d", fileExcludedMap["FirstFile.text"]))
	//count = count + 2
	//fileExcludedMap["FirstFile.text"] = count
	//println(fmt.Sprintf("new value is %d", fileExcludedMap["FirstFile.text"]))

	
}
