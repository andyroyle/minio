package cmd

import "net/http"

func (api blockStorageAPIHandlers) DescribeVolumesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "DescribeVolumes")

	// defer logger.AuditLog(w, r, "DescribeVolumes", mustGetClaimsFromToken(r))
	blockStorageAPI := api.BlockStorageAPI()

	if blockStorageAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	describeVolumes := blockStorageAPI.DescribeVolumes

	/*
	 */
	response, err := describeVolumes(ctx)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	// Generate response.
	// response := generateListBucketsResponse(bucketsInfo)
	encodedSuccessResponse := encodeResponse(response)

	// Write response.
	writeSuccessResponseXML(w, encodedSuccessResponse)
}
