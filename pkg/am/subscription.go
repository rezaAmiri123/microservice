package am

type Subscription interface {
	Unsubscribe() error
}

type SubscriptionFunc func() error

func (f SubscriptionFunc) Unsubscribe() error  {
	return f()
}