package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/nembis/gator/internal/config"
	"github.com/nembis/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		slog.Error("failed to read config", "error", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		slog.Error("failed to open database connection", "error", err)
		os.Exit(0)
	}
	defer db.Close()

	s := &state{
		cfg: cfg,
		db:  database.New(db),
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlersListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("browse", handlerBrowse)

	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	if len(os.Args) < 2 {
		slog.Error("requires 2+ arguemnts")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err = cmds.run(s, cmd); err != nil {
		slog.Error("failed to run command", "command", cmd.name, "args", cmd.args, "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}
