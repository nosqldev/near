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
    . "../index"
)

const max_near_poi_count = 10000

/* {{{ FetchNearPOI(poi_idx *POI_index, X float64, Y float64, count int) (guid_slice []uint64, retval int)  */

func FetchNearPOI(poi_idx *POI_index, X float64, Y float64, count int) (guid_slice []uint64, retval int) {
    retval = 0

    /*
     *if count > max_near_poi_count/2 { return nil, -100 }
     *result_x, retcode := fetch_from_index(poi_idx.PoiXIdx, X, count)
     *if retcode != 0 { return nil, -1 }
     *result_y, retcode := fetch_from_index(poi_idx.PoiYIdx, Y, count)
     *if retcode != 0 { return nil, -2 }
     */
    /*
     *result := intersect(result_x, result_y)
     *farthest_poi, cache_distances := find_farthest_poi(poi_idx.GuidArray, result)
     *result = scan_near_poi(poi_idx, farthest_poi, cache_distances)
     *sortby_distance(poi_idx.GuidArray, result)
     *guid_slice = translate_guid(poi_idx.GuidArray, result)
     */

    return
}

/* }}} */

/* {{{ fetch_from_index(poi_1d_slice Poi_1d_slice_t, v float64, count int) ([]uint32, int)  */

func fetch_from_index(poi_1d_slice Poi_1d_slice_t, v float64, count int) ([]uint32, int) {
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

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
