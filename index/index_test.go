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
)

/* {{{ TestVerifyIndex(t *testing.T) */

func TestVerifyIndex(t *testing.T) {
    var poi_idx1 POI_index = POI_index {
        {10.0, 20},
        {20.0, 1},
    }
    if !VerifyIndex(poi_idx1) {
        t.Fatal("VerifyIndex() failed")
    }

    var poi_idx2 POI_index = [2]POI_1d_index {
        {1, 2, 3},
        {2, 3},
    }
    if VerifyIndex(poi_idx2) {
        t.Fatal("VerifyIndex() failed")
    }
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
