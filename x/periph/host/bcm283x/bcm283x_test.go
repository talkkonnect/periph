// Copyright 2017 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package bcm283x

import "periph/x/periph/host/fs"

func init() {
	fs.Inhibit()
}
