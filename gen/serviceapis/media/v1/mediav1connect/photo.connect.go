// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: serviceapis/media/v1/photo.proto

package mediav1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/mcorrigan89/media/gen/serviceapis/media/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PhotoServiceName is the fully-qualified name of the PhotoService service.
	PhotoServiceName = "serviceapis.media.v1.PhotoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PhotoServiceGetPhotoByIdProcedure is the fully-qualified name of the PhotoService's GetPhotoById
	// RPC.
	PhotoServiceGetPhotoByIdProcedure = "/serviceapis.media.v1.PhotoService/GetPhotoById"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	photoServiceServiceDescriptor            = v1.File_serviceapis_media_v1_photo_proto.Services().ByName("PhotoService")
	photoServiceGetPhotoByIdMethodDescriptor = photoServiceServiceDescriptor.Methods().ByName("GetPhotoById")
)

// PhotoServiceClient is a client for the serviceapis.media.v1.PhotoService service.
type PhotoServiceClient interface {
	GetPhotoById(context.Context, *connect.Request[v1.GetPhotoByIdRequest]) (*connect.Response[v1.GetPhotoByIdResponse], error)
}

// NewPhotoServiceClient constructs a client for the serviceapis.media.v1.PhotoService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPhotoServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PhotoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &photoServiceClient{
		getPhotoById: connect.NewClient[v1.GetPhotoByIdRequest, v1.GetPhotoByIdResponse](
			httpClient,
			baseURL+PhotoServiceGetPhotoByIdProcedure,
			connect.WithSchema(photoServiceGetPhotoByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// photoServiceClient implements PhotoServiceClient.
type photoServiceClient struct {
	getPhotoById *connect.Client[v1.GetPhotoByIdRequest, v1.GetPhotoByIdResponse]
}

// GetPhotoById calls serviceapis.media.v1.PhotoService.GetPhotoById.
func (c *photoServiceClient) GetPhotoById(ctx context.Context, req *connect.Request[v1.GetPhotoByIdRequest]) (*connect.Response[v1.GetPhotoByIdResponse], error) {
	return c.getPhotoById.CallUnary(ctx, req)
}

// PhotoServiceHandler is an implementation of the serviceapis.media.v1.PhotoService service.
type PhotoServiceHandler interface {
	GetPhotoById(context.Context, *connect.Request[v1.GetPhotoByIdRequest]) (*connect.Response[v1.GetPhotoByIdResponse], error)
}

// NewPhotoServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPhotoServiceHandler(svc PhotoServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	photoServiceGetPhotoByIdHandler := connect.NewUnaryHandler(
		PhotoServiceGetPhotoByIdProcedure,
		svc.GetPhotoById,
		connect.WithSchema(photoServiceGetPhotoByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/serviceapis.media.v1.PhotoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PhotoServiceGetPhotoByIdProcedure:
			photoServiceGetPhotoByIdHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPhotoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPhotoServiceHandler struct{}

func (UnimplementedPhotoServiceHandler) GetPhotoById(context.Context, *connect.Request[v1.GetPhotoByIdRequest]) (*connect.Response[v1.GetPhotoByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("serviceapis.media.v1.PhotoService.GetPhotoById is not implemented"))
}
