package db

import (
	"fmt"
	"sync"
	"time"

	gomocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type Connection interface {
	Connect() (*gorm.DB, error)
	DB() *gorm.DB
	MockConnect() *gorm.DB
	Connected() bool
	DatabaseName() string
	Disconnect() bool
}

type MariaDBConnection struct {
	db        *gorm.DB
	mutex     *sync.Mutex
	config    Config
	connected bool
}

func New(config Config) Connection {
	conn := &MariaDBConnection{}
	conn.connected = false
	conn.mutex = &sync.Mutex{}
	conn.config = config
	return conn
}

func (conn *MariaDBConnection) Connect() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	//allow 3 attempts on connection before returning error
	for attempt := 0; attempt < 3; attempt++ {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true&maxAllowedPacket=0&query_cache_type=0",
			conn.config.User, conn.config.Password, conn.config.Host, conn.config.Port, conn.config.Database))
		if err != nil {
			continue
		}
		db.DB().SetConnMaxLifetime(time.Minute * 5)
		db.DB().SetMaxIdleConns(100)
		db.DB().SetMaxOpenConns(100)
		return db, nil
	}
	return nil, fmt.Errorf("could not connect to db %s", err)
}

func (conn *MariaDBConnection) Connected() bool {
	return conn.connected
}

func (conn *MariaDBConnection) MockConnect() *gorm.DB {
	gomocket.Catcher.Register()
	conn.connected = true
	conn.db, _ = gorm.Open(gomocket.DriverName, "")
	return conn.db
}

func (conn *MariaDBConnection) DB() *gorm.DB {
	if !conn.connected {
		conn.mutex.Lock()
		defer conn.mutex.Unlock()
		if !conn.connected {
			var err error
			conn.db, err = conn.Connect()
			if err != nil {
				logrus.Fatalf("message: %s config: %#v", err, conn.config)
			}
			conn.connected = true
		}
	}
	return conn.db
}

func (conn *MariaDBConnection) DatabaseName() string {
	return conn.config.Database
}

func (conn *MariaDBConnection) Disconnect() bool {
	if conn.Connected() {
		conn.db.Close()
		conn.connected = false
		return true
	}
	return false
}
