package entity

import (
	"github.com/lowl11/boost/storage/sql"
)

type Task struct {
	sql.Entity `ef:"table:tasks,alias:task"`

	Status string `db:"status"`

	RequestMethod  string `db:"request_method"`
	RequestURL     string `db:"request_url"`
	RequestHeaders string `db:"request_headers"`
	RequestBody    []byte `db:"request_body"`

	ResponseStatus     *string `db:"response_status"`
	ResponseStatusCode *int    `db:"response_status_code"`
	ResponseHeaders    *string `db:"response_headers"`
	ResponseLength     *int    `db:"response_length"`
	ResponseBody       []byte  `db:"response_body"`

	ResponseErrorReason *string `db:"response_error_reason"`
}
