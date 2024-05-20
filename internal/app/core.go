package app

import (
	"aabbcc-Agent/internal/pkg"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	var c pkg.Config
	c.GetConf()

	var httpClient = pkg.HttpClient{}
	httpClient.SetHost(c.Server)

	go func() {
		for {
			fragments := make([]pkg.AspectFragment, 0)
			c.Aspects.Each(func(fragment *pkg.AspectFragment, err error) {
				if fragment != nil {
					fragments = append(fragments, *fragment)
				}
			})
			httpClient.SendAspect(fragments)
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

}
