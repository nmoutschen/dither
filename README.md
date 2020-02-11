Dithering Threshold Map Generators
==================================

This library provide generators to create dithering threshold maps of various sizes.

## Usage

```go
import (
    "fmt"

    "github.com/nmoutschen/dither"
)

const size uint = 4

func myFunc() {
    //Ordered threshold map
    d := dither.NewOrdered(size)
    thr := d.Threshold(size*size/2)

    for i := uint(0); i < size; i++ {
        fmt.Println(thr[i*size:(i+1)*size])
    }
    //Output:
    //[1 0 1 0]
    //[0 1 0 1]
    //[1 0 1 0]
    //[0 1 0 1]

    //Random threshold map
    dr := dith.NewOrdered(size)
    thr := d.Threshold(size*size/2)

    for i := uint(0); i < size; i++ {
        fmt.Println(thr[i*size:(i+1)*size])
    }
    //Output
    //[0 1 0 0]
    //[1 0 0 1]
    //[0 1 1 0]
    //[1 0 1 1]
}
```