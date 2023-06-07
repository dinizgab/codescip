package main

import (
	"codescip/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api/generate", controllers.GenerateDocumentation)

	// TODO - Utilizar de um framework de API pra criar a API que conecta o front com o GPT e ter as outras rotas
	// TODO - Colocar essa API_KEY dentro de um .env

	e.Logger.Fatal(e.Start(":8000"))
}
