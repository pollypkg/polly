// package fuge wraps the centrifuge API to make using it more natural in Go:
// - func values instead of full types
// - forced synchronization when connecting and subscribing (and thus error return values)
package fuge

import (
	"errors"
	"fmt"
	"log"

	"github.com/centrifugal/centrifuge-go"
)

func New(u string, config centrifuge.Config) *Fuge {
	c := *centrifuge.New(u, config)
	f := Fuge(c)
	return &f
}

type (
	Real = *centrifuge.Client
	Fuge centrifuge.Client
)

func (f *Fuge) Connect() error {
	ch := make(chan error, 1)
	defer close(ch)

	f.OnConnect(func(c *centrifuge.Client, ce centrifuge.ConnectEvent) {
		ch <- nil
	})
	f.OnError(func(c *centrifuge.Client, ee centrifuge.ErrorEvent) {
		ch <- errors.New(ee.Message)
	})

	if err := Real(f).Connect(); err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	if err := <-ch; err != nil {
		return err
	}

	f.OnError(func(c *centrifuge.Client, ee centrifuge.ErrorEvent) {
		log.Println("error:", ee.Message)
	})

	return nil
}

func (f *Fuge) Token(token string) {
	Real(f).SetHeader("Authorization", "Bearer "+token)
}

func (f *Fuge) OnConnect(fn func(*centrifuge.Client, centrifuge.ConnectEvent)) {
	Real(f).OnConnect(handler{onConnect: fn})
}
func (f *Fuge) OnError(fn func(*centrifuge.Client, centrifuge.ErrorEvent)) {
	Real(f).OnError(handler{onError: fn})
}

type (
	Sub     centrifuge.Subscription
	RealSub = *centrifuge.Subscription
)

func (f *Fuge) Sub(channel string, onPub func(*centrifuge.Subscription, centrifuge.PublishEvent)) (*Sub, error) {
	rs, err := Real(f).NewSubscription(channel)
	if err != nil {
		return nil, err
	}

	tmp := Sub(*rs)
	s := &tmp

	RealSub(s).OnPublish(handler{onPublish: onPub})
	if err := s.Subscribe(); err != nil {
		return nil, fmt.Errorf("subscribing to '%s': %w", channel, err)
	}

	return s, nil
}

func (s *Sub) Subscribe() error {
	ch := make(chan error, 1)
	s.OnSubscribeError(func(s *centrifuge.Subscription, e centrifuge.SubscribeErrorEvent) {
		ch <- errors.New(e.Error)
	})
	s.OnSubscribeSuccess(func(s *centrifuge.Subscription, sse centrifuge.SubscribeSuccessEvent) {
		ch <- nil
	})

	if err := RealSub(s).Subscribe(); err != nil {
		return err
	}

	return <-ch
}

func (s *Sub) OnSubscribeSuccess(fn func(*centrifuge.Subscription, centrifuge.SubscribeSuccessEvent)) {
	RealSub(s).OnSubscribeSuccess(handler{onSubscribeSuccess: fn})
}
func (s *Sub) OnSubscribeError(fn func(*centrifuge.Subscription, centrifuge.SubscribeErrorEvent)) {
	RealSub(s).OnSubscribeError(handler{onSubscribeError: fn})
}

func (s *Sub) Close() error {
	return RealSub(s).Close()
}

type handler struct {
	onConnect func(*centrifuge.Client, centrifuge.ConnectEvent)
	onError   func(*centrifuge.Client, centrifuge.ErrorEvent)

	onSubscribeSuccess func(*centrifuge.Subscription, centrifuge.SubscribeSuccessEvent)
	onSubscribeError   func(*centrifuge.Subscription, centrifuge.SubscribeErrorEvent)
	onPublish          func(*centrifuge.Subscription, centrifuge.PublishEvent)
}

var (
	_ centrifuge.ConnectHandler = handler{}
	_ centrifuge.ErrorHandler   = handler{}

	_ centrifuge.SubscribeErrorHandler   = handler{}
	_ centrifuge.SubscribeSuccessHandler = handler{}
	_ centrifuge.PublishHandler          = handler{}
)

func (h handler) OnConnect(c *centrifuge.Client, e centrifuge.ConnectEvent) {
	h.onConnect(c, e)
}
func (h handler) OnError(c *centrifuge.Client, e centrifuge.ErrorEvent) {
	h.onError(c, e)
}

func (h handler) OnSubscribeSuccess(sub *centrifuge.Subscription, e centrifuge.SubscribeSuccessEvent) {
	h.onSubscribeSuccess(sub, e)
}
func (h handler) OnSubscribeError(sub *centrifuge.Subscription, e centrifuge.SubscribeErrorEvent) {
	h.onSubscribeError(sub, e)
}
func (h handler) OnPublish(sub *centrifuge.Subscription, e centrifuge.PublishEvent) {
	h.onPublish(sub, e)
}
