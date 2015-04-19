package common

import (
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "io/ioutil"
    "log"
)

type IMConfig struct {
    IMPort     int      `json:"im_port"`
    HttpPort   int      `json:"http_port"`
    MaxClients int      `json:"max_clients"`
    DBConfig   DBConfig `json:"db_config"`
}

type DBConfig struct {
    Host         string `json:"host"`
    Username     string `json:"username"`
    Password     string `json:"password"`
    Name         string `json:"name"`
    MaxIdleConns int    `json:"max_idle_conns"`
    MaxOpenConns int    `json:"max_open_conns"`
}

// 读取配置文件
func ReadConfig(path string) (*IMConfig, error) {
    config := new(IMConfig)
    err := config.Parse(path)
    if err != nil {
        return nil, err
    }
    return config, nil
}

// 解析配置文件
func (this *IMConfig) Parse(path string) error {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        return &ConfigurationError{err.Error()}
    }
    err = json.Unmarshal(file, &this)
    if err != nil {
        return &ConfigurationError{err.Error()}
    }
    return nil
}

// 连接数据库
func (this *DBConfig) Connect() (*sql.DB, error) {
    // 从配置文件中读取配置信息并初始化连接池(go中含有连接池处理机制)
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", this.Username, this.Password, this.Host, this.Name))
    db.SetMaxIdleConns(this.MaxIdleConns) // 最大空闲连接
    db.SetMaxOpenConns(this.MaxOpenConns) // 最大连接数

    if err != nil {
        return nil, &DatabaseError{err.Error()}
    }
    if err := db.Ping(); err != nil {
        return nil, &DatabaseError{err.Error()}
    }
    log.Printf("连接数据库成功 !")
    return db, nil
}
