package errors

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"

	"github.com/gogo/status"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

// CustomContentTypeMarshaler for customize content-type
type CustomContentTypeMarshaler interface {
	// ContentTypeFromMessage returns the Content-Type this marshaler produces from the provided message
	ContentTypeFromMessage(v interface{}) string
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorBodyWithDetails struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func statusToErrorBody(s *status.Status) interface{} {

	domainErr, ok := FromRPCStatus(s)
	if !ok {
		return errorBody{
			Code:    "13",
			Message: s.Message(),
		}
	}

	if domainErr.Details == nil || len(domainErr.Details) == 0 {
		return errorBody{
			Code:    domainErr.Code,
			Message: domainErr.Message,
		}
	}

	return errorBodyWithDetails{
		Code:    domainErr.Code,
		Message: domainErr.Message,
		Details: domainErr.Details,
	}
}

// CustomHTTPError for customizing error returning of gRPC gateway
func CustomHTTPError(
	ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler,
	w http.ResponseWriter, _ *http.Request, err error,
) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")

	contentType := marshaler.ContentType()
	// Check marshaler on run time in order to keep backwards compatibility
	// An interface param needs to be added to the ContentType() function on
	// the Marshal interface to be able to remove this check
	pb := s.Proto()
	if typeMarshaler, ok := marshaler.(CustomContentTypeMarshaler); ok {
		contentType = typeMarshaler.ContentTypeFromMessage(pb)
	}
	w.Header().Set("Content-Type", contentType)

	body := statusToErrorBody(s)

	buf, merr := marshaler.Marshal(body)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", body, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	handleForwardResponseServerMetadata(w, mux, md)
	handleForwardResponseTrailerHeader(w, md)
	st := runtime.HTTPStatusFromCode(s.Code())
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

	handleForwardResponseTrailer(w, md)
}

func outgoingHeaderMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", runtime.MetadataHeaderPrefix, key), true
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, mux *runtime.ServeMux, md runtime.ServerMetadata) {
	for k, vs := range md.HeaderMD {
		if h, ok := outgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}
}

func handleForwardResponseTrailerHeader(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k := range md.TrailerMD {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, k))
		w.Header().Add("Trailer", tKey)
	}
}

func handleForwardResponseTrailer(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.TrailerMD {
		tKey := fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, k)
		for _, v := range vs {
			w.Header().Add(tKey, v)
		}
	}
}
