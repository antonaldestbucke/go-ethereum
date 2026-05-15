// Copyright 2014 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// geth is the official command-line client for Ethereum.
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/internal/debug"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/internal/flags"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/node"
	"go.uber.org/automaxprocs/maxprocs"

	// Force-load the tracer engines to trigger registration
	_ "github.com/ethereum/go-ethereum/eth/tracers/js"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"

	"github.com/urfave/cli/v2"
)

const (
	clientIdentifier = "geth" // Client identifier to advertise over the network
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	gitDate   = ""

	// The app that holds all commands and flags.
	app = flags.NewApp(gitCommit, gitDate, "the go-ethereum command line interface")

	// flags that configure the node
	// NOTE: BuilderFlags and TxPoolFlags are included here for builder/mempool tuning.
	nodeFlags = flags.Merge(
		utils.NodeFlags,
		utils.RpcFlags,
		utils.ConsoleFlags,
		debug.Flags,
		utils.EthFlags,
		utils.AccountFlags,
		utils.NetworkFlags,
		utils.TxPoolFlags,
		utils.BuilderFlags,
		utils.MinerFlags,
		utils.GpoFlags,
		utils.MetricsFlags,
	)
)

func init() {
	// Initialize the CLI app and start Geth
	app.Action = geth
	app.Commands = []*cli.Command{
		// See chaincmd.go:
		initCommand,
		importCommand,
		exportCommand,
		importPreimagesCommand,
		removedbCommand,
		dumpCommand,
		dumpGenesisCommand,
		// See accountcmd.go:
		accountCommand,
		// See consolecmd.go:
		consoleCommand,
		attachCommand,
		javascriptCommand,
		// See misccmd.go:
		versionCommand,
		versionCheckCommand,
		licenseCommand,
		// See config.go:
		dumpConfigCommand,
		// See cmd/utils/flags.go:
		utils.ShowDeprecated,
		// See snapshot.go:
		snapshotCommand,
		// See verkle.go
		vverkleCommand,
		// Se
