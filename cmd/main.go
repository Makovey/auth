package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	m "github.com/Makovey/microservice_auth/pkg/user/v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

const (
	dbDSN    = "host=localhost port=5432 dbname=postgres user=admin password=admin sslmode=disable"
	grpcPort = 3000
)

type server struct {
	m.UnimplementedUserV1Server
}

func (s *server) Create(_ context.Context, _ *m.User) (*m.CreateResponse, error) {
	return &m.CreateResponse{Id: 1000}, nil
}

func main() {
	//addAndReadSqlData()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	m.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func addAndReadSqlData() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	pool.Ping(ctx)
	defer pool.Close()

	ib := sq.Insert("auth").
		PlaceholderFormat(sq.Dollar).
		Columns("title").
		Values("Test").
		Suffix("RETURNING id")

	query, args, err := ib.ToSql()
	if err != nil {
		log.Fatalf("failed to generate query: %v", err)
	}

	var id int
	err = pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	fmt.Printf("Added new id -> %d", id)
	fmt.Println()

	sb := sq.Select("id", "title").
		PlaceholderFormat(sq.Dollar).
		From("auth")

	query, args, err = sb.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to select notes: %v", err)
	}

	id = 0
	var title string

	for rows.Next() {
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		log.Printf("id -> %d, title -> %s", id, title)
	}
}
