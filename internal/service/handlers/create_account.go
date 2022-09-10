package handlers

import (
	"blob-base/internal/service/requests"
	"blob-base/resources"
	"bytes"
	"encoding/json"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	keypair2 "gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	"io"
	"net/http"
)

type CoreTxRequest struct {
	Tx string `json:"tx"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateAccountRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	txBuilder := xdrbuild.NewBuilder("TokenD Developer Network", 601200)
	tx := txBuilder.Transaction(keypair.MustParseAddress(request.Data.Attributes.SourceAccount))

	tx = tx.Op(&xdrbuild.CreateAccount{
		Destination: request.Data.Attributes.Operation.Body.CreateAccountOp.Destination,
		RoleID:      request.Data.Attributes.Operation.Body.CreateAccountOp.RoleId,
		Signers:     getSigners(request.Data.Attributes.Operation.Body.CreateAccountOp.SignersData),
	})

	envelopeString, err := tx.Marshal()
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "failed to build tx envelope"))...)
		return
	}

	var envelope xdr.TransactionEnvelope
	err = envelope.Scan(envelopeString)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "failed to extract envelope info"))...)
		return
	}

	txEnvelope, err := addSignatureAndHash(envelope, txBuilder, "SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4")

	data, err := json.Marshal(CoreTxRequest{Tx: txEnvelope})
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "failed to encode envelope request"))...)
		return
	}

	respBody, err := postData(data, "SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4")
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(errors.Wrap(err, "failed to send post envelope"))...)
		return

	}

	w.Header().Set("content-type", "application/vnd.api+json")
	w.Write(respBody)
}

func postData(body []byte, seed string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8001/v3/transactions", bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create envelope request")
	}

	err = signcontrol.SignRequest(req, keypair2.MustParse(seed))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign envelope request")
	}

	return sendRequest(req)
}

func sendRequest(req *http.Request) ([]byte, error) {
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make envelope request")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read envelope response body")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "notification request not accepted")
	}

	return body, nil
}

func addSignatureAndHash(envelope xdr.TransactionEnvelope, txBuilder *xdrbuild.Builder, seed string) (string, error) {
	signedEnvelope, err := txBuilder.Sign(&envelope, keypair.MustParseSeed(seed))
	if err != nil {
		return "", errors.Wrap(err, "failed to sign tx")
	}

	tx, err := xdr.MarshalBase64(signedEnvelope)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal tx envelope")
	}

	return tx, nil
}

func getSigners(accountSigners []resources.SignerData) []xdrbuild.SignerData {
	var signers []xdrbuild.SignerData
	for _, accountSigner := range accountSigners {
		signers = append(signers, xdrbuild.SignerData{
			PublicKey: accountSigner.PublicKey,
			RoleID:    accountSigner.RoleId,
			Weight:    accountSigner.Weight,
			Identity:  accountSigner.Identity,
			Details:   Details{},
		})
	}
	return signers
}

type Details map[string]interface{}

func (d Details) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(d))
}
