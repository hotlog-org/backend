package service

type ProjectStore interface {
	GetProjectEventsCount(id string) (int64, error)
}

type ProjectService struct {
	store ProjectStore
}

func NewProjectService(store ProjectStore) *ProjectService {
	return &ProjectService{store: store}
}

func (s *ProjectService) GetProjectEventsCount(id string) (int64, error) {
	return s.store.GetProjectEventsCount(id)
}
