package main

import (
    "errors"
    "fmt"
    "net"
    "github.com/StackExchange/wmi"
    "unsafe"
    "syscall"
)

type cpuInfo struct {
    Name          string
    NumberOfCores uint32
    ThreadCount   uint32
}

type operatingSystem struct {
    Name    string
    Version string
}

type memoryStatusEx struct {
    cbSize                  uint32
    dwMemoryLoad            uint32
    ullTotalPhys            uint64 // in bytes
    ullAvailPhys            uint64
    ullTotalPageFile        uint64
    ullAvailPageFile        uint64
    ullTotalVirtual         uint64
    ullAvailVirtual         uint64
    ullAvailExtendedVirtual uint64
}

type Storage struct {
    Name       string
    FileSystem string
    Total      uint64
    Free       uint64
}

type storageInfo struct {
    Name       string
    Size       uint64
    FreeSpace  uint64
    FileSystem string
}


var kernel = syscall.NewLazyDLL("Kernel32.dll")

func externalIP() (net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return nil, err
        }
        for _, addr := range addrs {
            ip := getIpFromAddr(addr)
            if ip == nil {
                continue
            }
            return ip, nil
        }
    }
    return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
    var ip net.IP
    switch v := addr.(type) {
    case *net.IPNet:
        ip = v.IP
    case *net.IPAddr:
        ip = v.IP
    }
    if ip == nil || ip.IsLoopback() {
        return nil
    }
    ip = ip.To4()
    if ip == nil {
        return nil // not an ipv4 address
    }

    return ip
}

func getCPUInfo() {

    var cpuinfo []cpuInfo

    err := wmi.Query("Select * from Win32_Processor", &cpuinfo)
    if err != nil {
        return
    }
    fmt.Println(cpuinfo[0].Name)
}

func getOSInfo() {
    var operatingsystem []operatingSystem
    err := wmi.Query("Select * from Win32_OperatingSystem", &operatingsystem)
    if err != nil {
        return
    }
    fmt.Println(operatingsystem[0].Name)
}

func getMemoryInfo() {

    GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
    var memInfo memoryStatusEx
    memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
    mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
    if mem == 0 {
        return
    }
    fmt.Println("total=:",memInfo.ullTotalPhys)
    fmt.Printf("%.2fGB", float64(memInfo.ullTotalPhys)/float64(1024*1024*1024))
    fmt.Printf("%.2fGB", float64(memInfo.ullAvailPhys)/float64(1024*1024*1024))
}

func getStorageInfo() {
    var storageinfo []storageInfo
    var localStorages []Storage
    err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
    if err != nil {
        return
    }

    for _, storage := range storageinfo {
        info := Storage{
            Name:       storage.Name,
            FileSystem: storage.FileSystem,
            Total:      storage.Size,
            Free:       storage.FreeSpace,
        }
        localStorages = append(localStorages, info)
    }
    fmt.Printf("%.2fGB", float64(localStorages[0].Total)/float64(1024*1024*1024))
}


func main() {
    ip, err := externalIP()
    if err != nil {
        fmt.Println(err)
    }
    getCPUInfo()
    getOSInfo()
    getMemoryInfo()
    getStorageInfo()

    fmt.Println(ip.String())
}