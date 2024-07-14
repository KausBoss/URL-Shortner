package shortner

import "tinyURL/internal/storage/memory"

type Service struct {
	store memory.Store
}

func New(store *memory.Store) *Service {
	return &Service{store: *store}
}

func (s *Service) Shorten(url string) (string, error) {
	return s.store.Save(url)
}

func (s *Service) Expand(code string) (string, error) {
	return s.store.Load(code)
}
