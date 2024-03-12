package task_repo

import "github.com/lowl11/boost/storage/sql"

const (
	table = "tasks"
	alias = "task"
)

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (repo Repo) all() sql.SelectBuilder {
	return sql.
		Select().
		From(table).
		As(alias)
}

func (repo Repo) add(entity any) sql.InsertBuilder {
	return sql.
		Insert().
		Entity(entity)
}

func (repo Repo) update(entity any) sql.UpdateBuilder {
	return sql.
		Update().
		Entity(entity)
}

func (repo Repo) delete() sql.DeleteBuilder {
	return sql.
		Delete(table)
}
