package repository

import (
	"SigmatechTechnicalTest-Golang/api/repository/r_order"
	"database/sql"
)

type Repository struct {
	OrderRepo r_order.OrderRepoInterface
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		OrderRepo: r_order.NewOrderRepo(db),
	}
}
