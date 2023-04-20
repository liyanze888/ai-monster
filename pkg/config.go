package pkg

import (
	"log"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// IsProdEnv TODO
var IsProdEnv bool

// EnableMockUser TODO
var EnableMockUser bool

// Env TODO
var Env string

// IsDebug TODO
func IsDebug() bool {
	return viper.GetBool("debug")
}

//IsDev TODO
func IsDev() bool {
	return Env == "dev"
}

// GetString TODO
func GetString(key string) string {
	return viper.GetString(key)
}

// GetBool TODO
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetInt32 TODO
func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

//GetInt64Arr TODO
func GetInt64Arr(key string) []int64 {
	str := viper.GetString(key)
	if len(str) <= 0 {
		return make([]int64, 0)
	}
	var result []int64
	for _, value := range strings.Split(str, ",") {
		iVal, _ := strconv.ParseInt(value, 10, 64)
		result = append(result, iVal)
	}
	return result
}

//GetInt64Map TODO
func GetInt64Map(key string) map[int64]struct{} {
	result := make(map[int64]struct{}, 0)
	str := viper.GetString(key)
	if len(str) <= 0 {
		return result
	}
	for _, value := range strings.Split(str, ",") {
		iVal, _ := strconv.ParseInt(value, 10, 64)
		result[iVal] = struct{}{}
	}
	return result
}

// SetDefault TODO
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

func init() {
	viper.AutomaticEnv()
	viper.SetConfigFile(`config/config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to read cfg !!!!!!!!!!!!!!!!,err:%v", err)
		viper.Set("debug", true)
		viper.Set("env", "testing")
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode !!!!!!!!!")
	}
	log.Println("Init viper cfg")
	initEnv()
	initMockUser()
}

func initEnv() {
	Env = viper.GetString("env")
	noneProdEnvs := make(map[string]struct{}, 3)
	noneProdEnvs["dev"] = struct{}{}
	noneProdEnvs["test"] = struct{}{}
	noneProdEnvs["pre"] = struct{}{}
	if _, ok := noneProdEnvs[Env]; ok {
		log.Println("NOT_PRODUCTION ENV")
		IsProdEnv = false
	} else {
		log.Println("PRODUCTION ENV")
		IsProdEnv = true
	}
	log.Println("Config Env")
}

func initMockUser() {
	EnableMockUser = viper.GetBool("config.context.x_mock_user.enable")
	log.Printf("EnableMockUser: %v", EnableMockUser)
}
