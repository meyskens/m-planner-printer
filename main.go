package main

import (
	"machine"
	"time"

	"github.com/mect/go-escpos"
	"github.com/meyskens/m-planner-printer/pkg/api"
	"tinygo.org/x/drivers/wifinina"
)

// these are the default pins for the Arduino RP2040 Connect
var (
	uart = machine.UART0
	tx   = machine.GPIO0
	rx   = machine.GPIO1

	console = machine.Serial

	// NINA
	spi     = machine.NINA_SPI
	adaptor *wifinina.Device
)

func main() {
	print("HELLO TINY GOPHERS")

	// Configure SPI for NINA
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	adaptor = wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()

	uart.Configure(machine.UARTConfig{
		TX:       tx,
		RX:       rx,
		BaudRate: 38400, // you can get this by holding the feed button when turning on the printer
	})

	// connect to access point
	connectToAP()

	api := api.NewApi(server, key)
	p, _ := escpos.NewPrinterByUART(uart)

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Low()
		jobs, err := api.GetPrintJobs()
		if err != nil {
			println(err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		for _, job := range jobs {
			p.Init()
			uart.Write(job)
			p.Cut()
			p.End()
			p.Close()
		}

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
