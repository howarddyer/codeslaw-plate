package config

type Port struct {
    Live, Test string
}

var Ports = Port {
    Live: ":80",
    Test: ":8000",
}
