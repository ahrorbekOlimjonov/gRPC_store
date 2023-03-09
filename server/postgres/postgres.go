package postgres

import (
	pb "GRPC-TODO/proto"
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func CreateStore(store *pb.Store) (*pb.Store, error) {

	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "ahrorbek", "3108", "store")

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	newStore := pb.Store{}

	err = db.QueryRow(`INSERT INTO stores (name, description, is_open, addresses) 
	VALUES($1, $2, $3, $4) 
	RETURNING id, name, description, is_open, addresses`,
		store.Name,
		store.Description,
		store.IsOpen,
		pq.Array(store.Addresses)).Scan(
		&newStore.Id,
		&newStore.Name,
		&newStore.Description,
		&newStore.IsOpen,
		pq.Array(&newStore.Addresses))

	if err != nil {
		log.Fatalf("Failed to insert store from postgres: %v",err)
	}

	return &newStore, nil
}
