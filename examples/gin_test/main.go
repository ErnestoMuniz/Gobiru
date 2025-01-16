package main

import (
	"log"
)

func main() {
	router := SetupRouter()

	// Servir arquivos estáticos do diretório docs
	router.Static("/docs", "./docs")

	log.Printf("Servidor rodando em http://localhost:8080")
	log.Printf("Documentação disponível em:")
	log.Printf("- Swagger UI: http://localhost:8080/docs/index.html")
	log.Printf("- OpenAPI JSON: http://localhost:8080/docs/openapi.json")
	log.Printf("- Routes JSON: http://localhost:8080/docs/routes.json")

	log.Fatal(router.Run(":8080"))
}
