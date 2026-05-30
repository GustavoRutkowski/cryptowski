package cli

import (
	"flag"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski/v1"
)

func ParseVersionFlags() cryptowski.ICrypto {
	isV1 := flag.Bool("v1", false, "Use version 1")
	// isV2 := flag.Bool("v2", false, "Use version 2")
	// ...
	flag.Parse()
	
	// If you pass many flags, always get the latest version
	switch {
	case *isV1:
		return v1.Crypto{}
	default:
		latest := v1.Crypto{}
		return latest
	}
}
