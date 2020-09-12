package conf

// Config 配置信息
type Config struct {
	DB DataBase
}

// DataBase ...
type DataBase struct {
	DriverName string
	Host       string
	Port       int
	User       string
	PassWord   string
	DBName     string
}
