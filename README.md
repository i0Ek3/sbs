# sbs

sbs is a forced type conversion wrapper between string and []byte in Go.

## Usage

```
package main 

import (
    sbs "github.com/i0Ek3/sbs" 
)

func main() {
    //...

    sbs.B2S(byte)
    sbs.S2B(str)

    //...
}
```

## Declaration

sbs's main source code from [here](https://mp.weixin.qq.com/s/REtrm292mlIwzaYtJrV7bw), I modified something and rewrite test file to reuse easily, thanks to original author.

