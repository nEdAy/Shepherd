package config

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/rs/zerolog/log"
)

var (
	Config *ini.File

	App      app
	Server   server
	Path     path
	Database database
	Dataoke  dataoke
	WeChat   weChat
	Redis    redis
	Mob      mob
)

func Setup() {
	var err error
	Config, err = ini.Load("./build/app.ini")
	if err != nil {
		log.Fatal().Msgf("Fail to parse './build/app.ini': %v", err)
	}
	Config.BlockMode = false // if false, only reading, speed up read operations about 50-70% faster
	loadApp()
}

func loadApp() {
	sec, err := Config.GetSection("App")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'App': %v", err)
	}
	App.RunMode = sec.Key("RUN_MODE").In("debug", []string{"debug", "release"})
	if App.RunMode == "debug" {
		Config, err = ini.Load("./build/debug.ini")
		if err != nil {
			log.Fatal().Msgf("Fail to parse './build/debug.ini': %v", err)
		}
	} else if App.RunMode == "release" {
		Config, err = ini.Load("./build/release.ini")
		if err != nil {
			log.Fatal().Msgf("Fail to parse './build/release.ini': %v", err)
		}
	} else {
		log.Fatal().Msgf("Fail to parse Error RunMode")
	}
	Config.BlockMode = false // if false, only reading, speed up read operations about 50-70% faster
	App.HmacSecret = sec.Key("Hmac_Secret").MustString("nEdAy")
	App.HmacSecret = sec.Key("Password_Salt").MustString("nEdAy")

	loadServer()
	loadPath()
	loadDatabase()
	loadDataoke()
	loadWeChat()
	loadRedis()
	loadMob()
}

func loadServer() {
	sec, err := Config.GetSection("Server")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Server': %v", err)
	}
	Server.Protocol = sec.Key("PROTOCOL").In("http", []string{"http", "https"})
	Server.Host = sec.Key("HOST").MustString("127.0.0.1")
	Server.Port = sec.Key("PORT").MustInt(80)
	Server.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	Server.WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadPath() {
	sec, err := Config.GetSection("Paths")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Paths': %v", err)
	}
	Path.DataPath = sec.Key("DATA_PATH").MustString("./runtime")
	Path.LogPath = sec.Key("LOG_PATH").MustString("./runtime/log")
	Path.CachePath = sec.Key("CACHE_PATH").MustString("./runtime/cache")
	Path.CertFilePath = sec.Key("CERT_FILE_PATH").MustString("./data/ssl/www.neday.cn.pem")
	Path.KeyFilePath = sec.Key("KEY_FILE_PATH").MustString("./data/ssl/www.neday.cn.key")
}

func loadDatabase() {
	sec, err := Config.GetSection("Database")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Database': %v", err)
	}
	Database.Debug = sec.Key("DEBUG").MustBool(false)
	Database.Type = sec.Key("TYPE").MustString("mysql")
	Database.User = sec.Key("USER").MustString("root")
	Database.Password = sec.Key("PASSWORD").String()
	Database.Host = sec.Key("HOST").String()
	Database.Port = sec.Key("PORT").MustInt(3306)
	Database.Name = sec.Key("NAME").String()
	Database.MaxIdleConns = sec.Key("MAX_IDLE_CONNS").MustInt(64)
	Database.MaxOpenConns = sec.Key("MAX_OPEN_CONNS").MustInt(24)
	Database.PingInterval = time.Duration(sec.Key("PING_INTERVAL").MustInt(30)) * time.Second
}

func loadDataoke() {
	sec, err := Config.GetSection("Dataoke")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Dataoke': %v", err)
	}
	Dataoke.AppSecret = sec.Key("APP_SECRET").String()
	Dataoke.AppKey = sec.Key("APP_KEY").String()
}

func loadWeChat() {
	sec, err := Config.GetSection("WeChat")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'WeChat': %v", err)
	}
	WeChat.CodeToSessionUrl = sec.Key("CODE_TO_SESSION_URL").String()
	WeChat.AppID = sec.Key("APP_ID").String()
	WeChat.AppSecret = sec.Key("APP_SECRET").String()
	WeChat.SessionMagicID = sec.Key("SESSION_MAGIC_ID").String()
}

func loadRedis() {
	sec, err := Config.GetSection("Redis")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Redis': %v", err)
	}
	Redis.Host = sec.Key("Host").String()
	Redis.Password = sec.Key("Password").String()
	Redis.MaxIdle = sec.Key("MaxIdle").MustInt()
	Redis.MaxActive = sec.Key("MaxActive").MustInt()
	Redis.IdleTimeout = time.Duration(sec.Key("IdleTimeout").MustInt(30)) * time.Second
}

func loadMob() {
	sec, err := Config.GetSection("Mob")
	if err != nil {
		log.Fatal().Msgf("Fail to get section 'Mob': %v", err)
	}
	Mob.AppKey = sec.Key("APP_KEY").String()
}

type app struct {
	RunMode      string
	HmacSecret   string
	PasswordSalt string
}

type server struct {
	Protocol     string
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type path struct {
	DataPath     string
	LogPath      string
	CachePath    string
	CertFilePath string
	KeyFilePath  string
}

type database struct {
	Debug        bool
	Type         string
	User         string
	Password     string
	Host         string
	Port         int
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	PingInterval time.Duration
}

type dataoke struct {
	AppSecret string
	AppKey    string
}

type redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type weChat struct {
	CodeToSessionUrl string
	AppID            string
	AppSecret        string
	SessionMagicID   string
}

type mob struct {
	AppKey string
}
