package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
	"kasir.api/database"
	"kasir.api/handler"
	"kasir.api/repository"
	"kasir.api/routes"
	service "kasir.api/services"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	repo := repository.NewInventoryPostgres(db)
	svc := service.NewInventoryService(repo)
	h := handler.NewInventoryHandler(svc)

	routes.RegisterRoutes(h)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	port := config.Port
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
