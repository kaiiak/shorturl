package config

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

func init() {
	flag.String("dbtype", "", "database type")
	flag.String("dbpath", "", "open database path")
	flag.String("cachepath", "", "open cache path")
	flag.String("cachepwd", "", "cache password")
	flag.Int("port", 80, "listen port, default 80")
}

// Config 系统配置项
type Config struct {
	DBType    string `json:"dbtype"`
	DBPath    string `json:"dbpath"`
	CachePwd  string `json:"cachepwd"`
	CachePath string `json:"cachepath"`
	Port      int    `json:"port"`
}

// New 读取系统配置
// 如果path为空，则只从argv中读取
// 命令行中设置的属性值的优先级比配置文件中的优先级高
func New(path string) (*Config, error) {
	cnf := &Config{}
	if path != "" {
		fd, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer fd.Close()
		temp, err := ioutil.ReadAll(fd)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(temp, cnf); err != nil {
			return nil, err
		}
	}
	if err := cnf.parseFlag(); err != nil {
		return nil, err
	}
	return cnf, nil
}

// ParseFlag 处理cli
func (c *Config) parseFlag() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	vc := reflect.Indirect(reflect.ValueOf(c))
	ct := vc.Type()
	valueCache := map[string]reflect.Value{}
	for i := 0; i < ct.NumField(); i++ {
		sf := ct.Field(i)
		tag := sf.Tag.Get("json")
		valueCache[tag] = vc.Field(i)
	}
	var err error
	flag.Visit(func(f *flag.Flag) {
		field, ok := valueCache[f.Name]
		if !ok {
			return
		}
		if err != nil {
			return
		}
		err = setField(field, f.Value.String())
	})
	return err
}

func setField(filed reflect.Value, value string) error {
	if !filed.CanSet() {
		return errors.New("field can not set")
	}
	switch filed.Type().Kind() {
	case reflect.String:
		filed.SetString(value)
	case reflect.Bool:
		temp, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		filed.SetBool(temp)
	case reflect.Int, reflect.Int64:
		temp, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		filed.SetInt(temp)
	case reflect.Uint, reflect.Uint64:
		temp, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		filed.SetUint(temp)
	default:
		return errors.New("unsupport type")
	}
	return nil
}
