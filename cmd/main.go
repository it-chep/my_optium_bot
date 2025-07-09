package main

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal"
)

func main() {
	ctx := context.Background()

	internal.New(ctx).Run(ctx)
}
