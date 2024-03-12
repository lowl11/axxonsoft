package task_repo

import (
	"axxonsoft/data/entity"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/lowl11/boost/storage/sql"
)

func (repo Repo) FindByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	var task entity.Task
	if err := repo.
		all().
		Where(func(where sql.Where) {
			where.Equal("id", "$1")
		}).
		ScanSingle(ctx, &task, id); err != nil {
		if errors.Is(err, sql.RecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (repo Repo) Add(ctx context.Context, task entity.Task) error {
	return repo.add(task).Exec(ctx)
}

func (repo Repo) Update(ctx context.Context, task entity.Task) error {
	return repo.
		update(task).
		Where(func(where sql.Where) {
			where.Equal("id", "$1")
		}).
		Exec(ctx, task.ID)
}
