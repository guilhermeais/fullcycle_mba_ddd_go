package common

type UnitOfWork struct {
	aggregateRoots []AggregateRoot
}

func (uow *UnitOfWork) GetAggregateRoots() []AggregateRoot {
	return uow.aggregateRoots
}

func (uow *UnitOfWork) RegisterAggregate(aggregates ...AggregateRoot) {
	uow.aggregateRoots = append(uow.aggregateRoots, aggregates...)
}
