package config

import "fmt"

type MySql struct {
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" yaml:"port"`
	UserName  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Database  string `json:"database" yaml:"database"`
	Prefix    string `json:"prefix" yaml:"prefix"`
	Charset   string `json:"charset" yaml:"charset"`
	Collation string `json:"collation" yaml:"collation"`
}

func (m *MySql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		m.UserName, m.Password, m.Host, m.Port, m.Database, m.Charset)

}
