package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    firebase "firebase.google.com/go/v4"
    "google.golang.org/api/option"
    
    "github.com/unb-tp2-daniel/trabalhofinal-angrybirdspace24-2/backend/routes"
)

func main() {
    fmt.Println("Iniciando backend")

    // inicializar firebase admin SDK
    opt := option.WithCredentialsFile("serviceAccountKey.json") // função depreceada, mudar dps
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("Erro ao inicializar Firebase: %v", err)
    }

    // configurar as rotas
    routes.StartRoutes(app)

    port := ":8080"
    fmt.Printf("Servidor rodando em http://localhost%s\n", port)
    
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}