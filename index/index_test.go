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
    /*"bytes"*/
    /*"encoding/binary"*/
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
            {"123", 100.0, 3.0},
            {"789", 200.0, 3.14159265},
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
            {"123", 100.0, 3.0},
            {"789", 200.0, 3.14159265},
            {"abc", 1, 1},
        },
    }
    if verifyIndex(poi_idx2) {
        t.Fatal("VerifyIndex() failed")
    }
}

/* }}} */
/* {{{ TestLoadPOI(t *testing.T)  */

/*
 *func TestLoadPOI(t *testing.T) {
 *    data := make([]byte, 0, 1000)
 *    buf  := bytes.NewBuffer(data)
 *
 *    binary.Write(buf, binary.LittleEndian, float64(1.0))
 *    binary.Write(buf, binary.LittleEndian, float64(200.0))
 *    binary.Write(buf, binary.LittleEndian, float64(19.0))
 *    binary.Write(buf, binary.LittleEndian, float64(3.14159265))
 *
 *    poi_idx, err := LoadPOI(buf)
 *    if poi_idx[0][0] != 1.0 { t.Fatal("LoadPOI() failed") }
 *    if poi_idx[1][0] != 200.0 { t.Fatal("LoadPOI() failed") }
 *    if poi_idx[0][1] != 19.0 { t.Fatal("LoadPOI() failed") }
 *    if poi_idx[1][1] != 3.14159265 { t.Fatal("LoadPOI() failed") }
 *    if err != 0 { t.Fatal("LoadPOI() failed") }
 *}
 */

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
