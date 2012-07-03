/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | topnear test                                                         |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
 * +----------------------------------------------------------------------+
 * | Created: 2012-06-26 00:41                                            |
 * +----------------------------------------------------------------------+
 */

package top

import(
    "fmt"
    "runtime"
    "testing"
    "strings"
    "math"
    . "../index"
)

/* {{{ build_assert_func(mesg string, t *testing.T) (func(expr bool))  */

func build_assert_func(mesg string, t *testing.T) (func(expr bool)int) {
    var counter int = 0

    return func(expr bool) int {
        if expr {
            counter ++
            fmt.Print(".")
        } else {
            _, fn, lineno, _ := runtime.Caller(1)
            mesg = "{" + fn[strings.LastIndex(fn, "/")+1:] + ":" + fmt.Sprint(lineno) + "} " + mesg

            t.Fatalf("[%dth cases] %s\n", counter, mesg)
        }
        return counter
    }
}

/* }}} */

/* {{{ Test_union(t *testing.T)  */

func Test_union(t *testing.T) {
    assert := build_assert_func("union() failed", t)

    result := union([]uint32{1, 2, 1, 3}, []uint32{1, 5, 2})
    assert(len(result) == 4)
    assert(result[0] + result[1] + result[2] + result[3] == 11)
}

/* }}} */
/* {{{ Test_fetch_from_index_by_count(t *testing.T)  */

func Test_fetch_from_index_by_count(t *testing.T) {
    assert := build_assert_func("fetch_from_index_by_count() failed", t)
    index_slice := Poi_1d_slice_t {
        {float64(1.0), 0},
        {float64(2.0), 1},
        {float64(3.0), 2},
        {float64(4.0), 3},
        {float64(5.0), 4},
        {float64(6.0), 5},
        {float64(7.0), 6},
        {float64(8.0), 7},
    }

    result, retval := fetch_from_index_by_count(index_slice, 3.0, 2)
    assert(retval == 0)
    assert(len(result) == 3)
    assert(result[0] == 1)
    assert(result[1] == 2)
    assert(result[2] == 3)

    result, retval = fetch_from_index_by_count(index_slice, 3.0, 1)
    assert(retval == 0)
    assert(len(result) == 3)

    result, retval = fetch_from_index_by_count(index_slice, 3.0, 3)
    assert(retval == 0)
    assert(len(result) == 5)
    assert(result[0] == 0)
    assert(result[1] == 1)
    assert(result[2] == 2)
    assert(result[3] == 3)
    assert(result[4] == 4)

    result, retval = fetch_from_index_by_count(index_slice, 7.0, 1)
    assert(retval == 0)
    assert(len(result) == 3)
    assert(result[0] == 5)
    assert(result[1] == 6)
    assert(result[2] == 7)

    result, retval = fetch_from_index_by_count(index_slice, 7.0, 2)
    assert(retval == 0)
    assert(len(result) == 3)
    assert(result[0] == 5)
    assert(result[1] == 6)
    assert(result[2] == 7)

    result, retval = fetch_from_index_by_count(index_slice, 7.0, 4)
    assert(retval == 0)
    assert(len(result) == 4)
    assert(result[0] == 4)
    assert(result[1] == 5)
    assert(result[2] == 6)
    assert(result[3] == 7)

    result, retval = fetch_from_index_by_count(index_slice, 7.0, 100)
    assert(retval == 0)
    assert(len(result) == 8)
    assert(result[0] == 0)
    assert(result[1] == 1)
    assert(result[2] == 2)
    assert(result[3] == 3)
    assert(result[4] == 4)
    assert(result[5] == 5)
    assert(result[6] == 6)
    assert(result[7] == 7)

    result, retval = fetch_from_index_by_count(index_slice, -3.0, 2)
    assert(retval == 0)
    assert(len(result) == 1)
    assert(result[0] == 0)

    index_slice = Poi_1d_slice_t {
        {1.0, 0},
        {2.0, 1},
        {2.2, 2},
        {3.3, 3},
    }
    result, retval = fetch_from_index_by_count(index_slice, 3.0, 1)
    assert(retval == 0)
    assert(len(result) == 2)
    assert(result[0] == 2)
    assert(result[1] == 3)

    result, retval = fetch_from_index_by_count(index_slice, 1.0, 1)
    assert(retval == 0)
    assert(len(result) == 2)
    assert(result[0] == 0)
    assert(result[1] == 1)

    result, retval = fetch_from_index_by_count(index_slice, 3.3, 1)
    assert(retval == 0)
    assert(len(result) == 2)
    assert(result[0] == 2)
    assert(result[1] == 3)

    result, retval = fetch_from_index_by_count(index_slice, 3.5, 1)
    assert(retval == 0)
    assert(len(result) == 1)
    assert(result[0] == 3)
}

/* }}} */
/* {{{ Test_intersect(t *testing.T)  */

func Test_intersect(t *testing.T) {
    assert := build_assert_func("intersect() failed", t)

    intersection := intersect([]uint32{1, 2, 3, 6, 0, 9}, []uint32{0, 8, 7, 3, 1})
    assert(len(intersection) == 3)
    assert(intersection[0] == 0)
    assert(intersection[1] == 3)
    assert(intersection[2] == 1)
}

/* }}} */
/* {{{ Test_calc_distance(t *testing.T)  */

func Test_calc_distance(t *testing.T) {
    assert := build_assert_func("calc_distance() failed", t)

    assert(5.0 == calc_distance(0.0, 3, 4.0, 0))
    assert(13 == calc_distance(0, 5, 12, 0))
}

/* }}} */
/* {{{ Test_find_farthest_poi(t *testing.T)  */

func Test_find_farthest_poi(t *testing.T) {
    assert := build_assert_func("find_farthest_poi() failed", t)

    point, max_distance, cache_distances := find_farthest_poi(
        []POI_Item {
            {0, 4, 3}, {1, 5, 12}, {2, 9, 40}, {3, 7, 24},
        },
        0, 0, []uint32{0, 1, 2, 3})
    assert(point == POI_Item{2, 9, 40})
    assert(len(cache_distances) == 4)
    assert(max_distance == 41)
    assert(cache_distances[POI_Item{0, 4, 3}] == 5)
    assert(cache_distances[POI_Item{1, 5, 12}] == 13)
    assert(cache_distances[POI_Item{3, 7, 24}] == 25)
    assert(cache_distances[POI_Item{2, 9, 40}] == 41)

    point, max_distance, cache_distances = find_farthest_poi(
        []POI_Item {
            {0, 4, 3}, {1, 5, 12}, {2, 9, 40}, {3, 7, 24},
        },
        0, 0, []uint32{1, 0})
    assert(point == POI_Item{1, 5, 12})
    assert(max_distance == 13)
    assert(len(cache_distances) == 2)
    assert(cache_distances[POI_Item{0, 4, 3}] == 5)
    assert(cache_distances[POI_Item{1, 5, 12}] == 13)
}

/* }}} */
/* {{{ Test_fetch_from_index_by_range(t *testing.T)  */

func Test_fetch_from_index_by_range(t *testing.T) {
    assert := build_assert_func("fetch_from_index_by_range() failed", t)

    index_slice := Poi_1d_slice_t{
        {1.0, 0}, {2.0, 3}, {3.0, 7},
        {4.0, 2}, {5.0, 1}, {6.0, 8},
        {7.0, 4}, {8.0, 5}, {9.0, 6},}

    result := fetch_from_index_by_range(index_slice, 3, 1)
    assert(len(result) == 3)
    assert(result[0] == 3)
    assert(result[1] == 7)
    assert(result[2] == 2)

    result = fetch_from_index_by_range(index_slice, 2.5, 1)
    assert(len(result) == 2)
    assert(result[0] == 3)
    assert(result[1] == 7)

    result = fetch_from_index_by_range(index_slice, 3.0, 3)
    assert(len(result) == 6)
    assert(result[0] == 0)
    assert(result[1] == 3)
    assert(result[2] == 7)
    assert(result[3] == 2)
    assert(result[4] == 1)
    assert(result[5] == 8)

    result = fetch_from_index_by_range(index_slice, 6.0, 4)
    assert(len(result) == 8)
    assert(result[0] == 3)
    assert(result[1] == 7)
    assert(result[2] == 2)
    assert(result[3] == 1)
    assert(result[4] == 8)
    assert(result[5] == 4)
    assert(result[6] == 5)
    assert(result[7] == 6)

    result = fetch_from_index_by_range(index_slice, 9.0, 4)
    assert(len(result) == 5)
    assert(result[0] == 1)
    assert(result[1] == 8)
    assert(result[2] == 4)
    assert(result[3] == 5)
    assert(result[4] == 6)

    result = fetch_from_index_by_range(index_slice, 31, 1)
    assert(result == nil)

    result = fetch_from_index_by_range(index_slice, 10, 3)
    assert(len(result) == 3)
    assert(result[0] == 4)
    assert(result[1] == 5)
    assert(result[2] == 6)

    result = fetch_from_index_by_range(index_slice, 0, 1.5)
    assert(len(result) == 1)
    assert(result[0] == 0)
}

/* }}} */
/* {{{ Test_filter_by_distance(t *testing.T)  */

func Test_filter_by_distance(t *testing.T) {
    assert := build_assert_func("filter_by_distance() failed", t)
    guid_slice := []POI_Item {
        {0, 3, 5},
        {1, 4, 6},
        {2, 6, 4},
        {3, 2.5, 2.5},
        {4, 4-math.Sqrt(2), 4-math.Sqrt(2)},
        {5, 2.6, 2.6},
        {6, 2, 2},
        {7, 4, 2},
        {8, 2, 4},
    }
    var cache distanceTable_t = make(distanceTable_t)

    result := filter_by_distance(4, 4, guid_slice, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8}, 2, cache)
    assert(len(result) == 7)
    assert(len(cache) == 9)
    assert(result[0] == 0)
    assert(result[1] == 1)
    assert(result[2] == 2)
    assert(result[3] == 4)
    assert(result[4] == 5)
    assert(result[5] == 7)
    assert(result[6] == 8)

    for k, _ := range cache {
        cache[k] = 1000
    }
    result = filter_by_distance(4, 4, guid_slice, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8}, 2, cache)
    assert(len(result) == 0)
    assert(len(cache) == 9)

    cache = distanceTable_t{}
    result = filter_by_distance(4, 4, guid_slice, []uint32{0, 1, 3, 4}, 2, cache)
    assert(len(result) == 3)
    assert(len(cache) == 4)
    assert(result[0] == 0)
    assert(result[1] == 1)
    assert(result[2] == 4)

    cache = distanceTable_t{}
    guid_slice = append(guid_slice, POI_Item{9, 4, 4})
    result = filter_by_distance(4, 4, guid_slice, []uint32{0, 6, 9}, 2, cache)
    assert(len(result) == 1)
    assert(result[0] == 0)
    assert(len(cache) == 3)
    assert(cache[POI_Item{9, 4, 4}] == 0)
}

/* }}} */
/* {{{ TestScanNearPOI(t testing.T)  */

func TestScanNearPOI(t *testing.T) {
    assert := build_assert_func("ScanNearPOI() failed", t)

    poi_idx := new(POI_index)
    poi_idx.PoiXIdx = make(Poi_1d_slice_t, 6)
    poi_idx.PoiYIdx = make(Poi_1d_slice_t, 6)
    poi_idx.GuidArray = make([]POI_Item, 6)

    poi_idx.PoiXIdx[0].XY = 2
    poi_idx.PoiXIdx[0].ID = 5
    poi_idx.PoiXIdx[1].XY = 2.5
    poi_idx.PoiXIdx[1].ID = 2
    poi_idx.PoiXIdx[2].XY = 4-math.Sqrt(2)
    poi_idx.PoiXIdx[2].ID = 3
    poi_idx.PoiXIdx[3].XY = 2.6
    poi_idx.PoiXIdx[3].ID = 4
    poi_idx.PoiXIdx[4].XY = 3
    poi_idx.PoiXIdx[4].ID = 0
    poi_idx.PoiXIdx[5].XY = 4
    poi_idx.PoiXIdx[5].ID = 1

    poi_idx.PoiYIdx[0].XY = 2
    poi_idx.PoiYIdx[0].ID = 5
    poi_idx.PoiYIdx[1].XY = 2.5
    poi_idx.PoiYIdx[1].ID = 2
    poi_idx.PoiYIdx[2].XY = 4-math.Sqrt(2)
    poi_idx.PoiYIdx[2].ID = 3
    poi_idx.PoiYIdx[3].XY = 2.6
    poi_idx.PoiYIdx[3].ID = 4
    poi_idx.PoiYIdx[4].XY = 5
    poi_idx.PoiYIdx[4].ID = 0
    poi_idx.PoiYIdx[5].XY = 6
    poi_idx.PoiYIdx[5].ID = 1

    poi_idx.GuidArray[0].GUID = 0
    poi_idx.GuidArray[0].X = 3
    poi_idx.GuidArray[0].Y = 5
    poi_idx.GuidArray[1].GUID = 1
    poi_idx.GuidArray[1].X = 4
    poi_idx.GuidArray[1].Y = 6
    poi_idx.GuidArray[2].GUID = 2
    poi_idx.GuidArray[2].X = 2.5
    poi_idx.GuidArray[2].Y = 2.5
    poi_idx.GuidArray[3].GUID = 3
    poi_idx.GuidArray[3].X = 4-math.Sqrt(2)
    poi_idx.GuidArray[3].Y = 4-math.Sqrt(2)
    poi_idx.GuidArray[4].GUID = 4
    poi_idx.GuidArray[4].X = 2.6
    poi_idx.GuidArray[4].Y = 2.6
    poi_idx.GuidArray[5].GUID = 5
    poi_idx.GuidArray[5].X = 2
    poi_idx.GuidArray[5].Y = 2

    var cache distanceTable_t = make(distanceTable_t)

    result := ScanNearPOI(4, 4, poi_idx, 100, 2, cache)
    assert(len(result) == 4)
    assert(len(cache) == 6)
    assert(result[0] == 0)
    assert(result[1] == 4)
    assert(result[2] == 3)
    assert(result[3] == 1)

    poi_idx.PoiXIdx = append(poi_idx.PoiXIdx, Poi_1d_item_t{100.0, 6})
    poi_idx.PoiYIdx = append(poi_idx.PoiXIdx, Poi_1d_item_t{100.0, 6})
    poi_idx.GuidArray = append(poi_idx.GuidArray, POI_Item{6, 100.0, 100.0})
    result = ScanNearPOI(4, 4, poi_idx, 100, 2, cache)
    assert(len(result) == 4)
    assert(len(cache) == 6) /* NOTE that cache table is not increased */
    assert(result[0] == 0)
    assert(result[1] == 4)
    assert(result[2] == 3)
    assert(result[3] == 1)

    result = ScanNearPOI(4, 4, poi_idx, 2, 2, cache)
    assert(len(result) == 2)
    assert(len(cache) == 6)
    assert(result[0] == 0)
    assert(result[1] == 4)

    result = ScanNearPOI(4, 4, poi_idx, 1, 2, cache)
    assert(len(result) == 1)
    assert(len(cache) == 6)
    assert(result[0] == 0)

    result = ScanNearPOI(4, 4, poi_idx, 10, 1.5, cache)
    assert(len(result) == 1)
    assert(len(cache) == 6)
    assert(result[0] == 0)
}

/* }}} */
/* {{{ Test_sortby_distance(t *testing.T)  */

func Test_sortby_distance(t *testing.T) {
    assert := build_assert_func("sortby_distance", t)

    guid_slice := make([]POI_Item, 3)
    guid_slice[0].GUID = 0
    guid_slice[0].X = 1
    guid_slice[0].Y = 3
    guid_slice[1].GUID = 1
    guid_slice[1].X = 1
    guid_slice[1].Y = 2
    guid_slice[2].GUID = 2
    guid_slice[2].X = 1
    guid_slice[2].Y = 1

    ids := []uint32{0, 1, 2}
    var cache distanceTable_t = make(distanceTable_t)
    cache[ guid_slice[0] ] = 3
    cache[ guid_slice[1] ] = 2
    cache[ guid_slice[2] ] = 1

    sortby_distance(guid_slice, ids, cache)
    assert(len(ids) == 3)
    assert(ids[0] == 2)
    assert(ids[1] == 1)
    assert(ids[2] == 0)

    ids = []uint32{2, 0, 1}
    sortby_distance(guid_slice, ids, cache)
    assert(len(ids) == 3)
    assert(ids[0] == 2)
    assert(ids[1] == 1)
    assert(ids[2] == 0)

    ids = []uint32{1, 2}
    sortby_distance(guid_slice, ids, cache)
    assert(len(ids) == 2)
    assert(ids[0] == 2)
    assert(ids[1] == 1)

    /* exception test */
    delete(cache, guid_slice[1])
    sortby_distance(guid_slice, ids, cache)
    assert(len(ids) == 2)
    assert(ids[0] == 2)
    assert(ids[1] == 0)
}

/* }}} */
/* {{{ Test_translate_guid(t *testing.T)  */

func Test_translate_guid(t *testing.T) {
    assert := build_assert_func("translate_guid() failed", t)
    var cache distanceTable_t = make(distanceTable_t)

    guid_slice := make([]POI_Item, 3)
    guid_slice[0].GUID = 1111
    guid_slice[0].X = 1
    guid_slice[0].Y = 3
    guid_slice[1].GUID = 2222
    guid_slice[1].X = 1
    guid_slice[1].Y = 2
    guid_slice[2].GUID = 3333
    guid_slice[2].X = 1
    guid_slice[2].Y = 1

    result := translate_guid(guid_slice, []uint32{0, 2, 1}, 0, 0, cache)
    assert(len(result) == 3)
    assert(result[0].GUID == 1111)
    assert(result[1].GUID == 3333)
    assert(result[2].GUID == 2222)

    result = translate_guid(guid_slice, []uint32{0, 2, 1}, 1, 1, cache)
    assert(len(result) == 2)
    assert(result[0].GUID == 1111)
    assert(result[1].GUID == 2222)

    result = translate_guid(guid_slice, []uint32{0, 2}, 0, 0, cache)
    assert(len(result) == 2)
    assert(result[0].GUID == 1111)
    assert(result[1].GUID == 3333)

    result = translate_guid(guid_slice, []uint32{0, 2}, 1, 1, cache)
    assert(len(result) == 1)
    assert(result[0].GUID == 1111)

    result = translate_guid(guid_slice, []uint32{0, 2}, 1, 2, cache)
    assert(len(result) == 2)
    assert(result[0].GUID == 1111)
    assert(result[1].GUID == 3333)
}

/* }}} */
/* {{{ TestFetchNearPOI(t *testing.T)  */

func TestFetchNearPOI(t *testing.T) {
    assert := build_assert_func("FetchNearPOI() failed", t)

    poi_idx := new(POI_index)
    poi_idx.PoiXIdx = make(Poi_1d_slice_t, 6)
    poi_idx.PoiYIdx = make(Poi_1d_slice_t, 6)
    poi_idx.GuidArray = make([]POI_Item, 6)

    poi_idx.PoiXIdx[0].XY = 2
    poi_idx.PoiXIdx[0].ID = 5
    poi_idx.PoiXIdx[1].XY = 2.5
    poi_idx.PoiXIdx[1].ID = 2
    poi_idx.PoiXIdx[2].XY = 4-math.Sqrt(2)
    poi_idx.PoiXIdx[2].ID = 3
    poi_idx.PoiXIdx[3].XY = 2.6
    poi_idx.PoiXIdx[3].ID = 4
    poi_idx.PoiXIdx[4].XY = 3
    poi_idx.PoiXIdx[4].ID = 0
    poi_idx.PoiXIdx[5].XY = 4
    poi_idx.PoiXIdx[5].ID = 1

    poi_idx.PoiYIdx[0].XY = 2
    poi_idx.PoiYIdx[0].ID = 5
    poi_idx.PoiYIdx[1].XY = 2.5
    poi_idx.PoiYIdx[1].ID = 2
    poi_idx.PoiYIdx[2].XY = 4-math.Sqrt(2)
    poi_idx.PoiYIdx[2].ID = 3
    poi_idx.PoiYIdx[3].XY = 2.6
    poi_idx.PoiYIdx[3].ID = 4
    poi_idx.PoiYIdx[4].XY = 5
    poi_idx.PoiYIdx[4].ID = 0
    poi_idx.PoiYIdx[5].XY = 6
    poi_idx.PoiYIdx[5].ID = 1

    poi_idx.GuidArray[0].GUID = 0
    poi_idx.GuidArray[0].X = 3
    poi_idx.GuidArray[0].Y = 5
    poi_idx.GuidArray[1].GUID = 1
    poi_idx.GuidArray[1].X = 4
    poi_idx.GuidArray[1].Y = 6
    poi_idx.GuidArray[2].GUID = 2
    poi_idx.GuidArray[2].X = 2.5
    poi_idx.GuidArray[2].Y = 2.5
    poi_idx.GuidArray[3].GUID = 3
    poi_idx.GuidArray[3].X = 4-math.Sqrt(2)
    poi_idx.GuidArray[3].Y = 4-math.Sqrt(2)
    poi_idx.GuidArray[4].GUID = 4
    poi_idx.GuidArray[4].X = 2.6
    poi_idx.GuidArray[4].Y = 2.6
    poi_idx.GuidArray[5].GUID = 5
    poi_idx.GuidArray[5].X = 2
    poi_idx.GuidArray[5].Y = 2

    guid_slice, err := FetchNearPOI(poi_idx, 4, 4, 1)
    assert(err == 0)
    assert(len(guid_slice) == 1)
    assert(guid_slice[0].GUID == 0)

    guid_slice, err = FetchNearPOI(poi_idx, 4, 4, 2)
    assert(err == 0)
    assert(len(guid_slice) == 2)
    assert(guid_slice[0].GUID == 0)
    assert(guid_slice[1].GUID == 4)
}

/* }}} */
/* {{{ TestFetchNearPOI_boundary_cond(t *testing.T)  */

func TestFetchNearPOI_boundary_cond(t *testing.T) {
    assert := build_assert_func("FetchNearPOI() failed", t)

    poi_idx := new(POI_index)
    poi_idx.PoiXIdx = make(Poi_1d_slice_t, 4)
    poi_idx.PoiYIdx = make(Poi_1d_slice_t, 4)
    poi_idx.GuidArray = make([]POI_Item, 4)

    poi_idx.GuidArray[0].GUID = 0
    poi_idx.GuidArray[0].X = 0
    poi_idx.GuidArray[0].Y = 0
    poi_idx.GuidArray[1].GUID = 1
    poi_idx.GuidArray[1].X = -1
    poi_idx.GuidArray[1].Y = 10
    poi_idx.GuidArray[2].GUID = 2
    poi_idx.GuidArray[2].X = -2
    poi_idx.GuidArray[2].Y = 2
    poi_idx.GuidArray[3].GUID = 3
    poi_idx.GuidArray[3].X = 10
    poi_idx.GuidArray[3].Y = 1

    poi_idx.PoiXIdx[0].XY = -2
    poi_idx.PoiXIdx[0].ID = 2
    poi_idx.PoiXIdx[1].XY = -1
    poi_idx.PoiXIdx[1].ID = 1
    poi_idx.PoiXIdx[2].XY = 0
    poi_idx.PoiXIdx[2].ID = 0
    poi_idx.PoiXIdx[3].XY = 10
    poi_idx.PoiXIdx[3].ID = 3

    poi_idx.PoiYIdx[0].XY = 0
    poi_idx.PoiYIdx[0].ID = 0
    poi_idx.PoiYIdx[1].XY = 1
    poi_idx.PoiYIdx[1].ID = 3
    poi_idx.PoiYIdx[2].XY = 2
    poi_idx.PoiYIdx[2].ID = 2
    poi_idx.PoiYIdx[3].XY = 10
    poi_idx.PoiYIdx[3].ID = 1

    guid_slice, err := FetchNearPOI(poi_idx, 0, 0, 1)
    assert(err == 0)
    assert(len(guid_slice) == 1)
    assert(guid_slice[0].GUID == 2)

    poi_idx.GuidArray = append(poi_idx.GuidArray, POI_Item{4, 1, 2})
    poi_idx.PoiXIdx   = make(Poi_1d_slice_t, 5)
    poi_idx.PoiYIdx   = make(Poi_1d_slice_t, 5)

    poi_idx.PoiXIdx[0].XY = -2
    poi_idx.PoiXIdx[0].ID = 2
    poi_idx.PoiXIdx[1].XY = -1
    poi_idx.PoiXIdx[1].ID = 1
    poi_idx.PoiXIdx[2].XY = 0
    poi_idx.PoiXIdx[2].ID = 0
    poi_idx.PoiXIdx[3].XY = 1
    poi_idx.PoiXIdx[3].ID = 4
    poi_idx.PoiXIdx[4].XY = 10
    poi_idx.PoiXIdx[4].ID = 3

    poi_idx.PoiYIdx[0].XY = 0
    poi_idx.PoiYIdx[0].ID = 0
    poi_idx.PoiYIdx[1].XY = 1
    poi_idx.PoiYIdx[1].ID = 3
    poi_idx.PoiYIdx[2].XY = 2
    poi_idx.PoiYIdx[2].ID = 2
    poi_idx.PoiYIdx[3].XY = 2
    poi_idx.PoiYIdx[3].ID = 4
    poi_idx.PoiYIdx[4].XY = 10
    poi_idx.PoiYIdx[4].ID = 1

    guid_slice, err = FetchNearPOI(poi_idx, 0, 0, 1)
    assert(err == 0)
    assert(len(guid_slice) == 1)
    assert(guid_slice[0].GUID == 4)
}

/* }}} */
/* {{{ TestFetchNearPOI_NotEnoughPOI(t *testing.T)  */

func TestFetchNearPOI_NotEnoughPOI(t *testing.T) {
    assert := build_assert_func("FetchNearPOI() failed", t)

    poi_idx := new(POI_index)
    poi_idx.PoiXIdx = make(Poi_1d_slice_t, 5)
    poi_idx.PoiYIdx = make(Poi_1d_slice_t, 5)
    poi_idx.GuidArray = make([]POI_Item, 5)

    poi_idx.GuidArray[0].GUID = 0
    poi_idx.GuidArray[0].X = 0
    poi_idx.GuidArray[0].Y = 0
    poi_idx.GuidArray[1].GUID = 1
    poi_idx.GuidArray[1].X = 1
    poi_idx.GuidArray[1].Y = 1
    poi_idx.GuidArray[2].GUID = 2
    poi_idx.GuidArray[2].X = 2
    poi_idx.GuidArray[2].Y = 2
    poi_idx.GuidArray[3].GUID = 3
    poi_idx.GuidArray[3].X = 3
    poi_idx.GuidArray[3].Y = 3
    poi_idx.GuidArray[4].GUID = 4
    poi_idx.GuidArray[4].X = 4
    poi_idx.GuidArray[4].Y = 4

    poi_idx.PoiXIdx[0].XY = 0
    poi_idx.PoiXIdx[0].ID = 0
    poi_idx.PoiXIdx[1].XY = 1
    poi_idx.PoiXIdx[1].ID = 1
    poi_idx.PoiXIdx[2].XY = 2
    poi_idx.PoiXIdx[2].ID = 2
    poi_idx.PoiXIdx[3].XY = 3
    poi_idx.PoiXIdx[3].ID = 3
    poi_idx.PoiXIdx[4].XY = 4
    poi_idx.PoiXIdx[4].ID = 4

    poi_idx.PoiYIdx[0].XY = 0
    poi_idx.PoiYIdx[0].ID = 0
    poi_idx.PoiYIdx[1].XY = 1
    poi_idx.PoiYIdx[1].ID = 1
    poi_idx.PoiYIdx[2].XY = 2
    poi_idx.PoiYIdx[2].ID = 2
    poi_idx.PoiYIdx[3].XY = 3
    poi_idx.PoiYIdx[3].ID = 3
    poi_idx.PoiYIdx[4].XY = 4
    poi_idx.PoiYIdx[4].ID = 4

    guid_slice, err := FetchNearPOI(poi_idx, 1, 1, 4)
    assert(err == 0)
    assert(len(guid_slice) == 4)
    assert(guid_slice[0].GUID == 0)
    assert(guid_slice[1].GUID == 2)
    assert(guid_slice[2].GUID == 3)
    assert(guid_slice[3].GUID == 4)
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
