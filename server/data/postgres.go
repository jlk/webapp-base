package data // import "github.com/jlk/webapp-base/server/data"

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jlk/webapp-base/server/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	// PG is a global session for psql (once initiated by dbConnect)
	PG *sqlx.DB
)

func init() {
	var err error

	dbHost := Config.GetString("dbHost")
	dbUsername := Config.GetString("dbUsername")
	dbPassword := Config.GetString("dbPassword")
	dbDatabaseName := Config.GetString("dbDatabaseName")
	dbSSLMode := Config.GetString("dbSSLMode")
	var maskedDbPassword = ""

	if len(dbHost) == 0 || len(dbUsername) == 0 || len(dbPassword) == 0 || len(dbDatabaseName) == 0 {
		logrus.Error("Missing config value(s) for database. Values I see are printed below:")
		logrus.Errorf("  dbHost: %s\n", dbHost)
		logrus.Errorf("  dbUser: %s\n", dbUsername)
		if len(dbPassword) == 0 {
			maskedDbPassword = "not set"
		} else {
			maskedDbPassword = util.MaskPassword(dbPassword)
		}
		logrus.Errorf("  dbPassword: %s\n", maskedDbPassword)
		logrus.Errorf("  dbDatabaseName: %s\n", dbDatabaseName)
		logrus.Fatal("Cannot connect to database. Exiting")
	}

	pgsqlConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUsername, dbPassword, dbDatabaseName, dbSSLMode)
	maskedPgsqlConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUsername, maskedDbPassword, dbDatabaseName, dbSSLMode)
	logrus.Info("Connecting to database at " + maskedPgsqlConnectionString)
	// TODO: Support private TLS certs
	PG, err = sqlx.Connect("pgx", pgsqlConnectionString)
	if err != nil {
		logrus.Fatal("error connecting to database: " + err.Error())
	}
	logrus.Info("Connected to database.")
}
