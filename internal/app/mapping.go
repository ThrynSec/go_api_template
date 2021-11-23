package app

func mapUrls() {
	router.GET("/ping", control.Ping())
}
