package main


import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "sync"
    "time"
)

//获取目录dir下的文件大小
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
    defer wg.Done()
    for _, entry := range dirents(dir) {
        if entry.IsDir() {//目录
            wg.Add(1)
            subDir := filepath.Join(dir, entry.Name())
            go walkDir(subDir, wg, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

//sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

//读取目录dir下的文件信息
func dirents(dir string) []os.FileInfo {
    sema <- struct{}{}
    defer func() { <-sema }()
    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du: %v\n", err)
        return nil
    }
    return entries
}

//输出文件数量的大小
func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}

//提供-v 参数会显示程序进度信息
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
    flag.Parse()
    roots := flag.Args()//需要统计的目录
    if len(roots) == 0 {
        roots = []string{"."}
    }
    fileSizes := make(chan int64)
    var wg sync.WaitGroup
    for _, root := range roots {
        wg.Add(1)
        go walkDir("C:\\Users\\Administrator\\Desktop", &wg, fileSizes)
        fmt.Println(root)
    }
    go func() {
        wg.Wait() //等待goroutine结束
        close(fileSizes)
    }()
    var tick <-chan time.Time
    if *verbose {
        tick = time.Tick(100 * time.Millisecond) //输出时间间隔
    }
    var nfiles, nbytes int64
loop:
    for {
        select {
        case size, ok := <-fileSizes:
            if !ok {
                break loop
            }
            nfiles++
            nbytes += size
        case <-tick:
            printDiskUsage(nfiles, nbytes)
        }
    }
    printDiskUsage(nfiles, nbytes)
}