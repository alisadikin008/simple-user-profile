package config

/*
	created by Ali Sadikin
	this file is configuration by viper that invoked config.json (env) in project directory
*/
import (
	general "simple-user-profile/general"

	_ "github.com/go-sql-driver/mysql" // _
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// LoadConfiguration -()
func LoadConfiguration() string {
	viperConfig()
	mode := viper.GetString("mode.name")
	return mode
}

// ConnectDB -()
func ConnectDB() (*gorm.DB, error) {
	viperConfig()
	DBDriver := viper.GetString("database.driver")
	DBName := viper.GetString("database.name")
	DBUser := viper.GetString("database.user")
	DBPassword := viper.GetString("database.password")
	DBHost := viper.GetString("tcp.host")
	DBPort := viper.GetString("tcp.port")
	db, err := gorm.Open(DBDriver, DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		general.WriteErrorLog(err.Error())
	}

	db.LogMode(true)
	db.Debug()
	return db, nil
}

// viperConfig -()
func viperConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		general.WriteErrorLog("the configuration may be not set clearly")
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// GetCredential -()
func GetCredential(credentialKey string) string {
	viperConfig()
	credential := ""
	if credentialKey == "key" {
		credential = viper.GetString("auth.key")
	} else if credentialKey == "token" {
		credential = viper.GetString("auth.token")
	}

	return credential
}
