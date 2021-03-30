package modules

import (
	"clinics-apis/mock"
	"os"
	"testing"
)

//TestMain setup mock data
func TestMain(m *testing.M) {

	mock.SetupAppConfig()

	Data, _ = ParseClinics(mock.MockClinicBuffer)

	code := m.Run()

	os.Exit(code)
}

//TestCriteriaName test to search clinics with criteria:name
func TestCriteriaName(t *testing.T) {

	list := Search("Good", "", "")

	if len(list) != 2 {
		t.Errorf("Count got: %v , want: 2", len(list))
	}
}

//TestCriteriaStateCode test to search clinics with criteria:state code
func TestCriteriaStateCode(t *testing.T) {

	list := Search("", "CA", "")

	if len(list) != 5 {
		t.Errorf("Count got: %v , want: 5", len(list))
	}
}

//TestCriteriaStateCode test to search clinics with criteria:state name
func TestCriteriaStateName(t *testing.T) {

	list := Search("", "California", "")

	if len(list) != 5 {
		t.Errorf("Count got: %v , want: 5", len(list))
	}
}

//TestCriteriaAvailability test to search clinics with criteria:state availability
func TestCriteriaAvailability(t *testing.T) {

	list := Search("", "", "10:30")

	if len(list) != 5 {
		t.Errorf("Count got: %v , want: 5", len(list))
	}

	list = Search("", "", "23:30")

	if len(list) != 2 {
		t.Errorf("Count got: %v , want: 2", len(list))
	}

}

//TestMultipleCriterias test to search clinics with mulitiple criteria
func TestMultipleCriterias(t *testing.T) {

	list := Search("Good Health Home", "FL", "")

	if len(list) != 1 {
		t.Errorf("Count got: %v , want: 1", len(list))
	}

	list = Search("", "CA", "23:30")

	if len(list) != 2 {
		t.Errorf("Count got: %v , want: 2", len(list))
	}

}
