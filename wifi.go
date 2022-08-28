package main

import (
	"time"

	"tinygo.org/x/drivers/net"
)

// connect to access point
func connectToAP() {
	adaptor.Disconnect()
	net.ActiveDevice = nil // dirty hack to reset the NINA
	adaptor.Configure()
	time.Sleep(time.Second)
	println("Connecting to " + ssid)

	err := adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
	if err != nil { // error connecting to AP
		for {
			println(err)
			time.Sleep(1 * time.Second)
		}
	}

	println("Connected.")

	time.Sleep(2 * time.Second)
	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		println(err.Error())
		time.Sleep(1 * time.Second)
	}
	println(ip.String())
}
