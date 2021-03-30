//Package modules implements the search and load function of clinics
package modules

import (
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

//Search find clinics with criteria
func Search(name, state string, availability string) []Clinic {
	mutex.RLock()
	defer mutex.RUnlock()

	list := make([]Clinic, 0)

	for _, v := range Data {

		if !contains(v.Name, name) {
			continue
		}

		if !(contains(v.StateCode, state) || contains(v.StateName, state)) {
			continue
		}

		if len(availability) > 0 {
			//it is closed
			if !(v.Opening.From <= availability && availability < v.Opening.To) {
				continue
			}
		}

		list = append(list, *v)

	}

	return list

}

var (
	searcher = search.New(language.English, search.IgnoreCase)
)

func contains(str, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	s, _ := searcher.IndexString(str, substr)

	return s > -1

}
