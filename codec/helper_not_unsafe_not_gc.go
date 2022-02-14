// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

package codec

import "reflect"

// This files contains safe versions of the code where the unsafe versions are not supported
// in either gccgo or gollvm.
//
// - rvType:
//   reflect.toType is not supported in gccgo, gollvm.

func rvType(rv reflect.Value) reflect.Type {
	return rv.Type()
}
