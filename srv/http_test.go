package srv

import (
	"clinics-apis/mock"
	"clinics-apis/modules"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

//TestMain setup mock data
func TestMain(m *testing.M) {
	mock.SetupAppConfig()

	modules.Data, _ = modules.ParseClinics(mock.MockClinicBuffer)

	code := m.Run()

	os.Exit(code)
}

//TestSearchAPI test the http response of online search api
func TestSearchAPI(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()

	w := httptest.NewRecorder()

	wanted := "National Veterinary Clinic"

	req, _ := http.NewRequest("GET", "/search?name="+url.QueryEscape(wanted), nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status got: %v , want: 200", w.Code)
	}

	var list []modules.Clinic

	json.Unmarshal([]byte(w.Body.String()), &list)

	if len(list) != 1 {
		t.Errorf("Count got: %v , want: 1", len(list))
		return
	}

	if list[0].Name != wanted {
		t.Errorf("Clinic got: %v , want: %v", list[0].Name, wanted)
	}

}
