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
    "io"
    "encoding/binary"
    "sort"
)

const (
    default_size = 1000000
    max_near_poi_count = 10000
)

type poi_1d_item struct {
    XY float64 /* X or Y of a given POI */
    ID uint32  /* index referred to GuidArray belongs to POI_index */
}
type poi_1d_slice []poi_1d_item
type POI_Item struct {
    GUID uint64
    X float64
    Y float64
}
type POI_index struct {
    PoiXIdx []poi_1d_item
    PoiYIdx []poi_1d_item
    GuidArray []POI_Item
}

/* {{{ verifyIndex(poi_idx POI_index) bool  */

// this func is just designed for test case
func verifyIndex(poi_idx POI_index) bool {
    return ((len(poi_idx.PoiXIdx) == len(poi_idx.PoiYIdx)) &&
            (len(poi_idx.PoiXIdx) == len(poi_idx.GuidArray)));
}

/* }}} */
/* {{{ LoadPOI(reader io.Reader) (poi_idx *POI_index, retval int) */

func LoadPOI(reader io.Reader) (poi_idx *POI_index, retval int) {
    var tmp_X, tmp_Y float64
    var guid uint64
    var poi_item POI_Item
    var poi_x_item poi_1d_item
    var poi_y_item poi_1d_item
    retval = 0
    poi_idx = new(POI_index)
    poi_idx.PoiXIdx = make([]poi_1d_item, 0, default_size)
    poi_idx.PoiYIdx = make([]poi_1d_item, 0, default_size)
    poi_idx.GuidArray = make([]POI_Item, 0, default_size)

    for i:=uint32(0); ; i++ {
        err := binary.Read(reader, binary.LittleEndian, &guid)
        if err != nil { break }
        err = binary.Read(reader, binary.LittleEndian, &tmp_X)
        if err != nil { retval = -1; break }
        err = binary.Read(reader, binary.LittleEndian, &tmp_Y)
        if err != nil { retval = -2; break }

        poi_item.GUID = guid
        poi_item.X = tmp_X
        poi_item.Y = tmp_Y
        poi_x_item.XY = tmp_X
        poi_x_item.ID = i
        poi_y_item.XY = tmp_Y
        poi_y_item.ID = i

        poi_idx.GuidArray = append(poi_idx.GuidArray, poi_item)
        poi_idx.PoiXIdx   = append(poi_idx.PoiXIdx, poi_x_item)
        poi_idx.PoiYIdx   = append(poi_idx.PoiYIdx, poi_y_item)
    }

    sortPoiIndex(poi_idx)

    return poi_idx, retval
}

/* }}} */
/* {{{ sortPoiIndex(poi_idx *POI_index)  */

func sortPoiIndex(poi_idx *POI_index) {
    sort.Sort(poi_1d_slice(poi_idx.PoiXIdx))
    sort.Sort(poi_1d_slice(poi_idx.PoiYIdx))
}

/* }}} */

/* {{{ (slice poi_1d_slice) Len () int */

func (slice poi_1d_slice) Len () int{
    return len(slice)
}

/* }}} */
/* {{{ (slice poi_1d_slice) Less(i,j int) bool */

func (slice poi_1d_slice) Less(i,j int) bool{
    if slice[i].XY < slice[j].XY { return true }
    return false
}

/* }}} */
/* {{{ (slice poi_1d_slice) Swap(i,j int)  */

func (slice poi_1d_slice) Swap(i,j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
