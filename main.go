package main

import (
	"fmt"
)

type Subscriber func(chn Channel, ch <-chan interface{})

func NewSubscriber() Subscriber {
	return func(chn Channel, ch <-chan interface{}) {
		fmt.Println(chn.Name + " starts")
		for {
			x := <-ch
			fmt.Println(chn.Name+" received :", x)
		}
	}
}

type Channel struct {
	Name string
}

func main() {

	var subs = []Channel{Channel{Name: "subscriber1"},
		Channel{Name: "subscriber2"},
		Channel{Name: "subscriber3"}}

	mp := map[string]chan interface{}{}

	for _, v := range subs {
		mp[v.Name] = make(chan interface{})
		go NewSubscriber()(v, mp[v.Name])
	}

	for {
		x := make([]string, 0)
		for {
			var w string
			fmt.Scanf("%s", &w)
			if w != "" {
				x = append(x, w)
			} else {
				break
			}
		}
		for _, v := range subs {
			mp[v.Name] <- x
		}
	}
}
