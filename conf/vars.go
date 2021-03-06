//Package conf provides settings of app
package conf

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/naoina/toml"
)

var (
	//Config configuration of app
	Config AppCfg
)

//init initiale config from ./conf.d/app.toml
func init() {

	fileName, err := filepath.Abs("./conf.d/app.toml")
	if err != nil {
		log.Println("cfg: ", err)
	}
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("cfg: ", err)
	}

	if err := toml.NewDecoder(strings.NewReader(string(buf))).Decode(&Config); err != nil {
		log.Println("cfg: ", err)
	}

}

//AppCfg configuration of app
type AppCfg struct {
	Addr   string            `toml:"addr"`
	URLs   []string          `toml:"urls"`
	States map[string]string `toml:"states"`
}
