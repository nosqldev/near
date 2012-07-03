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
    "time"
)

func main() {
    file, err := os.Open("/tmp/poi.data")
    if err != nil { fmt.Println(err); os.Exit(1); }

    poi_idx, ret := index.LoadPOI(file)
    if ret != 0 { fmt.Println("LoadPOI() failed: ", ret); os.Exit(1); }

    fmt.Println(len(poi_idx.GuidArray), "have been loaded", poi_idx.PoiXIdx[0], poi_idx.PoiXIdx[len(poi_idx.PoiXIdx)-1], poi_idx.PoiYIdx[0], poi_idx.PoiYIdx[len(poi_idx.PoiYIdx)-1])

    var x, y float64
    var cnt int

    for {
        fmt.Scanf("%f %f %d", &x, &y, &cnt)
        start_time := time.Now()
        fmt.Printf("got request: [%f %f] %d\n", x, y, cnt)
        guid_slice, ret := top.FetchNearPOI(poi_idx, x, y, cnt)
        if ret != 0 {
            fmt.Printf("[error] ret = %d\n", ret)
            continue
        }
        fmt.Printf("--- amount of results: %d---\n", len(guid_slice))
        for _, guid := range guid_slice {
           /*fmt.Printf("guid = %x, x = %f, y = %f\n", guid, poi_idx.GuidArray[guid].X, poi_idx.GuidArray[guid].Y)*/
           fmt.Printf("guid = %x, distance = %f\n", guid.GUID, guid.Distance)
        }
        end_time := time.Now()
        elapsed_time := end_time.Sub(start_time)
        fmt.Println(elapsed_time, "\n")
    }
}

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
