# sbs

sbs is a wrapper of forced type conversion between string and []byte in Go.

## Install 

`go get "github.com/i0Ek3/sbs"`

## Usage

```
package main 

import (
    "github.com/i0Ek3/sbs" 
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

