package main

import (
	"context"

	"go.einride.tech/sage/sg"
)

type CmdBackstage sg.Namespace

func (CmdBackstage) Default(ctx context.Context) error {
	sg.Deps(ctx, CmdBackstage.GoModTidy)
	return nil
}

func (CmdBackstage) GoModTidy(ctx context.Context) error {
	sg.Logger(ctx).Println("tidying Go module files...")
	cmd := sg.Command(ctx, "go", "mod", "tidy", "-v")
	cmd.Dir = sg.FromGitRoot("cmd", "backstage")
	return cmd.Run()
}
