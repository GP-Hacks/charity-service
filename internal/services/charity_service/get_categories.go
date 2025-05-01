package charity_service

import "context"

func (cs *CharityService) GetCategories(ctx context.Context) ([]string, error) {
	cat, err := cs.categoriesRepository.GetAll(ctx)
	if err != nil {
		return []string{}, err
	}

	s := make([]string, len(cat))
	for i, el := range cat {
		s[i] = el.Name
	}

	return s, nil
}
