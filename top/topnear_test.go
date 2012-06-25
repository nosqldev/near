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

/* {{{ Test_fetch_from_index(t *testing.T)  */

func Test_fetch_from_index(t *testing.T) {
    assert := build_assert_func("fetch_from_index() failed", t)
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

    result, retval := fetch_from_index(index_slice, 3.0, 2)
    assert(retval == 0)
    assert(len(result) == 3)
    assert(result[0] == 1)
    assert(result[1] == 2)
    assert(result[2] == 3)

    result, retval = fetch_from_index(index_slice, 3.0, 1)
    assert(retval == 0)
    assert(len(result) == 3)

    result, retval = fetch_from_index(index_slice, 3.0, 3)
    assert(retval == 0)
    assert(len(result) == 5)
    assert(result[0] == 0)
    assert(result[1] == 1)
    assert(result[2] == 2)
    assert(result[3] == 3)
    assert(result[4] == 4)

    result, retval = fetch_from_index(index_slice, 7.0, 1)
    assert(retval == 0)
    assert(len(result) == 3)
    assert(result[0] == 5)
    assert(result[1] == 6)
    assert(result[2] == 7)

    result, retval = fetch_from_index(index_slice, 7.0, 2)
    assert(retval == 0)
    assert(len(result) == 3)
    assert(result[0] == 5)
    assert(result[1] == 6)
    assert(result[2] == 7)

    result, retval = fetch_from_index(index_slice, 7.0, 4)
    assert(retval == 0)
    assert(len(result) == 4)
    assert(result[0] == 4)
    assert(result[1] == 5)
    assert(result[2] == 6)
    assert(result[3] == 7)

    result, retval = fetch_from_index(index_slice, 7.0, 100)
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
}

/* }}} */

/* vim: set expandtab tabstop=4 shiftwidth=4 foldmethod=marker: */
