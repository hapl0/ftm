// routes.go
// PUT : 200 entité, 201 creation entité + uri , 204 vide
// POST : 200 entité, 201 creation entité + uri , 204 vide
package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/cases", ShowUserCases)
	router.GET("/case/:id", showCase)
	router.PUT("/case/:id", putCase)
	router.POST("/case/", createCase)
	router.GET("/case/:id/wallets", showCaseWallets)
	router.GET("/case/:id/wallet/:wallet", showCaseWallet)
	router.PUT("/case/:id/wallet/:wallet", putCaseWallet)
	router.GET("/wallet/:id", showWallet)
	router.PUT("/wallet/:id", putWallet)
	router.GET("/bye", bye)
	router.GET("/admin", showAdminPanel)
	router.GET("/users", showUsersPanel)
	router.GET("/user/:id", showUserID)
	router.PUT("/user/:id", putUser)
	router.POST("/user/", createUser)
	router.POST("/auth/", auth)

}
