package hub

// A ServerOption sets options such as credentials, codec and keepalive parameters, etc.
type ClientOption interface {
	apply(*clientOptions)
}

type clientOptions struct {
	userInfo *UserInfo
	onWrite  func(msg string) error
}

var defaultClientOptions = clientOptions{}

// funcClientOption wraps a function that modifies funcClientOption into an
// implementation of the ServerOption interface.
type funcClientOption struct {
	f func(*clientOptions)
}

func (fdo *funcClientOption) apply(do *clientOptions) {
	fdo.f(do)
}

func newFuncClientOption(f func(*clientOptions)) *funcClientOption {
	return &funcClientOption{
		f: f,
	}
}

func NewHubClient(protocol Protocol, opt ...ClientOption) Client {
	var client Client

	opts := defaultClientOptions
	for _, o := range opt {
		o.apply(&opts)
	}

	switch protocol {
	case GraphQL:
		client = newGraphQLClient(opts)
	}

	return client
}

func OnWrite(f func(msg string) error) ClientOption {
	return newFuncClientOption(func(co *clientOptions) {
		co.onWrite = f
	})
}

func SetUserInfo(info *UserInfo) ClientOption {
	return newFuncClientOption(func(co *clientOptions) {
		co.userInfo = info
	})
}
