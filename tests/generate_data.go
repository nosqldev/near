/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * |                                                                  |
 * +----------------------------------------------------------------------+
 * | Author: jingmi@gmail.com                                             |
 * +----------------------------------------------------------------------+
 * | Created: 2012-06-28 16:16                                            |
 * +----------------------------------------------------------------------+
 */

package main

import(
    "os"
    "encoding/binary"
    "fmt"
)

func main() {
    var guid uint64
    var x, y float64

    file, err := os.OpenFile("/tmp/poi.data", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
    if err != nil {fmt.Println(err); os.Exit(-1); }
    for i:=0; i<10; i++ {
        guid = uint64(i)
        x = float64(i)
        y = float64(i)
        fmt.Printf("guid = %d, x = %f, y = %f\n", guid, x, y)
        binary.Write(file, binary.LittleEndian, guid)
        binary.Write(file, binary.LittleEndian, x)
        binary.Write(file, binary.LittleEndian, y)
    }
}


/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
