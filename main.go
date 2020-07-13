// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The simple command demonstrates a simple functionality which
// prompts the user for a GitHub username and lists all the public
// organization memberships of the specified username.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ianusmagnus/gh-release/ghutil"
	"github.com/ianusmagnus/gh-release/ghutil/client"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	username := flag.String("username", "", "github username (Required)")
	pat := flag.String("pat", "", "github personal access token (Required)")
	repo := flag.String("repo", "", "github repository (Required)")
	name := flag.String("name", "", "name and tag of the release (Required)")
	verbose := flag.Bool("verbose", false, "verbose mode")

	flag.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	if *username == "" || *pat == "" || *repo == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Info("Creating new client")
	ctx := context.Background()

	githubClient := client.NewClient(ctx, *pat)

	creator := ghutil.NewReleaseCreator(githubClient, ctx, *username, *repo)

	err := creator.CreateNewRelease(*name)
	if err != nil {
		fmt.Printf("Creating release failed: %+v", err)
		os.Exit(1)
	}
}
