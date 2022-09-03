package handlers

import (
	"blob-base/internal/data"
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobsQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func BlobsQ(r *http.Request) data.Blobs {
	return r.Context().Value(blobsQCtxKey).(data.Blobs).New()
}

func CtxBlobsQ(entry data.Blobs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobsQCtxKey, entry)
	}
}
