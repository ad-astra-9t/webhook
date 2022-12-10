package domain

type DomainType interface {
	Webhook |
		Event
}

type Adapt[D DomainType, T any] interface {
	AdaptTarget(domain D) (target T)
	AdaptDomain(target T) (domain D)
}
