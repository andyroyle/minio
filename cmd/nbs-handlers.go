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

	if err := reqSignatureV4Verify(r, globalServerConfig.GetRegion(), serviceEC2); err != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(err), r.URL, guessIsBrowserReq(r))
		return
	}

	describeVolumes := blockStorageAPI.DescribeVolumes

	/*
	 */
	volumesInfo, err := describeVolumes(ctx)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	response := generateDescribeVolumesResponse(volumesInfo)
	response.RequestId = w.Header().Get(responseRequestIDKey)

	// Generate response.
	// response := generateListBucketsResponse(bucketsInfo)
	encodedSuccessResponse := encodeResponse(response)

	// Write response.
	writeSuccessResponseXML(w, encodedSuccessResponse)
}
