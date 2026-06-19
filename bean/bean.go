// Package bean provides a lightweight dependency injection container.
// Repository → Eager singleton, Service → Factory (new instance each Get).
package bean

import "sync"

// Provider is a unified object creation abstraction.
//
//	Singleton: created on first Get(), cached thereafter.
//	Factory:   new instance on every Get().
//	Eager:     created immediately, cached thereafter (for Repositories).
type Provider[T any] struct {
	once      sync.Once
	instance  T
	factory   func() T
	singleton bool
}

// Singleton creates a lazy singleton — created on first Get, cached thereafter.
func Singleton[T any](fn func() T) *Provider[T] {
	return &Provider[T]{factory: fn, singleton: true}
}

// Factory creates a new instance on every Get call.
func Factory[T any](fn func() T) *Provider[T] {
	return &Provider[T]{factory: fn, singleton: false}
}

// Eager creates an eager singleton — created immediately, returned on every Get.
// Use for Repositories that should be initialized at startup.
func Eager[T any](fn func() T) *Provider[T] {
	p := &Provider[T]{factory: fn, singleton: true}
	p.instance = fn()
	p.once.Do(func() {}) // mark once as done so Get never calls fn again
	return p
}

// Get returns the instance. Factory calls the factory each time;
// Singleton/Eager return the cached instance.
func (p *Provider[T]) Get() T {
	if !p.singleton {
		return p.factory()
	}
	p.once.Do(func() {
		p.instance = p.factory()
	})
	return p.instance
}
