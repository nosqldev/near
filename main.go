/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | main of near service                                                 |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
 * +----------------------------------------------------------------------+
 * | Created: 2012-06-23 11:59                                            |
 * +----------------------------------------------------------------------+
 */

package main

import(
    index "./index"
    top "./top"
    "os"
    "fmt"
)

func main() {
    file, err := os.Open("/tmp/poi.data")
    if err != nil { fmt.Println(err); os.Exit(1); }

    poi_idx, ret := index.LoadPOI(file)
    if ret != 0 { fmt.Println("LoadPOI() failed: ", ret); os.Exit(1); }

    fmt.Println(len(poi_idx.GuidArray), "have been loaded")

    var x, y float64
    var cnt int

    for {
        fmt.Scanf("%f %f %d", &x, &y, &cnt)
        fmt.Printf("got request: [%f %f] %d\n", x, y, cnt)
        guid_slice, ret := top.FetchNearPOI(poi_idx, x, y, cnt)
        if ret != 0 {
            fmt.Printf("[error] ret = %d\n", ret)
            continue
        }
        fmt.Printf("--- amount of results: %d---\n", len(guid_slice))
        for id, guid := range guid_slice {
            fmt.Printf("%d, %x\n", id, guid)
        }
    }
}

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
