package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
    "net"
    "time"
)

const (
    deviceName = "en0" // for mac
)

var (
    DownStreamDataSize = 0 // 单位时间内下行的总字节数
    UpStreamDataSize   = 0 // 单位时间内上行的总字节数
)

func StartCatchSpeed() {
    ifs, err := pcap.FindAllDevs()
    checkErr(err)
    var device pcap.Interface
    for _, d := range ifs {
        if d.Name == deviceName {
            device = d
        }
    }
    ipv4 := findDeviceIpv4(device)
    fmt.Println("--------",ipv4)
    macAddr := findMacAddrByIp(ipv4)
    fmt.Println("ipv4:", ipv4)
    fmt.Println("mac:", macAddr)

    // 获取网卡handler，可用于读取或写入数据包
    handle, err := pcap.OpenLive(deviceName, 1024 /*每个数据包读取的最大值*/, true /*是否开启混杂模式*/, 30*time.Second /*读包超时时长*/)
    if err != nil {
        panic(err)
    }
    defer handle.Close()

    // 开始抓包
    monitor(handle, macAddr)
}

func monitor(handle *pcap.Handle, macAddr string) {
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        // 只获取以太网帧
        ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
        if ethernetLayer != nil {
            ethernet := ethernetLayer.(*layers.Ethernet)
            // 如果封包的目的MAC是本机则表示是下行的数据包，否则为上行
            if ethernet.DstMAC.String() == macAddr {
                DownStreamDataSize += len(packet.Data()) // 统计下行封包总大小
            } else {
                UpStreamDataSize += len(packet.Data()) // 统计上行封包总大小
            }
        }
    }
}

// 获取网卡的IPv4地址
func findDeviceIpv4(device pcap.Interface) string {
    for _, addr := range device.Addresses {
        if ipv4 := addr.IP.To4(); ipv4 != nil {
            return ipv4.String()
        }
    }
    panic("device has no IPv4")
}

// 根据网卡的IPv4地址获取MAC地址
// 有此方法是因为gopacket内部未封装获取MAC地址的方法，所以这里通过找到IPv4地址相同的网卡来寻找MAC地址
func findMacAddrByIp(ip string) (string) {
    interfaces, err := net.Interfaces()
    if err != nil {
        panic(interfaces)
    }

    for _, i := range interfaces {
        addrs, err := i.Addrs()
        if err != nil {
            panic(err)
        }

        for _, addr := range addrs {
            if a, ok := addr.(*net.IPNet); ok {
                if ip == a.IP.String() {
                    return i.HardwareAddr.String()
                }
            }
        }
    }
    panic(fmt.Sprintf("no device has given ip: %s", ip))
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    // 开启网速抓包
    go StartCatchSpeed()
}