package yamlcustom

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ConfigLDAP struct for yaml LDAP config
type ConfigLDAP struct {
	UserDN string `yaml:"userdn"`
}

// Config struct for mnc config
type Config struct {
	Conf []ConfigLDAP `yaml:"conf"`
}

// ParseYAML parse yaml config file
func ParseYAML() Config {
	// usr, err := user.Current()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	filename, _ := filepath.Abs("../config.yml")
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
