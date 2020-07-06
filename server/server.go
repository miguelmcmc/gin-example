package server

import (
	"time"

	"github.com/pkg/errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//httpServer
type httpServer struct {
	address string
	port    string
	router  *gin.Engine
}

//HTTPServer functionalities
func HTTPServer(address string, port string) serverINT {
	return &httpServer{
		address: address,
		port:    port,
	}
}

type serverINT interface {
	Init() (err error)
}

var (
	//ErrRunningServer http
	ErrRunningServer = errors.New("Error Running HTTP server").Error()
)

//Init server listen on address and port
func (server *httpServer) Init() (err error) {
	//SetUp server
	server.setUp()
	//Init server
	err = server.serve()
	return
}

func (server *httpServer) serve() error {
	err := server.router.Run(server.address + ":" + server.port) // listen and serve on 0.0.0.0:8080
	if err != nil {
		return errors.Wrap(err, ErrRunningServer)
	}
	return nil
}

func (server *httpServer) setUp() {
	server.router = gin.Default()

	//Cors config
	server.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Id", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Id", "content-type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Router
	server.initRoutes()
}

func (server *httpServer) initRoutes() {
	rgAny := server.router.Group("")
	any(rgAny)

	rgTest := server.router.Group("/test")
	test(rgTest)
}
