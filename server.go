package main

import (
	"echo-boilerplate/chat"
	"echo-boilerplate/facebook"
	"echo-boilerplate/conf"
	"echo-boilerplate/controllers/api"
	"echo-boilerplate/database/orm"
	"echo-boilerplate/migrate"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Template 템플릿
type Template struct {
	templates *template.Template
}

// Render 렌더
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func route(e *echo.Echo) *echo.Echo {
	apiGroup := e.Group("/api")
	api.Router(apiGroup)
	chatGroup := e.Group("/ws")
	chat.Router(chatGroup)
	facebookGroup := e.Group("/facebook")
	facebook.Router(facebookGroup)
	return e
}

// Init web server
func Init(e *echo.Echo) *echo.Echo {
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e = route(e)

	e.Static("/", "public")
	return e
}

func main() {
	if err := conf.InitConfig("conf/config.toml"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	db := orm.Init()
	migrate.Exec(db)
	defer db.Close()

	e := echo.New()

	e = Init(e)
	if conf.Conf.LogLevel == "DEBUG" {
		fmt.Println("DEBUG MODE")
		e.Debug = true
	}
	e.HideBanner = false

	var addr = ""
	if port := os.Getenv("PORT"); port != "" {
		addr = ":"+port
	} else {
		addr = conf.Conf.Server.Addr
	}

	if err := e.Start(addr); err != nil {
		fmt.Fprintf(os.Stderr, "Server Start Error: %v\n", err)
		os.Exit(1)
	}
}
