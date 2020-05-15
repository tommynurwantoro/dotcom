package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"tommynurwantoro/dotcom/internal/domain"
	"tommynurwantoro/dotcom/internal/http/handler"
	"tommynurwantoro/dotcom/internal/http/router"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &domain.CustomContext{c, readConfig()}
			return next(cc)
		}
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = getRendererConfig()

	// Init handler
	h := handler.NewHandler(e)

	// Init router
	router.NewRouter(e, h)

	e.Static("/assets", "assets")

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

// PRIVATE
func readConfig() *domain.StaticConfig {
	conf := &domain.StaticConfig{}
	yamlFile, err := ioutil.ReadFile("static_config.yaml")
	if err != nil {
		log.Printf("Error read yaml file: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return conf
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	errorPage := fmt.Sprintf("views/errors/%d.html", code)
	if err := c.File(errorPage); err != nil {
		fmt.Println("Error to load file")
	}
}

func getRendererConfig() *echoview.ViewEngine {
	return echoview.New(goview.Config{
		Root:         "internal/views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: false,
	})
}
