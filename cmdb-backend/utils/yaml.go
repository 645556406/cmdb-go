package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

//
//type ReadYamlConfig interface {
//	ReadYamlConfig(path string) (interface{}, error)
//}

type YamlConfig struct {
	HOSTNAME     string `yaml:"hostname"`
	PORT         int    `yaml:"port"`
	UserName     string `yaml:"username"`
	PASSWORD     string `yaml:"password"`
	DATABASE     string `yaml:"database"`
	SSLMode      bool   `yaml:"sslmode"`
	PoolSize     int    `yaml:"pool_size"`
	Timeout      int    `yaml:"timeout"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

// LoadYamlConfig ReadYamlConfig 从指定的yaml配置文件中读取数据并反序列化到指定的对象中
//
// 参数：
//
//	y *YamlConfig - YamlConfig类型的指针，用于表示读取配置的YamlConfig实例
//	path string - 配置文件路径
//	v interface{} - 需要反序列化的目标对象
//
// 返回值：
//
//	interface{} - 反序列化后的对象
//	error - 错误信息，如果读取和解析配置文件过程中发生错误，则返回非nil的error
func LoadYamlConfig(path string) (*YamlConfig, error) {
	var config YamlConfig
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	errYaml := yaml.Unmarshal(file, &config)
	if errYaml != nil {
		return nil, errYaml
	}
	return &config, nil
}

// LoadYamlConfigNew 从指定路径加载YAML配置文件，并返回一个map[string]interface{}类型的配置项
// path: YAML配置文件的路径
// 返回值: map[string]interface{}类型的配置项
func LoadYamlConfigNew(path string) map[string]interface{} {
	var config map[string]interface{}
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	errYaml := yaml.Unmarshal(file, &config)
	if errYaml != nil {
		panic(errYaml)
	}
	return config
}
