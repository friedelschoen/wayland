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

type Registrar map[string]Proxy

func (proxies Registrar) Add(entries ...Proxy) {
	for _, p := range entries {
		proxies[p.Name()] = p
	}
}

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
