package main

type Store interface {
  // users
  CreateUser() erro
}

type Storage struct{
  db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
  return &Storage{
    db: db,
  }
}

func (s *Storage) CreateUser() error {
  return nil
}
