package payment

import (
	"math/rand"
	"strconv"
	"time"
)

func Create_order() (order string) {
	order = time.Now().Format("20060102150405")
	rnd := strconv.Itoa(rand.Intn(1000))
	order += ("000" + rnd)[len(rnd):]
	return
}