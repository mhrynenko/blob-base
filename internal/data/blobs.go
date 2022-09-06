package data

import "gitlab.com/distributed_lab/kit/pgdb"

type Blobs interface {
	New() Blobs

	NewBlob(blob Blob) (Blob, error)
	GetBlob() (*Blob, error)
	DeleteBlob(blobId int64) error

	GetBlobs() ([]Blob, error)

	Page(pageParams pgdb.OffsetPageParams) Blobs

	FilterByOwner(owners ...string) Blobs
	FilterByID(ids ...int64) Blobs
}

type Blob struct {
	ID         int64  `db:"id" structs:"-"`
	Owner      string `db:"owner" structs:"-"`
	Attributes []byte `db:"attributes" structs:"-"`
}
