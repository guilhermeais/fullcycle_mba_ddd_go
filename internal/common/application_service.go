package common

func Run[T any](fn func() T, uow *UnitOfWork, eventManager *DomainEventManager) T {
	result := fn()
	aggregates := uow.GetAggregateRoots()
	for _, agg := range aggregates {
		eventManager.Publish(agg)
	}
	return result
}
