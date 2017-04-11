package db
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
  "fmt"
	_ "net/http"
	_ "time"
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
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "", "127.0.0.1", 3306, "sdp")
    mysqlConn, err = gorm.Open("mysql", connectionString)
    defer mysqlConn.Close()
    if err != nil {
      panic(err)
    }
    err = mysqlConn.DB().Ping()
    if err != nil {
      panic(err)
    }
    mysqlConn.LogMode(true)
    // mysqlConn.DB().SetMaxIdleConns(mysql.MaxIdleConns)
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}
