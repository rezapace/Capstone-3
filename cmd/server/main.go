package main

import (
	"Ticketing/internal/builder"
	"Ticketing/internal/config"
	"Ticketing/internal/http/binder"
	"Ticketing/internal/http/server"
	"Ticketing/internal/http/validator"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//menghubungkan ke postgresql atau database
	cfg, err := config.NewConfig(".env")
	checkError(err)

	splash()

	db, err := buildGormDB(cfg.Postgres)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db)

	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	srv := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)

	runServer(srv, cfg.Port)

	waitForShutdown(srv)
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

// berfungsi ketika API mati akan hidup sendiri lagi. ini untuk menghindari error ketika API mati
func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

// func untuk koneksi ke postgresql
func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// untuk membuat spalsh screen ini bisa menggunakan website
// ascii text style generator seperti patorjk.com
func splash() {
	colorReset := "\033[0m"

	splashText := `

	  ___________.__          __              __   .__                  
	  \__    ___/|__|  ____  |  | __  ____  _/  |_ |__|  ____     ____  
	    |    |   |  |_/ ___\ |  |/ /_/ __ \ \   __\|  | /    \   / ___\ 
	    |    |   |  |\  \___ |    < \  ___/  |  |  |  ||   |  \ / /_/  >
	    |____|   |__| \___  >|__|_ \ \___  > |__|  |__||___|  / \___  / 
	                      \/      \/     \/                 \/ /_____/      
`
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}

// func untuk cek error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 	// //memanggil entity user
// 	// users := make([]entity.User, 0)
// 	// if err := db.Find(&users).Error; err != nil {
// 	// 	checkError(err)
// 	// }
// 	// for _, v := range users { // ini untuk menampilkan data user secara looping
// 	// 	fmt.Println(v)
// 	// }
