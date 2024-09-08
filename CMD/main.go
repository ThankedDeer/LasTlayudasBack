package main

import (
	"github/thankeddeer/lastlayudas/cmd/provider"

	_ "github.com/lib/pq"
)

func main() {
	p := provider.NewProvider()

	err := p.Build().Run()
	if err != nil {
		return 
	}

}
