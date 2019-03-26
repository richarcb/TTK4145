package FSM

import (
	"time"
	//"fmt"
)


func DoorTimer(close_door_ch chan<- bool, reset_timer <-chan bool) {
	const doorOpenTime = 3 * time.Second
  //declare timer
	timer := time.NewTimer(0)
	timer.Stop()

	for {
		select {
		case <-reset_timer:
			timer.Reset(doorOpenTime)
		case <-timer.C:
			timer.Stop()
			close_door_ch <- true
		}
	}
}
/*
func Powerloss_timer(power_loss_ch chan<- bool, reset_timer_pl <-chan bool) {
	const no_floor_time = 5 * time.Second
  //declare timer
	timer := time.NewTimer(0)
	timer.Stop()

	for {
		select {
		case <-reset_timer_pl:
			timer.Reset(no_floor_time)
		case <-timer.C:
			timer.Stop()
			power_loss_ch <- true
		}
	}
}
*/
