package main

import (
	"fmt"
	"math/rand"
	"time"
)
//https://youtu.be/pokwMUv5vwc

func main(){
	ticker :=time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var alertTimerChan <-chan time.Time
	alertTimerActive:=false

	go func ()  {
		for {
			select{
			case <-ticker.C:
				cpuUsage:=rand.Intn(100)+100
				//fmt.Println("Tick at:",t)
				fmt.Printf("CPU usage: %d%%\n",cpuUsage)
				
				if cpuUsage>60 {
					if !alertTimerActive{
						fmt.Println("High CPU usage detected, starting alert timer.")
						alertTimer :=time.NewTicker(5*time.Second)
						alertTimerChan=alertTimer.C
						alertTimerActive=true				
					
					}

				} else{
					if alertTimerActive {
						fmt.Println("CPU usage normalized, stopping alert timer")
						alertTimerActive=false
						alertTimerChan=nil
					}

				}
			case <-alertTimerChan:
				if alertTimerActive {
					fmt.Println("ALERT: High CPU usage sustained for 10 seconds!")
					alertTimerActive=false
					alertTimerChan=nil
				}
			}
		}
	}()

	time.Sleep(20* time.Second)
	


}