package modules

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"clinics-apis/conf"
)

var (
	//Data clinics loaded
	Data  = make([]*Clinic, 0, 1000)
	mutex sync.RWMutex
)

//LoadClinics load clinics from remote url based on configuration
func LoadClinics() {
	mutex.Lock()
	defer mutex.Unlock()
	for _, u := range conf.Config.URLs {
		list, err := loadClinicsWithURL(u)
		if err != nil {
			log.Printf("load: %v %v\n", u, err)
			continue
		}

		Data = append(Data, list...)

	}

}

func loadClinicsWithURL(u string) ([]*Clinic, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	jsonBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ParseClinics(jsonBuf)
}

//ParseClinics parse clinics from json buffer
func ParseClinics(jsonBuf []byte) ([]*Clinic, error) {
	list := make([]*Clinic, 0)
	if err := json.Unmarshal(normalizeCloumnNames(jsonBuf), &list); err != nil {

		return list, err
	}

	for _, c := range list {

		fixState(c)
	}

	return list, nil
}

func fixState(c *Clinic) {
	//state code is missing
	if len(c.StateCode) == 0 {
		c.StateCode = getStateCode(c.StateName)
	}

	//state name is mssing
	if len(c.StateName) == 0 {
		c.StateName = getStateName(c.StateCode)
	}

}

func normalizeCloumnNames(jsonBuf []byte) []byte {
	normalizedName := strings.Replace(string(jsonBuf), "\"clinicName\"", "\"name\"", -1)
	normalizedAvailability := strings.Replace(normalizedName, "\"availability\"", "\"opening\"", -1)

	return []byte(normalizedAvailability)

}

func getStateName(code string) string {

	for k, v := range conf.Config.States {
		if strings.EqualFold(k, code) {
			return v
		}
	}

	return ""
}

func getStateCode(name string) string {
	for k, v := range conf.Config.States {
		if strings.EqualFold(v, name) {
			return k
		}
	}

	return ""
}

//Clinic info of clinic
type Clinic struct {
	Name      string  `json:"name"`
	StateCode string  `json:"stateCode"`
	StateName string  `json:"stateName"`
	Opening   Opening `json:"opening"`
}

//Opening opening hours of clinic
type Opening struct {
	From string `json:"from"`
	To   string `json:"to"`
}
