package domain

type DomainType interface {
	Webhook |
		Event
}

type Adapt[T any, D DomainType] interface {
	AdaptTarget(domain D) (target T)
	AdaptDomain(target T) (domain D)
}
