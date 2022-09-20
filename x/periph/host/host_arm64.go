// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package host

import (
	// Make sure CPU and board drivers are registered.
	_ "periph/x/periph/host/allwinner"
	_ "periph/x/periph/host/bcm283x"
	_ "periph/x/periph/host/pine64"
	_ "periph/x/periph/host/rpi"
	// for Talkkonnect on OrangePi Zero2 running Ubuntu 22.04 - Jammy server. This ubuntu version is 64bit OS comment above and uncomment below
        // _ "github.com/talkkonnect/periph/x/periph/host/allwinner"
        // _ "github.com/talkkonnect/periph/x/periph/host/bcm283x"
        // _ "github.com/talkkonnect/periph/x/periph/host/pine64"
        // _ "github.com/talkkonnect/periph/x/periph/host/rpi"	
)
