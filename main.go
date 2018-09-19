package main

import (
	"flag"
	"net"
	"os"

	"github.com/mdlayher/arp"
)

func main() {
	var mac string
	var ipstr string
	flag.StringVar(&mac, "macaddr", "", "arp replay mac address ff:ff:ff:ff:ff:ff")
	flag.StringVar(&ipstr, "ipaddr", "", "arp replay ipaddress 192.168.0.1")
	flag.Parse()

	if mac == "" || ipstr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ifce, err := net.InterfaceByName("bond0")
	if err != nil {
		panic(err)
	}

	c, err := arp.Dial(ifce)
	if err != nil {
		panic(err)
	}

	ip := net.ParseIP(ipstr).To4()
	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		panic(err)
	}
	destIP := net.ParseIP("0.0.0.0").To4()
	destHwAddr, err := net.ParseMAC("ff:ff:ff:ff:ff:ff")
	if err != nil {
		panic(err)
	}
	pkt, err := arp.NewPacket(arp.OperationReply, destHwAddr, destIP, hwAddr, ip)
	if err != nil {
		panic(err)
	}

	c.Reply(pkt, hwAddr, ip)

}
