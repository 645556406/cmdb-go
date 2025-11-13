package utils

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

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

// expandEnvVars 替换配置中的环境变量
func expandEnvVars(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		// 支持 ${VAR_NAME} 格式的环境变量
		if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
			envKey := strings.TrimPrefix(v, "${")
			envKey = strings.TrimSuffix(envKey, "}")
			if envValue := os.Getenv(envKey); envValue != "" {
				return envValue
			}
		}
		return v
	case map[string]interface{}:
		for key, value := range v {
			v[key] = expandEnvVars(value)
		}
		return v
	case []interface{}:
		for i, item := range v {
			v[i] = expandEnvVars(item)
		}
		return v
	default:
		return v
	}
}

// LoadYamlConfigNew 从指定路径加载YAML配置文件，并返回一个map[string]interface{}类型的配置项
// 支持环境变量替换
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
	
	// 展开环境变量
	config = expandEnvVars(config).(map[string]interface{})
	
	return config
}
