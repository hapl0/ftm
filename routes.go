// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/cases", ShowUserCases)
	router.GET("/case/:id", showCase)
	router.GET("/case/:id/wallets", showCaseWallets)
	router.GET("/case/:id/wallet/id:", showCaseWallet)
	router.GET("/wallet/:id", showWallet)
	router.GET("/bye", bye)
	router.POST("/auth/", auth)

}
