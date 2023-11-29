package gormlib

import (
	"erp/internal/api/request"
	"erp/internal/infrastructure"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GormRepository struct {
	infrastructure.Database
}

func NewGormRepository(db infrastructure.Database) infrastructure.DatabaseTransaction {
	return GormRepository{db}
}

func (g GormRepository) WithTransaction(db infrastructure.Database, txFunc func(tx *infrastructure.Database) error) (err error) {
	err = db.RDBMS.Transaction(func(tx *gorm.DB) error {
		// copy db
		copyDb := db
		copyDb.RDBMS = tx

		err := txFunc(&copyDb)
		return err
	})
	return err
}

type GormQueryPaginationBuilder[E any] struct {
	tx *gorm.DB
}

func GormQueryPagination[E any](tx *gorm.DB, o request.PageOptions, data *[]*E) *GormQueryPaginationBuilder[E] {
	q := &GormQueryPaginationBuilder[E]{
		tx: tx,
	}
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	offset := (o.Page - 1) * o.Limit

	q.tx = q.tx.Debug().Offset(int(offset)).Limit(int(o.Limit)).Find(&data)

	fmt.Println(q.tx.Error)
	return q
}

func (q *GormQueryPaginationBuilder[E]) Count(total *int64) *GormQueryPaginationBuilder[E] {
	if total == nil {
		total = new(int64)
	}
	q.tx.Count(total)
	return q
}

func (q *GormQueryPaginationBuilder[E]) Error() error {
	return errors.WithStack(q.tx.Error)
}