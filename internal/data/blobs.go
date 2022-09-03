package data

import "gitlab.com/distributed_lab/kit/pgdb"

type Blobs interface {
	New() Blobs

	NewBlob(blob Blob) (Blob, error)
	GetBlob(blobId int64) (*Blob, error)
	DeleteBlob(blobId int64) (*Blob, error)

	GetBlobs() ([]Blob, error)

	Page(pageParams pgdb.OffsetPageParams) Blobs
}

type Blob struct {
	ID         int64  `db:"id" structs:"-"`
	Attributes []byte `db:"attributes" structs:"-"`
}
