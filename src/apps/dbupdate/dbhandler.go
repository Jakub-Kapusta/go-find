// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/Jakub-Kapusta/go-find/apps/find"
)

type dbHandler struct {
	ctx          context.Context
	wg           sync.WaitGroup
	db           *sql.DB
	tx           *sql.Tx // Set to nil before creating actual transaction.
	c            chan *find.FileInfo
	rootDir      string
	isSearchPath bool
	searchPath   string
	skipDirs     []string
}

func newDbHandler(ctx context.Context, rootDir string, isSearchPath bool, searchPath string) (*dbHandler, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	return &dbHandler{
		ctx:          ctx,
		db:           db,
		tx:           nil,
		c:            make(chan *find.FileInfo, 32),
		rootDir:      rootDir,
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
		skipDirs: []string{
			"/dev",
			"/mnt",
			"/proc",
			"/run",
			"/sys",
			"/tmp",
		},
	}, nil
}

func (dbh *dbHandler) run() error {
	// TODO do something about this error.
	var rollback bool
	defer func() {
		if err := dbh.close(rollback); err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
		}
	}()

	_, err := dbh.db.Exec(prepDb)
	if err != nil {
		return err
	}

	if dbh.tx == nil {
		dbh.tx, err = dbh.db.Begin()
		if err != nil {
			rollback = true
			return err
		}
	} else {
		rollback = true
		return fmt.Errorf("something wrong with tx")
	}

	// Not using a single transaction and prepared statements will make things slow.
	stmt, err := dbh.tx.Prepare(preparedInsert)
	if err != nil {
		rollback = true
		return err
	}
	defer stmt.Close()

	for {
		select {
		case <-dbh.ctx.Done():
			// TODO add rollback
			rollback = true
			return dbh.ctx.Err()

		case fi, isOpen := <-dbh.c:
			if !isOpen {
				return nil
			}

			// Default to unknown type.
			type_id := 0
			if fi.D.IsDir() {
				// Directory.
				type_id = 1
			}

			_, err = stmt.Exec(fi.Path, type_id)
			if err != nil {
				rollback = true
				fmt.Println(err)
			}
		}

	}
}

func (dbh *dbHandler) getChan() chan<- *find.FileInfo {
	return dbh.c
}

// Do not call directly.
func (dbh *dbHandler) close(rollback bool) error {
	if dbh.tx != nil {
		if rollback {
			fmt.Println("ROLLBACK")
			if err := dbh.tx.Rollback(); err != nil {
				// TODO wait for the wg
				return err
			}
		} else {
			if err := dbh.tx.Commit(); err != nil {
				return err
			}
		}
	}

	dbh.wg.Wait()

	if err := dbh.db.Close(); err != nil {
		return err
	}

	return nil
}
