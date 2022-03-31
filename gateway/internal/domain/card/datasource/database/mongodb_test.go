package database

import (
	"testing"
)

//go:generate mockgen -package mock -source postgres.go -destination=../../mock/mock_postgres.go

func TestMain(m *testing.M) {

}

func TestCardRepository_Create(t *testing.T) {
	migrate.Start()

	migrate.Down()
}

func TestRepository_Find(t *testing.T) {
	migrate.Start()

	migrate.Down()
}

func TestRepository_List(t *testing.T) {
	migrate.Start()

	migrate.Down()
}

func TestRepository_Update(t *testing.T) {
	migrate.Start()

	migrate.Down()
}

func TestRepository_Delete(t *testing.T) {
	migrate.Start()

	migrate.Down()
}

func TestRepository_Search(t *testing.T) {
	migrate.Start()

	migrate.Down()
}
