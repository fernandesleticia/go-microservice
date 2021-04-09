func main() {
	flag.Parse()
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, Heartbeat())
	})

	if err := e.Start(fmt.Sprintf("%s:%s", *host, *port)); err != nil {
		fmt.Println("Failed to start server!", err)
		os.Exit(1)
	}

	return
}