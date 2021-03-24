package sql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/sei-ri/microservice.io/pkg/pigeon"
)

const (
	DefaultDriver          = "mysql"
	DefaultURL             = "root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=True"
	DefaultMaxIdleConns    = 10
	DefaultMaxOpenConns    = 20
	DefaultConnMaxLifetime = time.Hour
)

type Storage struct {
	driver          string
	url             string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration

	db  *sql.DB
	Now func() time.Time
}

type Option func(s *Storage) error

func WithDriver(v string) Option {
	return func(s *Storage) error {
		s.driver = v
		return nil
	}
}

func WithURL(v string) Option {
	return func(s *Storage) error {
		s.url = v
		return nil
	}
}

func WithMaxIdleConns(v int) Option {
	return func(s *Storage) error {
		s.maxIdleConns = v
		return nil
	}
}

func WithMaxOpenConns(v int) Option {
	return func(s *Storage) error {
		s.maxOpenConns = v
		return nil
	}
}

func WithConnMaxLifetime(v string) Option {
	return func(s *Storage) error {
		if d, err := time.ParseDuration(v); err != nil {
			return err
		} else {
			s.connMaxLifetime = d
		}
		return nil
	}
}

func New(ctx context.Context, opts ...Option) (*Storage, error) {
	s := &Storage{
		driver:          DefaultDriver,
		url:             DefaultURL,
		maxIdleConns:    DefaultMaxIdleConns,
		maxOpenConns:    DefaultMaxOpenConns,
		connMaxLifetime: DefaultConnMaxLifetime,
		Now:             time.Now,
	}

	for i := range opts {
		if err := opts[i](s); err != nil {
			return nil, err
		}
	}

	return s, s.open(ctx)
}

func (s *Storage) open(ctx context.Context) error {
	db, err := sql.Open(s.driver, s.url)
	if err != nil {
		return err
	}
	if err := db.PingContext(ctx); err != nil {
		return err
	}

	db.SetMaxIdleConns(s.maxIdleConns)
	db.SetMaxOpenConns(s.maxOpenConns)
	db.SetConnMaxLifetime(s.connMaxLifetime)

	// TODO: automigrate with sql-migrate pkg

	s.db = db
	log.Printf("[PIGEON] %s connected at %s", s.driver, s.url)
	return nil
}

func (s *Storage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *Storage) Put(ctx context.Context, data ...*pigeon.Data) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	now := time.Now()
	args, values := []string{}, []interface{}{}
	for i := range data {
		args = append(args, "(?, ?, ?, ?, ?)")
		values = append(values, data[i].ID, data[i].Version, data[i].Type, data[i].Dump, now)
	}

	result, err := tx.ExecContext(ctx, `INSERT INTO event_log(id, version, type, dump, created_at) VALUES `+strings.Join(args, ","), values...)
	if err != nil {
		return err
	}

	if _, err := result.RowsAffected(); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Storage) All(ctx context.Context, f pigeon.Filter) ([]pigeon.Data, error) {
	where, args := filter(ctx, f)
	stmt := fmt.Sprintf("SELECT id, version, type, dump FROM event_log WHERE %s", where)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]pigeon.Data, 0)
	for rows.Next() {
		var row pigeon.Data
		if rows.Scan(
			&row.ID,
			&row.Version,
			&row.Type,
			&row.Dump,
		); err != nil {
			return nil, err
		}

		data = append(data, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) Get(ctx context.Context, f pigeon.Filter) (*pigeon.Data, error) {
	where, args := filter(ctx, f)
	stmt := fmt.Sprintf("SELECT id, version, type, dump FROM event_log WHERE %s ORDER BY version DESC LIMIT 1", where)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var row pigeon.Data
	if err := tx.QueryRowContext(ctx, stmt, args...).Scan(
		&row.ID,
		&row.Version,
		&row.Type,
		&row.Dump,
	); err != nil {
		return nil, err
	}

	return &row, nil
}

func filter(ctx context.Context, f pigeon.Filter) (string, []interface{}) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := f.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := f.Version; v != nil {
		where, args = append(where, "version = ?"), append(args, *v)
	}

	return strings.Join(where, " AND "), args
}
