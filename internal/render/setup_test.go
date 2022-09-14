package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rafaelmardones/go-tour/internal/config"
	"github.com/rafaelmardones/go-tour/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.User{})

	// change when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (mw myWriter) Header() http.Header {
	return http.Header{}
}

func (mw myWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (mw myWriter) WriteHeader(i int) {
}
