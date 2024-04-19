package app

import (
	"aabbcc-Agent/internal/pkg"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	var c pkg.Config
	c.GetConf()

	var httpClient = pkg.HttpClient{}
	httpClient.SetHost(c.Server)

	for {

		for _, aspect := range c.Aspects {
			if aspect.NextTime.Before(time.Now()) || aspect.NextTime.IsZero() {
				fragment, err := aspect.FetchData()
				if err != nil {
					//fmt.Println("some error: " + err.Error())
				} else {
					httpClient.SendAspect(fragment)
				}
			}

		}
		time.Sleep(time.Millisecond * 1)
	}
}
