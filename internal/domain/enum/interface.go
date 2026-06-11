package enum

type Enum interface {
    String() string
    Valid() bool
    Value() int16
}
