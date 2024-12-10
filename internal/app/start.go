package app

import (
	"errors"
	"io"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
)

type starter interface {
	Start() error
}

type starterCloser interface {
	starter
	io.Closer
}

func (a *App) addStarter(i starter) {
	if isNil(i) {
		panic("nil starter")
	}
	a.starters = append(a.starters, i)
}

func (a *App) addCloser(i io.Closer) {
	if isNil(i) {
		panic("nil closer")
	}
	a.closers = append(a.closers, i)
}

func (a *App) addComponent(component any) {
	if isNil(component) {
		panic("nil component")
	}
	switch cmpType := component.(type) {
	case starterCloser:
		a.addStarter(cmpType)
		a.closers = append(a.closers, cmpType)
	case starter:
		a.addStarter(cmpType)
	case io.Closer:
		a.addCloser(cmpType)
	}
}

func (a *App) Start() error {
	if len(a.starters) == 0 {
		return errors.New("no starters")
	}
	ch := make(chan error)
	var err error
	a.gracefulShutdown()

	wg := &sync.WaitGroup{}
	for _, s := range a.starters {
		wg.Add(1)
		go func(s starter) {
			defer wg.Done()
			startErr := s.Start()
			if startErr != nil {
				ch <- startErr
			}
		}(s)
	}
	go func() {
		chErr, ok := <-ch
		if !ok {
			return
		}
		err = chErr
	}()
	wg.Wait()
	close(ch)
	if err != nil {
		a.startErrCh <- err
	}
	return err
}

func (a *App) close() error {
	for _, c := range a.closers {
		err := c.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) gracefulShutdown() {
	sc := make(chan os.Signal, 1)
	go func() {
		for {
			select {
			case <-sc:
				_ = a.close()
				break
			case <-a.startErrCh:
				_ = a.close()
				break
			}
		}
	}()
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	os.Exit(0)
}

func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}
