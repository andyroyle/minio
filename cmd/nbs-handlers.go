package cmd

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func (api blockStorageAPIHandlers) CreateVolumeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "CreateVolume")

	blockStorageAPI := api.BlockStorageAPI()

	if blockStorageAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	if err := reqSignatureV4Verify(r, globalServerConfig.GetRegion(), serviceEC2); err != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(err), r.URL, guessIsBrowserReq(r))
		return
	}

	createVolume := blockStorageAPI.CreateVolume

	vars := mux.Vars(r)

	size, _ := strconv.ParseInt(vars["size"], 10, 32)
	iops, _ := strconv.ParseInt(vars["iops"], 10, 32)

	/*
	 */
	volumeInfo, err := createVolume(ctx,
		size,
		vars["availabilityZone"],
		vars["snapshotId"],
		iops,
		vars["kmsKeyId"],
		vars["volumeType"])

	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	response := generateCreateVolumeResponse(volumeInfo)
	response.RequestId = w.Header().Get(responseRequestIDKey)

	// Generate response.
	encodedSuccessResponse := encodeResponse(response)

	// Write response.
	writeSuccessResponseXML(w, encodedSuccessResponse)
}

func (api blockStorageAPIHandlers) DeleteVolumeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "DeleteVolume")

	blockStorageAPI := api.BlockStorageAPI()

	if blockStorageAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	if err := reqSignatureV4Verify(r, globalServerConfig.GetRegion(), serviceEC2); err != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(err), r.URL, guessIsBrowserReq(r))
		return
	}

	deleteVolume := blockStorageAPI.DeleteVolume

	vars := mux.Vars(r)

	/*
	 */
	err := deleteVolume(ctx, vars["volumeId"])

	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	response := DeleteVolumeResponse{
		RequestId: w.Header().Get(responseRequestIDKey),
		Return:    true,
	}

	// Generate response.
	encodedSuccessResponse := encodeResponse(response)

	// Write response.
	writeSuccessResponseXML(w, encodedSuccessResponse)
}

func (api blockStorageAPIHandlers) AttachVolumeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "AttachVolume")

	blockStorageAPI := api.BlockStorageAPI()

	if blockStorageAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL, guessIsBrowserReq(r))
		return
	}

	if err := reqSignatureV4Verify(r, globalServerConfig.GetRegion(), serviceEC2); err != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(err), r.URL, guessIsBrowserReq(r))
		return
	}

	attachVolume := blockStorageAPI.AttachVolume

	vars := mux.Vars(r)

	/*
	 */
	err := attachVolume(ctx, vars["volumeId"], vars["instanceId"], vars["device"])

	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL, guessIsBrowserReq(r))
		return
	}

	response := AttachVolumeResponse{
		RequestId:  w.Header().Get(responseRequestIDKey),
		VolumeId:   vars["volumeId"],
		InstanceId: vars["instanceId"],
		Device:     vars["device"],
	}

	// Generate response.
	encodedSuccessResponse := encodeResponse(response)

	// Write response.
	writeSuccessResponseXML(w, encodedSuccessResponse)
}
