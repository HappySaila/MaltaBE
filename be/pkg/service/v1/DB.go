package v1

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type DB struct {
	db *sql.DB
}

func (d *DB) AddData(ctx context.Context, data string) (bool, error) {
	// get SQL connection from pool
	c, err := d.connect(ctx)
	if err != nil {
		log.Printf("Error: v%", err.Error())
		return false, err
	}
	defer c.Close()

	// insert User into db
	_, err = c.ExecContext(ctx, "insert into malta_be(`StatsData`) values(?)", data)
	if err != nil {
		return false, status.Error(codes.Unknown, "failed to insert into malta_be-> "+err.Error())
	}

	return true, nil
}

// connect returns SQL database connection from the pool
func (d *DB) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := d.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
