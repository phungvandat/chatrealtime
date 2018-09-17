package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chatsex/config/database/pg"
	"github.com/chatsex/endpoints"
	serviceHttp "github.com/chatsex/http"
	"github.com/chatsex/service"
	messageSvc "github.com/chatsex/service/message"
	roomSvc "github.com/chatsex/service/room"
	userSvc "github.com/chatsex/service/user"
	"github.com/chatsex/websoc"
	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	// setup env on local
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by errors: %v", err))
		}
	}

	// setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	//setupService
	var (
		pgDB, closeDB = pg.New(os.Getenv("PG_DATASOURCE"))
		s             = service.Service{
			UserService: service.Compose(
				userSvc.NewPGService(pgDB),
			).(userSvc.Service),
			RoomService: service.Compose(
				roomSvc.NewPGService(pgDB),
			).(roomSvc.Service),
			MessageService: service.Compose(
				messageSvc.NewPGService(pgDB),
			).(messageSvc.Service),
		}
	)
	defer closeDB()

	websoc, err := websoc.NewWebSoc(&s)
	if err != nil {
		logger.Log("cannot create websoc ", err)
	}

	err = http.ListenAndServe(":3000", serviceHttp.NewHTTPHandler(endpoints.MakeServerEndpoints(s), websoc, logger))
	if err != nil {
		logger.Log("exit", err)
	}
}
