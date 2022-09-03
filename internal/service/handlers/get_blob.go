package handlers

import (
	"net/http"

	"blob-base/internal/service/requests"
	"blob-base/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	notification, err := BlobsQ(r).GetBlob(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if notification == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	result := resources.BlobResponse{
		Data: newBlobModel(*notification),
	}

	ape.Render(w, result)
}
