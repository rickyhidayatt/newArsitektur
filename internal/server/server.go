package server

import (
	"bni.co.id/xpora/medias/internal/server/decoding"
	"bni.co.id/xpora/medias/internal/server/encoding"
	"github.com/go-kit/kit/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
)

type Option struct {
	DecodeModel interface{}
	Encoder     kitHttp.EncodeResponseFunc
	Decoder     kitHttp.DecodeRequestFunc
}

func NewHttpServer(endpoint endpoint.Endpoint, option Option, serverOption []kitHttp.ServerOption) *kitHttp.Server {
	if option.Encoder == nil {
		option.Encoder = encoding.Encode()
	}
	if option.Decoder == nil {
		option.Decoder = decoding.Decode(option.DecodeModel)
	}

	return kitHttp.NewServer(endpoint, option.Decoder, option.Encoder, serverOption...)
}
