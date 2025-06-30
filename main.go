package main

import (
	"log"
	"net/http"

	"server/internal/config"
	"server/internal/db"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/repository"
	"server/internal/service"
	"server/internal/session"
)

func main() {
	// 1. Load configuration
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2. serve static files
	middleware.ServeStatic(cfg.StaticDir)

	// 3. initialize db pool
	dbConn, err := db.NewDB(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	sess := session.NewSession(cfg.Session.Key)

	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	// 5. wire repositories and services
	repo := repository.NewRepo(dbConn)

	userSvc := service.NewUserService(repo)
	productSvc := service.NewProductService(repo)

	svcs := service.Services{
		UserSvc:    userSvc,
		ProductSvc: productSvc,
	}

	// 6. Build HTTP Handler
	h := handler.NewHandler(svcs, sess)

	// 7. register routes
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/login", h.LoginHandler)
	http.HandleFunc("/register", h.RegisterHandler)
	http.HandleFunc("/logout", h.LogoutHandler)
	http.HandleFunc("/products", h.ProductsHandler)
	http.HandleFunc("GET /products/{id}", h.ProductDetailsHandler)
	http.HandleFunc("/cart/add", h.AddToCartHandler)
	http.HandleFunc("/cart/remove", h.RemoveFromCartHandler)
	http.HandleFunc("/cart", h.CartHandler)

	// 8. start server and listen
	srv := &http.Server{
		Addr: cfg.Server.HOST + ":" + cfg.Server.PORT,
	}
	log.Printf("listening on :%v", cfg.Server.PORT)
	log.Fatal(http.ListenAndServe(srv.Addr, nil))
}
