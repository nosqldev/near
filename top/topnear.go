/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | find top nearest POI                                                 |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
 * +----------------------------------------------------------------------+
 * | Created: 2012-06-23 12:07                                            |
 * +----------------------------------------------------------------------+
 */

package top

import(
    "sort"
    "math"
    . "../index"
)

const max_near_poi_count = 10000

type distanceTable_t map[POI_Item]float64
type sort_poi_t struct {
    ID uint32
    Key POI_Item
    Distance float64
}
type sort_poi_slice_t []sort_poi_t

/* {{{ FetchNearPOI(poi_idx *POI_index, X float64, Y float64, count int) (guid_slice []uint64, retval int)  */

func FetchNearPOI(poi_idx *POI_index, x float64, y float64, count int) (guid_slice []uint64, retval int) {
    retval = 0

/*
 *    if count > max_near_poi_count/2 { return nil, -100 }
 *    result_x, retcode := fetch_from_index_by_count(poi_idx.PoiXIdx, x, count)
 *    if retcode != 0 { return nil, -1 }
 *    result_y, retcode := fetch_from_index_by_count(poi_idx.PoiYIdx, y, count)
 *    if retcode != 0 { return nil, -2 }
 *    result := intersect(result_x, result_y)
 *    farthest_point, max_distance, cache_distances := find_farthest_poi(poi_idx.GuidArray, x, y, result)
 *    result = ScanNearPOI(x, y, poi_idx, count, max_distance, cache_distances)
 *    guid_slice = translate_guid(poi_idx.GuidArray, result, x, y)
 *
 */
    return
}

/* }}} */
/* {{{ ScanNearPOI(x float64, y float64, poi_idx *POI_index, count int, max_distance float64, cache_distances distanceTable_t) []uint32  */

func ScanNearPOI(x float64, y float64, poi_idx *POI_index, count int, max_distance float64, cache_distances distanceTable_t) []uint32 {
    result_x := fetch_from_index_by_range(poi_idx.PoiXIdx, x, max_distance)
    if result_x == nil { return nil }
    result_y := fetch_from_index_by_range(poi_idx.PoiYIdx, y, max_distance)
    if result_y == nil { return nil }

    result := intersect(result_x, result_y)
    if result == nil { return nil }
    result = filter_by_distance(x, y, poi_idx.GuidArray, result, max_distance, cache_distances)

    var result_size int
    if count == 0 {
        result_size = len(result)
    } else {
        result_size = int(math.Min(float64(count), float64(len(result))))
    }
    sortby_distance(poi_idx.GuidArray, result, cache_distances)

    return result[:result_size]
}

/* }}} */

/* {{{ sortby_distance(guid_slice []POI_Item, ids []uint32, cache_distances distanceTable_t)  */

/* This function will drop those id which are not caculated in cache_distances!
 * Basically, the distances corresponding to the points to be sorted should be caculated before.
 */
func sortby_distance(guid_slice []POI_Item, ids []uint32, cache_distances distanceTable_t) {
    sort_slice := make(sort_poi_slice_t, 0, len(ids))
    var sort_item sort_poi_t
    for _, id := range ids {
        distance, exists := cache_distances[ guid_slice[id] ]
        if !exists { continue /* TODO Throw exception here */ }
        sort_item.ID = id
        sort_item.Key = guid_slice[id]
        sort_item.Distance = distance
        sort_slice = append(sort_slice, sort_item)
    }
    sort.Sort(sort_slice)

    for id, item := range sort_slice {
        ids[id] = item.ID
    }
    /* set the tailing ids slice to zero POI if querying non-existence item in cache_distances */
    for i := len(sort_slice); i < len(ids); i++ {
        ids[i] = 0
    }
}

/* }}} */
/* {{{ (s sort_poi_slice_t) Len() int  */

func (s sort_poi_slice_t) Len() int {
    return len(s)
}

/* }}} */
/* {{{ (s sort_poi_slice_t) Less(i,j int) bool */

func (s sort_poi_slice_t) Less(i,j int) bool{
    return s[i].Distance < s[j].Distance
}

/* }}} */
/* {{{ (s sort_poi_slice_t) Swap(i,j int)  */

func (s sort_poi_slice_t) Swap(i,j int) {
    s[i], s[j] = s[j], s[i]
}

/* }}} */
/* {{{ fetch_from_index_by_range(poi_1d_slice Poi_1d_slice_t, v float64, max_distance float64) []uint32 */

func fetch_from_index_by_range(poi_1d_slice Poi_1d_slice_t, v float64, max_distance float64) []uint32 {
    pos := sort.Search(len(poi_1d_slice), func(i int)bool { return poi_1d_slice[i].XY >= v })
    if pos >= len(poi_1d_slice) { pos = len(poi_1d_slice) - 1 }

    var start, end int = -1, -1
    for p := pos; p >= 0 && poi_1d_slice[p].XY >= v - max_distance; p-- {
        start = p
    }
    for p := pos; p < len(poi_1d_slice) && poi_1d_slice[p].XY <= v + max_distance; p++ {
        end = p
    }
    if start == -1 || end == -1 { return nil }
    end ++
    if end > start + max_near_poi_count { end = start + max_near_poi_count }

    result := make([]uint32, 0, max_near_poi_count)
    for i := start; i < end ; i++ {
        result = append(result, poi_1d_slice[i].ID)
    }

    if len(result) == 0 {
        return nil
    }
    return result
}

/* }}} */
/* {{{ fetch_from_index_by_count(poi_1d_slice Poi_1d_slice_t, v float64, count int) ([]uint32, int)  */

func fetch_from_index_by_count(poi_1d_slice Poi_1d_slice_t, v float64, count int) ([]uint32, int) {
    var retval int = 0

    pos := sort.Search(len(poi_1d_slice),
               func(i int)bool {
                   return poi_1d_slice[i].XY >= v
               })
    if pos >= len(poi_1d_slice) || poi_1d_slice[pos].XY != v {
        return nil, 0
    }

    var start, end int
    delta := int(float32(count) / 2 + 0.5)
    if pos - delta >= 0 {
        start = pos - delta
    } else {
        start = 0
    }
    if pos + delta >= len(poi_1d_slice) {
        end = len(poi_1d_slice)
    } else {
        end = pos + delta + 1
    }

    var result []uint32 = make([]uint32, 0, max_near_poi_count)
    for i := start; i < end; i++ {
        result = append(result, poi_1d_slice[i].ID)
    }

    return result, retval
}

/* }}} */
/* {{{ intersect(x_slice, y_slice []uint32) []uint32  */

func intersect(x_slice, y_slice []uint32) []uint32 {
    var result_len int = 0
    if len(x_slice) > len(y_slice) {
        result_len = len(x_slice)
    } else {
        result_len = len(y_slice)
    }
    var result []uint32 = make([]uint32, 0, result_len)
    var x_set map[uint32]bool = make(map[uint32]bool)

    for _, x := range x_slice {
        x_set[x] = true
    }
    for _, y := range y_slice {
        if _, err := x_set[y]; err {
            result = append(result, y)
        }
    }

    return result
}

/* }}} */
/* {{{ find_farthest_poi(guid_slice []POI_Item, x float64, y float64, ids []uint32) (point POI_Item, max_distance float64, cache_distances distanceTable_t) */

func find_farthest_poi(guid_slice []POI_Item, x float64, y float64, ids []uint32) (point POI_Item, max_distance float64, cache_distances distanceTable_t) {
    max_distance = 0
    cache_distances = make(distanceTable_t)

    for _, id := range ids {
        pnt := guid_slice[id]
        cache_distances[pnt] = calc_distance(pnt.X, pnt.Y, x, y)
        if max_distance < cache_distances[pnt] {
            max_distance = cache_distances[pnt]
            point = pnt
        }
    }

    return
}

/* }}} */
/* {{{ calc_distance(x1, y1, x2, y2 float64) distance float64  */

func calc_distance(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Abs(x1 - x2) * math.Abs(x1 - x2) + math.Abs(y1 - y2) * math.Abs(y1 - y2))
}

/* }}} */
/* {{{ filter_by_distance(x float64, y float64, guid_slice []POI_Item, ids []uint32, max_distance float64, cache_distances distanceTable_t) []uint32  */

func filter_by_distance(x float64, y float64, guid_slice []POI_Item, ids []uint32, max_distance float64, cache_distances distanceTable_t) []uint32 {
    var pnt POI_Item
    var result_ids []uint32 = make([]uint32, 0, max_near_poi_count)

    for _, id := range ids {
        if id > uint32(len(guid_slice)) { continue }
        pnt = guid_slice[id]
        distance, exists := cache_distances[pnt]
        if !exists {
            distance = calc_distance(x, y, pnt.X, pnt.Y)
            cache_distances[pnt] = distance
        }

        if distance <= max_distance {
            result_ids = append(result_ids, id)
        }
    }

    return result_ids
}

/* }}} */
/* {{{ translate_guid(guid_slice []POI_Item, ids []uint32, x float64, y float64) []uint64  */

func translate_guid(guid_slice []POI_Item, ids []uint32, x float64, y float64) []uint64 {
    var result []uint64 = make([]uint64, 0, len(ids))

    for _, id := range ids {
        if guid_slice[id].X == x && guid_slice[id].Y == y {
            /* filter the querying POI */
            continue
        }
        result = append(result, guid_slice[id].GUID)
    }

    return result
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
