/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | test index package                                                   |
 * +----------------------------------------------------------------------+
 * | Author: jingmi@gmail.com                                             |
 * +----------------------------------------------------------------------+
 * | Created: 2012-06-24 11:45                                            |
 * +----------------------------------------------------------------------+
 */

package index

import(
    "testing"
    "bytes"
    "encoding/binary"
)

/* {{{ Test_verifyIndex(t *testing.T) */

func Test_verifyIndex(t *testing.T) {
    var poi_idx1 POI_index = POI_index {
        []poi_1d_item {
            {100.0, 0},
            {200.0, 1},
        },
        []poi_1d_item {
            {3.0, 0},
            {3.14159265, 1},
        },
        []POI_Item {
            {123, 100.0, 3.0},
            {789, 200.0, 3.14159265},
        },
    }
    if !verifyIndex(poi_idx1) {
        t.Fatal("VerifyIndex() failed")
    }

    var poi_idx2 POI_index = POI_index {
        []poi_1d_item {
            {100.0, 0},
            {200.0, 1},
        },
        []poi_1d_item {
            {3.0, 0},
            {3.14159265, 1},
        },
        []POI_Item {
            {123, 100.0, 3.0},
            {789, 200.0, 3.14159265},
            {111, 1, 1},
        },
    }
    if verifyIndex(poi_idx2) {
        t.Fatal("VerifyIndex() failed")
    }
}

/* }}} */
/* {{{ TestLoadPOI(t *testing.T)  */

func TestLoadPOI(t *testing.T) {
    data := make([]byte, 0, 1000)
    buf  := bytes.NewBuffer(data)

    binary.Write(buf, binary.LittleEndian, uint64(1))
    binary.Write(buf, binary.LittleEndian, float64(1.0))
    binary.Write(buf, binary.LittleEndian, float64(200.0))

    binary.Write(buf, binary.LittleEndian, uint64(2))
    binary.Write(buf, binary.LittleEndian, float64(19.0))
    binary.Write(buf, binary.LittleEndian, float64(3.14159265))

    poi_idx, err := LoadPOI(buf)

    if err != 0 { t.Fatal("LoadPOI() failed") }

    if poi_idx.PoiXIdx[0].XY != 1.0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.PoiXIdx[0].ID != 0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.PoiXIdx[1].XY != 19.0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.PoiXIdx[1].ID != 1 { t.Fatal("LoadPOI() failed") }

    if poi_idx.PoiYIdx[0].XY != 3.14159265 { t.Fatal("LoadPOI() failed") }
    if poi_idx.PoiYIdx[0].ID != 1 { t.Fatal("LoadPOI() failed") }
    if poi_idx.PoiYIdx[1].XY != 200.0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.PoiYIdx[1].ID != 0 { t.Fatal("LoadPOI() failed") }

    if poi_idx.GuidArray[0].GUID != 1 { t.Fatal("LoadPOI() failed") }
    if poi_idx.GuidArray[0].X != 1.0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.GuidArray[0].Y != 200.0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.GuidArray[1].GUID != 2 { t.Fatal("LoadPOI() failed") }
    if poi_idx.GuidArray[1].X != 19.0 { t.Fatal("LoadPOI() failed") }
    if poi_idx.GuidArray[1].Y != 3.14159265 { t.Fatal("LoadPOI() failed") }
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
