package modules

import "github.com/ngrok/ngrok-go/internal/tunnel/proto"

type TCPOption interface {
	ApplyTCP(cfg *tcpOptions)
}

type tcpOptionFunc func(cfg *tcpOptions)

func (of tcpOptionFunc) ApplyTCP(cfg *tcpOptions) {
	of(cfg)
}

// Construct a new set of HTTP tunnel options.
func TCPOptions(opts ...TCPOption) TunnelOptions {
	cfg := tcpOptions{}
	for _, opt := range opts {
		opt.ApplyTCP(&cfg)
	}
	return cfg
}

// The options for a TCP edge.
type tcpOptions struct {
	// Common tunnel configuration options.
	commonOpts
	// The TCP address to request for this edge.
	RemoteAddr string
}

// Set the TCP address to request for this edge.
func WithRemoteAddr(addr string) TCPOption {
	return tcpOptionFunc(func(cfg *tcpOptions) {
		cfg.RemoteAddr = addr
	})
}

func (cfg *tcpOptions) toProtoConfig() *proto.TCPOptions {
	return &proto.TCPOptions{
		Addr:          cfg.RemoteAddr,
		IPRestriction: cfg.commonOpts.CIDRRestrictions.toProtoConfig(),
		ProxyProto:    proto.ProxyProto(cfg.commonOpts.ProxyProto),
	}
}

func (cfg tcpOptions) tunnelOptions() {}

func (cfg tcpOptions) ForwardsTo() string {
	return cfg.commonOpts.getForwardsTo()
}
func (cfg tcpOptions) Extra() proto.BindExtra {
	return proto.BindExtra{
		Metadata: cfg.Metadata,
	}
}
func (cfg tcpOptions) Proto() string {
	return "tcp"
}
func (cfg tcpOptions) Opts() any {
	return cfg.toProtoConfig()
}
func (cfg tcpOptions) Labels() map[string]string {
	return nil
}