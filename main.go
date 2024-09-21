package main

import (
	"fmt"
	"github.com/askaroe/hexago/internal/adapter/database/gorm"
	"github.com/askaroe/hexago/internal/adapter/database/repo"
	"github.com/askaroe/hexago/internal/domain/service"
	"github.com/askaroe/hexago/internal/http"
	"github.com/askaroe/jsonlog"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	err := godotenv.Load()
	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}

	logger.PrintInfo("initialized env variables", nil)

	idb, err := gorm.NewDB()

	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}
	logger.PrintInfo("initialized database", nil)

	repos := repo.Init(idb)

	srv := service.Init(repos)

	gin, err := http.NewRouter(logger, srv)

	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}

	logger.PrintInfo("initialized router", nil)

	server, err := http.New(gin)

	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}

	logger.PrintInfo("intialized server", nil)

	startServerError := server.Start()

	fmt.Println("SERVER STARTED")

	var stopReason string
	select {
	case err = <-startServerError:
		stopReason = fmt.Sprintf("start server error: %v", err)
	case qs := <-quit:
		stopReason = fmt.Sprintf("reveived signal %s", qs.String())
	}

	logger.PrintInfo(fmt.Sprintf("%s\nshutting down server...\n", stopReason), nil)

	err = server.Stop()

	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}

	err = idb.Close()
	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}

	logger.PrintInfo("server stopped successfully", nil)
}
