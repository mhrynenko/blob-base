package postgres

import (
	"blob-base/internal/data"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const tableName = "blobs"

type Blobs struct {
	database *pgdb.DB
	sql      squirrel.SelectBuilder
}

var selectedTable = squirrel.Select("*").From(tableName)

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
	query := squirrel.Insert(tableName).Columns("attributes").Values(blob.Attributes).Suffix("returning *")

	var result data.Blob
	err := q.database.Get(&result, query)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *Blobs) GetBlob(blobId int64) (*data.Blob, error) {
	var result data.Blob

	query := q.sql.Where("id = ?", blobId).
		PlaceholderFormat(squirrel.Dollar)

	err := q.database.Get(&result, query)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *Blobs) DeleteBlob(blobId int64) (*data.Blob, error) {
	blob, err := q.GetBlob(blobId)
	if err != nil {
		return blob, err
	}

	query := squirrel.Delete(tableName).Where("id = ?", blobId)

	err = q.database.Exec(query)
	if err != nil {
		return blob, err
	}

	return blob, nil
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
