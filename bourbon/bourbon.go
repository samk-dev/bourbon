package bourbon

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Bourbon struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	config   config
}

type config struct {
	port           string
	templateEngine string
}

func (b *Bourbon) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:     rootPath,
		foldersNames: []string{"handlers", "migrations", "views", "data", "public", "temp", "logs", "middlware"},
	}

	err := b.Init(pathConfig)
	if err != nil {
		return err
	}

	err = b.CheckDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	infoLog, errorLog := b.startLoggers()
	b.InfoLog = infoLog
	b.ErrorLog = errorLog
	b.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	b.Version = version
	b.RootPath = rootPath
	b.Routes = b.routes().(*chi.Mux)

	b.config = config{
		port:           os.Getenv("PORT"),
		templateEngine: os.Getenv("TEMPLATE_ENGINE"),
	}

	return nil
}

func (b *Bourbon) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.foldersNames {
		err := b.CreateDirIfDoesntExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bourbon) CheckDotEnv(path string) error {
	err := b.CreateFileIfDoesntExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

// ListenAndServe start the web server
func (b *Bourbon) ListenAndServe() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     b.ErrorLog,
		Handler:      b.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	b.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := server.ListenAndServe()
	b.ErrorLog.Fatal(err)
}
