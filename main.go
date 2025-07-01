package main

import (
	"fmt"
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
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2. serve static files
	static := middleware.ServeStatic(cfg.StaticDir)
	

	// 3. initialize db pool
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.USER, cfg.DB.PWD, cfg.DB.HOST, cfg.DB.PORT, cfg.DB.DBNAME)
	dbConn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	sess := session.NewSession(cfg.Session.Key)

	// 5. wire repositories and services
	userRepo := repository.NewUserRepo(dbConn)
	userSvc := service.NewUserService(userRepo)
	
	productRepo := repository.NewProductRepo(dbConn)
	productSvc := service.NewProductService(productRepo)


	// Handlers
	uh := handler.NewAuthHandler(userSvc, sess)
	ph := handler.NewProdHandler(productSvc, sess)
	hh := handler.NewHomeHandler(sess)
	ch := handler.NewCartHandler(sess, productSvc)

	http.Handle("/static/", static)

	// 7. register routes
	http.HandleFunc("/", hh.HomeHandler)
	http.HandleFunc("/login", uh.LoginHandler)
	http.HandleFunc("/register", uh.RegisterHandler)
	http.HandleFunc("/logout", uh.LogoutHandler)
	http.HandleFunc("/products", ph.ProductsHandler)
	http.HandleFunc("GET /products/{id}", ph.ProductDetailsHandler)
	http.HandleFunc("/cart/add", ch.AddToCartHandler)
	http.HandleFunc("/cart/remove", ch.RemoveFromCartHandler)
	http.HandleFunc("/cart", ch.HandleCart)

	// 8. start server and listen
	srv := &http.Server{
		Addr: cfg.Server.HOST + ":" + cfg.Server.PORT,
	}
	log.Printf("listening on :%v", cfg.Server.PORT)
	log.Fatal(http.ListenAndServe(srv.Addr, nil))
}
