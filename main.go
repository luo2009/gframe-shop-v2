package main

import (
	_ "gframe-shop-v2/internal/logic"
	_ "gframe-shop-v2/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gframe-shop-v2/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
