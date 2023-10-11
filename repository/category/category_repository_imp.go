package categoryRepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/model/entity"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	sql := "insert into categories(name) value (?)"

	result, err := tx.ExecContext(ctx, sql, category.Name)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)

	category.Id = int(id)

	return category

}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	sql := "update categories set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helpers.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	sql := "delete from categories where id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Id)
	helpers.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	sql := "select id,name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helpers.PanicIfError(err)
	defer func() {
		rows.Close()
	}()

	category := entity.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category is Not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	sql := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, sql)
	helpers.PanicIfError(err)
	defer func() {
		rows.Close()
	}()

	var categories []entity.Category

	for rows.Next() {
		category := entity.Category{}

		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories

}
