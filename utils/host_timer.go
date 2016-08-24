package utils

import (
	"distribute_file_system/log"
	"time"
)

const (
	time_seconde = 1
)

func Init_host_timer() {
	log.Debugf("init ping db")
	go func() {
		ch := make(chan int, 1)
		for {
			// check host status
			select {
			case <-ch:
				// notification node update
			case <-time.After(time_seconde * time.Second):
				log.Debugf("no host operation")
			}
		}
	}()

}
