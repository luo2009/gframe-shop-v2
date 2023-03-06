package main

import (
	_ "gframe-shop-v2/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gframe-shop-v2/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
