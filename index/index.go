/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | index of POI                                                         |
 * +----------------------------------------------------------------------+
 * | Author: jingmi@gmail.com                                             |
 * +----------------------------------------------------------------------+
 * | Created: 2012-06-23 12:23                                            |
 * +----------------------------------------------------------------------+
 */

package index

import (
    /*"io"*/
    /*"encoding/binary"*/
)

type poi_1d_item struct {
    XY float64 /* X or Y of a given POI */
    ID uint32
}
type POI_Item struct {
    GUID string
    X float64
    Y float64
}
type POI_index struct {
    PoiXIdx []poi_1d_item
    PoiYIdx []poi_1d_item
    GuiArray []POI_Item
}

/* {{{ verifyIndex(poi_idx POI_index) bool  */

func verifyIndex(poi_idx POI_index) bool {
    return ((len(poi_idx.PoiXIdx) == len(poi_idx.PoiYIdx)) &&
            (len(poi_idx.PoiXIdx) == len(poi_idx.GuiArray)));
}

/* }}} */
/* {{{ LoadPOI(reader io.Reader) (poi_idx *POI_index, retval int) */

/*
 *func LoadPOI(reader io.Reader) (poi_idx *POI_index, retval int) {
 *    var tmp_X, tmp_Y float64
 *    retval = 0
 *    poi_idx = new(POI_index)
 *    poi_idx[0] = make(poi_1d_index, 0, 100000)
 *    poi_idx[1] = make(poi_1d_index, 0, 100000)
 *
 *    for i:=0; ; i++ {
 *        err := binary.Read(reader, binary.LittleEndian, &tmp_X)
 *        if err != nil { break }
 *        err = binary.Read(reader, binary.LittleEndian, &tmp_Y)
 *        if err != nil { retval = -1; break }
 *        poi_idx[0] = append(poi_idx[0], tmp_X)
 *        poi_idx[1] = append(poi_idx[1], tmp_Y)
 *    }
 *
 *    return poi_idx, retval
 *}
 */

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
