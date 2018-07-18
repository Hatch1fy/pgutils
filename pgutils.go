package pgutils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Hatch1fy/errors"
)

// Dump will dump a postgres database
func Dump(cfg Config, w io.Writer) (err error) {
	errBuf := bytes.NewBuffer(nil)
	cmd := exec.Command("pg_dump",
		"-h", cfg.Host,
		"-p", strconv.Itoa(int(cfg.Port)),
		"-U", cfg.User,
		cfg.Database,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", cfg.Password))

	if cfg.SSL {
		cmd.Env = append(cmd.Env, "PGSSLMODE=allow")
	}

	cmd.Stdout = w
	cmd.Stderr = errBuf

	if err = cmd.Run(); err != nil {
		return errors.Error(errBuf.String())
	}

	if err = os.Setenv("PGPASSWORD", ""); err != nil {
		return
	}

	return
}

// Import will import a PGDump file into the specified database
func Import(database, filename string) (err error) {
	cmd := exec.Command("psql", "-d", database, "-f", filename)
	errBuf := bytes.NewBuffer(nil)
	cmd.Stderr = errBuf
	if err = cmd.Run(); err != nil {
		return errors.Error(errBuf.String())
	}

	return
}

// ReplaceArgs will replace arguments for queries
// This is quite useful when preparing statements with replacement variables
// without having to rely on an ORM
func ReplaceArgs(query string, args ...string) (out string) {
	for i, arg := range args {
		search := fmt.Sprintf("$%d", i+1)
		query = strings.Replace(query, search, arg, -1)
	}

	out = query
	return
}
