package mysql

import (
	"fmt"

	plugin "github.com/jsmzr/boot-plugin"
	"github.com/jsmzr/boot-plugin-gorm-mysql/db"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormMysqlPlugin struct{}

const configPrefix = "boot.gorm."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled":       true,
	"order":         4,
	"port":          3306,
	"singularTable": true,
	"noLowerCase":   false,
}

var replacer schema.Replacer

func (m *GormMysqlPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (m *GormMysqlPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (m *GormMysqlPlugin) Load() error {
	username := viper.GetString(configPrefix + "username")
	host := viper.GetString(configPrefix + "host")
	database := viper.GetString(configPrefix + "database")
	if username == "" || host == "" || database == "" {
		return fmt.Errorf("username, host and database should be not null")
	}
	password := viper.GetString(configPrefix + "password")
	port := viper.GetInt(configPrefix + "port")
	options := viper.GetString(configPrefix + "options")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, database, options)
	config := mysql.Config{DSN: dsn}

	conn, err := gorm.Open(mysql.New(config), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString(configPrefix + "tablePrefix"),
			SingularTable: viper.GetBool(configPrefix + "singularTable"),
			NoLowerCase:   viper.GetBool(configPrefix + "noLowerCase"),
			NameReplacer:  replacer,
		},
	})
	if err != nil {
		return err
	}
	db.DB = conn
	return nil
}

// 在 main.go 中设置
func SetReplacer(r schema.Replacer) {
	replacer = r
}

func init() {
	for key := range defaultConfig {
		viper.SetDefault(configPrefix+key, defaultConfig[key])
	}
	plugin.Register("mysql", &GormMysqlPlugin{})
}
