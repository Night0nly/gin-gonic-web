package infrastructure

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/controller"
)

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()

	//CORS setup
	corsConfig := cors.DefaultConfig()
	router.Use(cors.New(corsConfig))

	return &Router{
		router: router,
	}
}

func (s *Router) Run() {
	s.router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	api := s.router.Group("/api")

	// Song controller
	api.GET("/song", func(context *gin.Context) {
		songController := controller.NewSongController(context)
		songController.SearchSongByName()
	})

	s.router.Run(":8088")
}
