package main

func initializeRoutes() {
	routes := router.Group("/v1")
	{
		routes.GET("/today", todayThumbs)
		routes.GET("/day/:day", somedayThumbs)
		routes.GET("/thumb/:id", getThumb)
		routes.GET("/video/:id", getVideo)
	}
}
