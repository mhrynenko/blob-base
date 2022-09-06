package handlers

import (
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

	err = BlobsQ(r).DeleteBlob(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusOK)
}
