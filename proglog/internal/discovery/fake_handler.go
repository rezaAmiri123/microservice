package discovery

var _ Handler = (*FakeHandler)(nil)

type FakeHandler struct {
	Joins  chan map[string]string
	Leaves chan string
}

func (h *FakeHandler) Join(id, addr string) error {
	if h.Joins != nil {
		h.Joins <- map[string]string{
			"id":   id,
			"addr": addr,
		}
	}

	return nil
}

func (h *FakeHandler) Leave(id string) error {
	if h.Leaves != nil {
		h.Leaves <- id
	}
	return nil
}
