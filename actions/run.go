package actions

// func RunAction() {
// 	logrus.Info("initializing database")

// 	database.InitializeDatabase()
// 	database.InitializeTables()

// 	if !database.AdminExists(config.MainAdminUsername) {
// 		database.AddAdmin(config.MainAdminUsername, config.MainAdminDefaultPassword)
// 	}

// 	http.HandleFunc("/api/admin/get-session/", handlers.AdminNewSessionHandler)

// 	fileServer := http.FileServer(http.Dir("./static"))
// 	safeHandler := http.StripPrefix("/", limitRequestBodyMiddleware(fileServer))
// 	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
// 		if request.Method != http.MethodGet && request.Method != http.MethodHead {
// 			http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
// 			return
// 		}
// 		safeHandler.ServeHTTP(responseWriter, request)
// 	})

// 	fullAddress := fmt.Sprintf("%s:%d", config.ServerAddress, config.ServerPort)

// 	server := &http.Server{
// 		Addr:         fullAddress,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 	}

// 	logrus.WithFields(logrus.Fields{
// 		"ip_address": config.ServerAddress,
// 		"port":       config.ServerPort,
// 	}).Infof("Static server is listening on http://%s", fullAddress)

// 	// Graceful shutdown (for future)
// 	go func() {
// 		err := server.ListenAndServe()
// 		if err != nil && err != http.ErrServerClosed {
// 			logrus.WithError(err).Fatal("HTTP server failed")
// 		}
// 	}()

// 	select {}
// }

// func limitRequestBodyMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
// 		request.Body = http.MaxBytesReader(responseWriter, request.Body, config.MaxBodyBytes)
// 		next.ServeHTTP(responseWriter, request)
// 	})
// }
