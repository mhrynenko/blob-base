package postgres

import (
	"blob-base/internal/data"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const tableName = "blobs"

type Blobs struct {
	database *pgdb.DB
	sql      squirrel.SelectBuilder
}

var selectedTable = squirrel.Select("n.*").From(fmt.Sprintf("%s as n", tableName))

func NewBlobsTable(blobs *pgdb.DB) data.Blobs {
	return &Blobs{
		database: blobs.Clone(),
		sql:      selectedTable,
	}
}

func (q *Blobs) New() data.Blobs {
	return NewBlobsTable(q.database)
}

func (q *Blobs) NewBlob(blob data.Blob) (data.Blob, error) {
	query := squirrel.Insert(tableName).Columns("attributes", "owner").
		Values(blob.Attributes, blob.Owner).Suffix("returning *")

	var result data.Blob
	err := q.database.Get(&result, query)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *Blobs) GetBlob() (*data.Blob, error) {
	var result data.Blob
	err := q.database.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, errors.New("No blob with such ID")
	}

	return &result, err
}

func (q *Blobs) DeleteBlob(blobId int64) error {
	query := squirrel.Delete(tableName).Where("id = ?", blobId)

	result, err := q.database.ExecWithResult(query)
	if err != nil {
		return err
	}

	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return errors.New("No blob with such ID")
	}

	return nil
}

func (q *Blobs) GetBlobs() ([]data.Blob, error) {
	var result []data.Blob

	err := q.database.Select(&result, q.sql)
	if err != nil {
		return nil, err
	}

	if len(result) < 1 {
		return nil, errors.New("no blob was found")
	}

	return result, nil
}

func (q *Blobs) Page(pageParams pgdb.OffsetPageParams) data.Blobs {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *Blobs) FilterByOwner(owners ...string) data.Blobs {
	q.sql = q.sql.Where(squirrel.Eq{"n.owner": owners})
	return q
}

func (q *Blobs) FilterByID(ids ...int64) data.Blobs {
	q.sql = q.sql.Where(squirrel.Eq{"n.id": ids})
	return q
}
