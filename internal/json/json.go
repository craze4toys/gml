/*
 * GML - Go QML
 *
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 Roland Singer <roland.singer[at]desertbit.com>
 * Copyright (c) 2019 Sebastian Borchers <sebastian[at]desertbit.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
json-iterator (https://github.com/json-iterator/go) seems to be even faster
than easyjson (https://github.com/mailru/easyjson) and does not require code
generation, but is in fact a 100% drop-in replacement for encoding/json.
*/
package json

import (
	jsoniter "github.com/json-iterator/go"
)

func Marshal(v interface{}) (data []byte, err error) {
	/*if m, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(m)
	}*/

	return jsoniter.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) (err error) {
	/*if m, ok := v.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(data, m)
	}*/

	return jsoniter.Unmarshal(data, v)
}