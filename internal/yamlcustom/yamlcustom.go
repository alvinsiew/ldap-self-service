package yamlcustom

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ConfigLDAP struct for yaml LDAP config
type ConfigLDAP struct {
	UserDN string `yaml:"userdn"`
	LDAP   string `yaml:"ldap"`
}

type ConfigSMTP struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	From     string `yaml:"from"`
}

// Config struct for mnc config
type Config struct {
	Ldap []ConfigLDAP `yaml:"ldap"`
	Smtp []ConfigSMTP `yaml:"smtp"`
}

// ParseYAML parse yaml config file
func ParseYAML() Config {
	filename, err := filepath.Abs("../conf/config.yml")

	if err != nil {
		panic(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
