package handlers

import (
	"context"
	"github.com/CloudyKit/jet/v6"
	"github.com/FernandoJVideira/velox"
	"github.com/FernandoJVideira/velox/mailer"
	"github.com/FernandoJVideira/velox/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var vel velox.Velox
var testSession *scs.SessionManager
var testHandlers Handlers

func TestMain(m *testing.M) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	testSession = scs.New()
	testSession.Lifetime = 24 * time.Hour
	testSession.Cookie.Persist = true
	testSession.Cookie.SameSite = http.SameSiteLaxMode
	testSession.Cookie.Secure = false

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader("../views"),
		jet.InDevelopmentMode(),
	)

	testRenderer := render.Render{
		Renderer: "jet",
		RootPath: "../",
		Port:     "4000",
		JetViews: views,
		Session:  testSession,
	}

	vel := velox.Velox{
		AppName:       "myapp",
		Debug:         true,
		Version:       "1.0.1",
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		RootPath:      "../",
		Routes:        nil,
		Render:        &testRenderer,
		Session:       testSession,
		DB:            velox.Database{},
		JetViews:      views,
		EncryptionKey: vel.RandomString(32),
		Cache:         nil,
		Scheduler:     nil,
		Mail:          mailer.Mail{},
		Server:        velox.Server{},
	}

	testHandlers.App = &vel

	os.Exit(m.Run())
}

func getRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(vel.SessionLoad)

	mux.Get("/", testHandlers.Home)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./../public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))
	return mux
}

func getCtx(req *http.Request) context.Context {
	ctx, err := testSession.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
