package wayland

import (
	"slices"
)

type RegistryProxy interface {
	Proxy
	Bind(uint32, string, uint32, Proxy)
}

type BindInterfaceEvent interface {
	Event
	BindInterface() (uint32, string, uint32)
}

type Registrar []Proxy

func (proxies Registrar) GetInterface(name string) (int, Proxy) {
	for i, p := range proxies {
		if p.Name() == name {
			return i, p
		}
	}
	return -1, nil
}

func (proxies *Registrar) Handler(evt Event) {
	e := evt.(BindInterfaceEvent)
	r := evt.Proxy().(RegistryProxy)

	name, iface, version := e.BindInterface()

	i, p := proxies.GetInterface(iface)
	if p == nil {
		return
	}

	r.Conn().Register(p)
	r.Bind(name, iface, version, p)
	*proxies = slices.Delete(*proxies, i, i)
}
