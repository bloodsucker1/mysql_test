package main

import (
	"encoding/base64"
	"github.com/ngaut/log"
	"github.com/wumansgy/goEncrypt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var conf = ConfigYaml{}

type ConfigYaml struct {
	DbConfig struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"config"`
}

func loadConfig(configFile string) {
	conf := new(ConfigYaml)
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Errorf("read file error: %v", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Errorf("Unmarshal: %v", err)
		return
	}
}

func EncryptPassword(plaintext string) string {
	// 传入明文和自己定义的密钥，密钥为8字节
	cryptText, err := goEncrypt.DesCbcEncrypt([]byte(plaintext), []byte("asd12345"), []byte{}) //得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(cryptText)
}

func DecryptPassword(cryptText string) string {
	decodeString, err := base64.StdEncoding.DecodeString(cryptText)
	if err != nil {
		panic(err)
	}
	newPlaintext, err := goEncrypt.DesCbcDecrypt(decodeString, []byte("asd12345"), []byte{})
	if err != nil {
		panic(err)
	}
	return string(newPlaintext)
}
