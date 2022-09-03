package handlers

import (
	"net/http"

	"blob-base/resources"

	"blob-base/internal/data"

	"blob-base/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBlobsList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobsListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blobsQ := BlobsQ(r)
	blobsQ.Page(request.OffsetPageParams)
	blobs, err := blobsQ.GetBlobs()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blobs")
		ape.Render(w, problems.InternalError())
		return
	}

	blobsIds := make([]int64, len(blobs))
	for i, item := range blobs {
		blobsIds[i] = item.ID
	}

	response := resources.BlobListResponse{
		Data:  newBlobsList(blobs),
		Links: GetOffsetLinks(r, request.OffsetPageParams),
	}

	ape.Render(w, response)
}

func newBlobsList(blobs []data.Blob) []resources.Blob {
	result := make([]resources.Blob, len(blobs))
	for i, blob := range blobs {
		result[i] = newBlobModel(blob)
	}
	return result
}
