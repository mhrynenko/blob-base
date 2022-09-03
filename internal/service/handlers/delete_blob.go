package handlers

import (
	"blob-base/resources"
	"net/http"

	"blob-base/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, err := BlobsQ(r).DeleteBlob(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if blob == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := resources.BlobResponse{
		Data: newBlobModel(*blob),
	}
	ape.Render(w, response)
}
