package wayland

type registryProxy interface {
	Proxy
	Bind(uint32, string, uint32, Proxy)
}

type registryGlobalEvent interface {
	Event
	Name() uint32
	Interface() string
	Version() uint32
}

// Registrar is an object containing multiple proxies which can be registered by wl_registry::global.
type Registrar map[string]Proxy

// Add adds one or more proxies to the registrar.
func (proxies Registrar) Add(entries ...Proxy) {
	for _, p := range entries {
		proxies[p.Name()] = p
	}
}

// Handler should be set as handler for wl_registry::global.
func (proxies Registrar) Handler(evt Event) bool {
	e := evt.(registryGlobalEvent)
	r := evt.Proxy().(registryProxy)

	p, ok := proxies[e.Interface()]
	if !ok {
		return false
	}

	r.Conn().Register(p)
	r.Bind(e.Name(), e.Interface(), e.Version(), p)
	delete(proxies, p.Name())
	return true
}
