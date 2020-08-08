package controllers

import (
	"iris/repositories"
	"iris/services"

	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	movieResult := movieService.ShowMovieName()

	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}
