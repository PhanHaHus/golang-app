package db
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "net/http"
)

var (
	mysqlConn *gorm.DB
	err       error
)
// initialize database
func init() {
		setupMysqlConn()
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
    mysqlConn, err = gorm.Open("mysql", "root:@/sdp?charset=utf8&parseTime=True")

    if err != nil {
      panic(err)
    }
    err = mysqlConn.DB().Ping()
    if err != nil {
      panic(err)
    }
    mysqlConn.LogMode(true)
		// mysqlConn.DB().SetMaxIdleConns(10)

    // mysqlConn.DB().SetMaxIdleConns(mysql.MaxIdleConns)
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}
