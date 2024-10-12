package request

type TransportType uint

const (
	HTTP_TRANSPORT_TYPE TransportType = iota
	GRPC_TRANSPORT_TYPE
	WEBSOCKET_TRANSPORT_TYPE
)

type Context struct {
	transportType TransportType
	HttpContext   *HttpContext
}

func (ctx *Context) GetTransportType() TransportType {
	return ctx.transportType
}
