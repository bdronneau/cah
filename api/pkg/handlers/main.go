package handlers

import (
	"flag"
	"fmt"
	"net/http"

	"cah/pkg/models"
	"cah/pkg/services"
	"cah/pkg/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// App of package
type App interface {
	NewHTTP(string)
}

type app struct {
	httpCORS string
	httpPort int
}

// Config for flags
type Config struct {
	dbHost     *string
	dbUser     *string
	dbName     *string
	dbPassword *string
	dbPort     *int

	httpCORS *string
	httpPort *int
}

type tokenURL struct {
	Code  string `json:"code"`
	State string
	Scope string
}

type stats struct {
	History    []models.InformationsByTurn `json:"history"`
	Classement []models.RoomClassement     `json:"classement"`
}

// Flags defines cli args for Strava Client
func Flags(fs *flag.FlagSet) Config {
	return Config{
		httpCORS: fs.String("http-cors", "http://localhost:8080", "[http] CORS domain"),
		httpPort: fs.Int("http-port", 1324, "[http] api port"),

		dbHost:     fs.String("db-host", "localhost", "DB Host"),
		dbUser:     fs.String("db-user", "cah", "DB User"),
		dbName:     fs.String("db-name", "cah", "DB Name"),
		dbPassword: fs.String("db-password", "xoxo", "DB Password"),
		dbPort:     fs.Int("db-port", 5432, "DB Port"),
	}
}

// New function create an app
func New(config Config) (App, error) {
	services.InitDB(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", *config.dbUser, *config.dbPassword, *config.dbHost, *config.dbPort, *config.dbName))

	return &app{
		httpCORS: *config.httpCORS,
		httpPort: *config.httpPort,
	}, nil
}

func (a app) NewHTTP(env string) {
	e := echo.New()

	e.Use(utils.MiddlewareLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{a.httpCORS},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's okay, but not tested!")
	})

	g := e.Group("/api")

	g.GET("/users/:id", a.handleGetUser)
	g.POST("/users", a.handlePostUser)

	g.GET("/rooms", a.handleListRoom)
	g.POST("/rooms", a.handlePostRoom)

	g.GET("/rooms/:name", a.handleRoomStatus)
	g.GET("/rooms/:name/stats", a.handleRoomStats)
	g.GET("/rooms/:name/users", a.handleListUser)
	g.GET("/rooms/:name/users/:user_id/cards", a.handleListUserCard)
	g.HEAD("/rooms/:name", a.handleHeadRoomStatus)
	g.POST("/rooms/:name/next", a.handlePostTurn)
	g.POST("/rooms/:name/start", a.handleRoomStart)
	g.POST("/rooms/:name/stop", a.handleRoomStop)
	g.POST("/rooms/:name/users", a.handleRoomPostUser)
	g.POST("/rooms/:name/users/:user_id/card", a.handlePostUserCard)
	g.POST("/rooms/:name/users/:user_id/elected", a.handlePostVote)
	// TODO move in status
	g.GET("/rooms/:name/cards/current", a.handleRoomSCards)
	g.GET("/rooms/:name/judge", a.handleJudgeCards)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", a.httpPort)))
}
