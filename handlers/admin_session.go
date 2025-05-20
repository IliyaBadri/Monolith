package handlers

type AdminCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// func AdminNewSessionHandler(responseWriter http.ResponseWriter, request *http.Request) {
// 	responseWriter.Header().Set("Content-Type", "application/json")
// 	if request.Method != http.MethodPost {
// 		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(responseWriter).Encode(map[string]string{
// 			"error": "Method Not Allowed",
// 		})
// 		return
// 	}

// 	request.Body = http.MaxBytesReader(responseWriter, request.Body, config.MaxBodyBytes)
// 	defer request.Body.Close()

// 	var credentials AdminCredentials
// 	err := json.NewDecoder(request.Body).Decode(&credentials)
// 	if err != nil {
// 		responseWriter.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(responseWriter).Encode(map[string]string{
// 			"error": "Invalid JSON",
// 		})
// 		return
// 	}

// 	adminExists := database.AdminExists(credentials.Username)

// 	if !adminExists {
// 		responseWriter.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(responseWriter).Encode(map[string]string{
// 			"error": "Unauthorized",
// 		})
// 		log.WithFields(log.Fields{
// 			"entered_username": credentials.Username,
// 			"ip_address":       request.RemoteAddr,
// 			"method":           request.Method,
// 			"requested_url":    request.RequestURI,
// 			"user_agent":       request.UserAgent(),
// 			"referrer":         request.Referer(),
// 		}).Warn("Failed admin login with an incorrect username for gaining a session")
// 		return
// 	}

// 	passwordHash := database.GetAdminPasswordHash(credentials.Username)

// 	if passwordHash == "" {
// 		responseWriter.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(responseWriter).Encode(map[string]string{
// 			"error": "Internal Error",
// 		})
// 		return
// 	}

// 	passwordValid := cryptography.VerifyHash(passwordHash, credentials.Password)

// 	if !passwordValid {
// 		responseWriter.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(responseWriter).Encode(map[string]string{
// 			"error": "Unauthorized",
// 		})
// 		log.WithFields(log.Fields{
// 			"username":      credentials.Username,
// 			"ip_address":    request.RemoteAddr,
// 			"method":        request.Method,
// 			"requested_url": request.RequestURI,
// 			"user_agent":    request.UserAgent(),
// 			"referrer":      request.Referer(),
// 		}).Warn("Failed admin login with a correct username but an incorrect password for gaining a session")
// 		return
// 	}

// 	sessionToken := database.AddAdminSession(credentials.Username)

// 	if sessionToken == "" {
// 		responseWriter.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(responseWriter).Encode(map[string]string{
// 			"error": "Internal Error",
// 		})
// 		return
// 	}

// 	responseWriter.WriteHeader(http.StatusOK)
// 	json.NewEncoder(responseWriter).Encode(map[string]string{
// 		"sessionToken": sessionToken,
// 	})

// 	log.WithFields(log.Fields{
// 		"username":   credentials.Username,
// 		"ip_address": request.RemoteAddr,
// 	}).Info("An admin has gained a new session token")
// }
