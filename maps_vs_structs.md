# Golang: Maps vs Structs
##Map
* All keys must be of the same type
* All values must be of the same type
* Keys are indexed - we can iter over them.
* Use to represent a collection of related properties
* Don't need to know all the keys at compile time
* Reference Type!

```
// Declare Method 1
var colors map[string]string

// Declare Method 2
colors := make(map[string]string)

// Declare Method 3
colors := map[string]string{
    "red": "#ff0000",
	"green": "#4bf745",
	"white": "#ffffff",
}
```

## Struct
* Values can be of different type
* Keys don't support indexing
* You need to know all the different fields at compile time.
* Use to represent a "thing" with a lot of different properties.
* Value Type!