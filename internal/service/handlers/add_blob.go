package handlers

import (
	"encoding/json"
	"net/http"

	"blob-base/internal/data"
	"blob-base/internal/service/requests"
	"blob-base/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AddBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAddBlobRequest(r)

	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var resultBlob data.Blob
	newBlob := data.Blob{
		Attributes: request.Data.Attributes,
	}
	resultBlob, err = BlobsQ(r).NewBlob(newBlob)

	if err != nil {
		Log(r).WithError(err).Error(err, "failed to insert blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.BlobResponse{
		Data: newBlobModel(resultBlob),
	}
	ape.Render(w, result)
}

func newBlobModel(blob data.Blob) resources.Blob {
	result := resources.Blob{
		Key:        resources.NewKeyInt64(blob.ID, resources.BLOB),
		Attributes: json.RawMessage(blob.Attributes),
	}

	return result
}
