package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"path"
	"time"

	"github.com/go-kit/kit/sd/consul"
	"github.com/go-redis/redis"
	consuld "github.com/hashicorp/consul/api"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"userService/pkg/common"
	"userService/pkg/model"
	"userService/pkg/pb"
	"userService/pkg/rbac"
	"userService/pkg/userservice"
)

var (
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlUser     = "test"
	mysqlPassword = "test"
	mysqlDB       = "test"

	redisHost = "localhost"
	redisPort = 6379

	grpcHost = "localhost"
	grpcPort = 5001

	consulHost = "localhost"
	consulPort = 8500

	rbacFileName = ""

	logPath = ""
	logFile = ""
)

func main() {
	level := os.Getenv("LOG_LEVEL")
	if level == "debug" {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	var err error
	if err = parseConfigFile(); err != nil {
		logrus.Fatal("解析配置文件错误", err)
	}

	logrus.SetFormatter(&logFormatter{})
	if logFile != "" {
		os.MkdirAll(logPath, os.ModePerm)
		logFilePath := path.Join(logPath, logFile)
		writer, err := rotatelogs.New(
			logFilePath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(logFilePath),
			rotatelogs.WithMaxAge(time.Duration(24)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(30*24)*time.Hour),
		)
		if err != nil {
			logrus.Errorln(err)
		}
		logrus.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  writer,
				logrus.DebugLevel: writer,
				logrus.FatalLevel: writer,
				logrus.PanicLevel: writer,
				logrus.ErrorLevel: writer,
				logrus.WarnLevel:  writer,
				logrus.TraceLevel: writer,
			},
			&logFormatter{},
		))

	}

	// 初始化redis client
	common.RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisHost, redisPort),
	})

	// 初始化mysql client
	opts := model.Options{
		User:     mysqlUser,
		Password: mysqlPassword,
		DB:       mysqlDB,
		Addr:     fmt.Sprintf("%s:%d", mysqlHost, mysqlPort),
	}
	common.DB, err = model.NewDB(&opts)
	if err != nil {
		logrus.Fatal("启动mysql错误", err)
	}
	if level == "debug" {
		common.DB = common.DB.Debug()
	}

	common.Enforcer = rbac.NewCasbin(rbacFileName, &opts)

	// 初始化consul client
	consulClient, err := newConsulClient(fmt.Sprintf("%s:%d", consulHost, consulPort))
	if err != nil {
		logrus.Fatal("连接consul失败", err)
	}

	// 启动grpc server
	go func() {

		if err := runGRPCServer(fmt.Sprintf("%s:%d", "", grpcPort)); err != nil {
			logrus.Fatal("grpc server shutdown", err)
		}
	}()

	// 注册consul service
	err = registerConsulService(consulClient, "userService", grpcHost, grpcPort)
	if err != nil {
		logrus.Errorln("注册userService失败", err)
	}
	logrus.Infoln("启动成功")

	for {
		time.Sleep(time.Hour)
	}
}

func parseConfigFile() error {
	fileName := os.Getenv("CONFIG_FILE")

	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs/")

	if fileName != "" {
		viper.SetConfigFile(fileName)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	settings := viper.AllSettings()
	logrus.Infoln(settings)

	mysqlHost = viper.GetString("db.mysql.host")
	mysqlPort = viper.GetInt("db.mysql.port")
	mysqlUser = viper.GetString("db.mysql.user")
	mysqlPassword = viper.GetString("db.mysql.password")
	mysqlDB = viper.GetString("db.mysql.db")

	redisHost = viper.GetString("db.redis.host")
	redisPort = viper.GetInt("db.redis.port")

	grpcHost = viper.GetString("grpc.host")
	grpcPort = viper.GetInt("grpc.port")

	if host := os.Getenv("GRPC_HOST"); host != "" {
		grpcHost = host
	}

	consulHost = viper.GetString("consul.host")
	consulPort = viper.GetInt("consul.port")

	rbacFileName = viper.GetString("rbacFile")

	logPath = viper.GetString("log.path")
	logFile = viper.GetString("log.file")
	return nil
}

func newConsulClient(addr string) (consul.Client, error) {
	consulClient, err := consuld.NewClient(&consuld.Config{
		Address: addr,
	})
	if err != nil {
		return nil, err
	}
	return consul.NewClient(consulClient), nil
}

func registerConsulService(client consul.Client, name string, host string, port int) error {
	return client.Register(&consuld.AgentServiceRegistration{
		Address: host,
		Port:    port,
		Name:    name,
	})
}

func runGRPCServer(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	svr := grpc.NewServer()
	pb.RegisterUserServer(svr, userservice.New())
	return svr.Serve(l)
}

type logFormatter struct{}

func (l logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}
	buffer.Write([]byte("["))
	buffer.Write([]byte(entry.Time.Format("2006-01-02 15:04:05.000")))
	buffer.Write([]byte("] "))
	buffer.Write([]byte("["))
	buffer.Write([]byte(entry.Level.String()))
	buffer.Write([]byte("] "))
	if entry.HasCaller() {
		buffer.Write([]byte("["))
		buffer.Write([]byte(fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)))
		buffer.Write([]byte("] "))
	}
	buffer.Write([]byte(entry.Message))
	buffer.Write([]byte("\n"))
	return buffer.Bytes(), nil
}
