package common

type UnitOfWork struct {
	aggregateRoots []AggregateRoot
}

func (uow *UnitOfWork) GetAggregateRoots() []AggregateRoot {
	return uow.aggregateRoots
}

func (uow *UnitOfWork) RegisterAggregate(aggregates ...AggregateRoot) {
	for _, ag := range aggregates {
		for _, alreadyRegisteredAg := range uow.aggregateRoots {
			if ag.id.IsEqual(alreadyRegisteredAg.id) {
				return
			}
		}
	}
	uow.aggregateRoots = append(uow.aggregateRoots, aggregates...)
}
