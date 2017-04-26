package usethrift

import "thrifttest/duck"

func whatevs(f duck.Feeder) string {
	d := duck.NewDuck()
	f.Feed(d)
	return d.GetQuack()
}
