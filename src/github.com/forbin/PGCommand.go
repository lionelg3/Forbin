package forbin

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os/exec"
)

const (
	PG_SQL_SELECT_NOW                 = "select now()"
	PG_SQL_IS_IN_RECOVERY             = "select pg_is_in_recovery()"
	PG_SQL_CURRENT_XLOG_LOCATION      = "select pg_current_xlog_location()"
	PG_SQL_LAST_XLOG_RECEIVE_LOCATION = "select pg_last_xlog_receive_location()"
)

type PGCommand struct {
	serverCmd      *exec.Cmd
	db             *sql.DB
	DriverName     string
	DataSourceName string
}

func NewPGCommand(driverName, dataSourceName string) *PGCommand {
	_db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln("sql.Open err: %s", err)
		return nil
	}
	return &PGCommand{
		db: _db,
	}
}

func (self *PGCommand) HealCheck() (err error) {
	err = self.db.Ping()
	if err != nil {
		log.Fatalf("Ping Warn: %s", err)
		return err
	}
	_, err = self.db.Query(PG_SQL_SELECT_NOW)
	if err != nil {
		log.Fatalf("Healcheck query warn: %s", err)
	}
	return err
}

func (self *PGCommand) Promote() (err error) {
	return
}

func (self *PGCommand) ReParent() (err error) {
	return
}

func (self *PGCommand) LogStatus() (status string, err error) {
	rows, err := self.db.Query(PG_SQL_CURRENT_XLOG_LOCATION)
	//rows, err := self.db.Query(PG_SQL_SELECT_NOW)
	if err != nil {
		log.Fatalf("LogStatus query warn: %s", err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&status)
	}
	return status, err
}

func (self *PGCommand) Stats() (stats sql.DBStats) {
	return
}

func (self *PGCommand) InitDatabase() (err error) {
	return exec.Command("./pg-init-db.sh").Run()
}

func (self *PGCommand) ReConfigure(file string) (err error) {
	return
}

func (self *PGCommand) Startup() (err error) {
	self.serverCmd = exec.Command("./pg-start-db.sh")
	err = self.serverCmd.Start()
	return err
}

func (self *PGCommand) Graceful() error {
	return exec.Command("./pg-gracefull-db.sh").Run()
}

func (self *PGCommand) Shutdown() error {
	return exec.Command("./pg-stop-db.sh").Run()
}

func (self *PGCommand) Backup() ([]byte, error) {
	return exec.Command("./pg-backup-db.sh").CombinedOutput()
}

func (self *PGCommand) Restore() ([]byte, error) {
	return exec.Command("./pg-restore-db.sh").CombinedOutput()
}
