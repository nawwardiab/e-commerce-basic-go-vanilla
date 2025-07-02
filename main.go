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

	
	// 2. initialize db pool
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.USER, cfg.DB.PWD, cfg.DB.HOST, cfg.DB.PORT, cfg.DB.DBNAME)
	dbConn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	sess := session.NewSession(cfg.Session.Key)

	// 3. wire repositories and services
	userRepo := repository.NewUserRepo(dbConn)
	userSvc := service.NewUserService(userRepo)
	
	productRepo := repository.NewProductRepo(dbConn)
	productSvc := service.NewProductService(productRepo)


	// 4. Handlers
	uh := handler.NewAuthHandler(*userSvc, sess)
	ph := handler.NewProdHandler(*productSvc, sess)
	hh := handler.NewHomeHandler(sess)
	ch := handler.NewCartHandler(sess, *productSvc)

	// 5. serve static files
	static := middleware.Handler(middleware.ServeStatic(cfg.StaticDir))
	http.Handle("/staticFiles/", static)

	// 6. register routes
	http.HandleFunc("/", hh.HomeHandler) // catch all! 
	http.HandleFunc("GET /login", uh.GetLoginHandler)
	http.HandleFunc("POST /login", uh.PostLoginHandler)
	http.HandleFunc("POST /register", uh.PostRegisterHandler)
	http.HandleFunc("GET /register", uh.GetRegisterHandler)
	http.HandleFunc("GET /logout", uh.LogoutHandler)
	http.HandleFunc("GET /products", ph.ProductsHandler)
	http.HandleFunc("GET /products/{id}", ph.ProductDetailsHandler)
	http.HandleFunc("POST /cart/add", ch.AddToCartHandler)
	http.HandleFunc("POST /cart/remove", ch.RemoveFromCartHandler)
	http.HandleFunc("GET /cart", ch.CartHandler)

	// 7. start server and listen
	srv := &http.Server{
		Addr: cfg.Server.HOST + ":" + cfg.Server.PORT,
	}
	log.Printf("listening on :%v", cfg.Server.PORT)
	log.Fatal(http.ListenAndServe(srv.Addr, nil))
}
