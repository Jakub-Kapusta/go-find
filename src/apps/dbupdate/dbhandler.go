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
	c            chan *find.FileInfo
	rootDir      string
	isSearchPath bool
	searchPath   string
	paths        []string
}

func newDbHandler(ctx context.Context, rootDir string, isSearchPath bool, searchPath string) (*dbHandler, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}
	return &dbHandler{
		ctx:          ctx,
		db:           db,
		c:            make(chan *find.FileInfo, 32),
		rootDir:      rootDir,
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
	}, nil
}

func (dbh *dbHandler) run() error {
	// TODO do something about this error.
	defer func() {
		if err := dbh.close(); err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
		}
	}()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS foo (path TEXT NOT NULL PRIMARY KEY);
	DELETE from foo;
	`
	_, err := dbh.db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	for {
		select {
		case <-dbh.ctx.Done():
			return dbh.ctx.Err()
		case fi, ok := <-dbh.c:
			if !ok {
				return nil
			}
			dbh.paths = append(dbh.paths, fi.Path)

		}

	}
}

func (dbh *dbHandler) getChan() chan<- *find.FileInfo {
	return dbh.c
}

// Do not call directly.
func (dbh *dbHandler) close() error {
	fmt.Println("inserting")
	fmt.Println(len(dbh.paths))
	cnt := 0

	tx, err := dbh.db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	// Not using a sigle transaction will make things incredibly slow.
	stmt, err := tx.Prepare("INSERT INTO foo (path) VALUES (?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range dbh.paths {
		cnt += 1
		fmt.Println(cnt)
		_, err = stmt.Exec(p)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	dbh.wg.Wait()

	if err := dbh.db.Close(); err != nil {
		return err
	}

	return nil
}
