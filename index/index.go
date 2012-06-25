/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | index of POI                                                         |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
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
)

type Poi_1d_item_t struct {
    XY float64 /* X or Y of a given POI */
    ID uint32  /* index referred to GuidArray belongs to POI_index */
}
type Poi_1d_slice_t []Poi_1d_item_t
type POI_Item struct {
    GUID uint64
    X float64
    Y float64
}
type POI_index struct {
    PoiXIdx Poi_1d_slice_t
    PoiYIdx Poi_1d_slice_t
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
    var poi_x_item Poi_1d_item_t
    var poi_y_item Poi_1d_item_t
    retval = 0
    poi_idx = new(POI_index)
    poi_idx.PoiXIdx = make(Poi_1d_slice_t, 0, default_size)
    poi_idx.PoiYIdx = make(Poi_1d_slice_t, 0, default_size)
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
    sort.Sort(Poi_1d_slice_t(poi_idx.PoiXIdx))
    sort.Sort(Poi_1d_slice_t(poi_idx.PoiYIdx))
}

/* }}} */

/* {{{ (slice Poi_1d_slice_t) Len () int */

func (slice Poi_1d_slice_t) Len () int{
    return len(slice)
}

/* }}} */
/* {{{ (slice Poi_1d_slice_t) Less(i,j int) bool */

func (slice Poi_1d_slice_t) Less(i,j int) bool{
    if slice[i].XY < slice[j].XY { return true }
    return false
}

/* }}} */
/* {{{ (slice Poi_1d_slice_t) Swap(i,j int)  */

func (slice Poi_1d_slice_t) Swap(i,j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
