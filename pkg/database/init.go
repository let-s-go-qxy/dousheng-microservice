package database

import (
	"context"
	g "dousheng/pkg/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Config struct {
	Database Database
}

type Database struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type Mysql struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	Addr        string `yaml:"addr"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	DbVerify    int    `yaml:"dbVerify"`
	DbUserLike  int    `yaml:"dbUserLike"`
	DbVideoLike int    `yaml:"dbVideoLike"`
}

// InitDB 初始化数据库相关
func InitDB() {
	m, r := ParseYaml()
	MysqlDBSetup(m)
	RedisSetup(r)
}

// InitMysql 如果用不到redis，使用这个方法只连接mysql
func InitMysql() {
	m, _ := ParseYaml()
	MysqlDBSetup(m)
}

// ParseYaml 解析yaml
func ParseYaml() (mySql Mysql, redis Redis) {
	path, _ := os.Getwd()
	var configViperConfig = viper.New()
	configViperConfig.SetConfigName("config")
	configViperConfig.AddConfigPath(path + "/conf/")
	configViperConfig.SetConfigType("yaml")
	//读取配置文件内容
	if err := configViperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	var c Config
	if err := configViperConfig.Unmarshal(&c); err != nil {
		panic(err)
	}
	redis = c.Database.Redis
	mySql = c.Database.Mysql
	return
}

func MysqlDBSetup(m Mysql) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Australia%%2FMelbourne",
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.Db,
		m.Charset)), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetConnMaxLifetime(100 * time.Second)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("initialize mysql db successfully")
	g.MysqlDB = db
}

func RedisSetup(r Redis) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	verifyDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Addr, r.Port),
		Username: "",
		Password: r.Password,
		DB:       r.DbVerify,
		PoolSize: 10000,
	})
	_, err := verifyDb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	g.DbVerify = verifyDb

	fmt.Println("initialize verify redis client successfully")

	userLikeDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Addr, r.Port),
		Username: "",
		Password: r.Password,
		DB:       r.DbUserLike,
		PoolSize: 10000,
	})
	_, err = userLikeDb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	g.DbUserLike = userLikeDb

	fmt.Println("initialize userLike redis client successfully")

	videoLikeDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Addr, r.Port),
		Username: "",
		Password: r.Password,
		DB:       r.DbVideoLike,
		PoolSize: 10000,
	})
	_, err = videoLikeDb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	g.DbVideoLike = videoLikeDb
	//g.DbVideoLike.FlushAll(g.RedisContext)
	fmt.Println("initialize videoLike redis client successfully")
}
