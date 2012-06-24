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

type POI_1d_index []float64
type POI_index [2]POI_1d_index

/* {{{ VerifyIndex(poi_idx POI_index) bool  */

func VerifyIndex(poi_idx POI_index) bool {
    return (len(poi_idx[0]) == len(poi_idx[1]));
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
