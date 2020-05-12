// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The simple command demonstrates a simple functionality which
// prompts the user for a GitHub username and lists all the public
// organization memberships of the specified username.
package main

import (
	"flag"
	"fmt"
	"github.com/ianusmagnus/gh-release/ghutil"
	"os"
)

func main() {

	username := flag.String("username", "", "github username")
	pat := flag.String("pat", "", "github personal access token")
	repo := flag.String("repo", "", "github repository")

	creator := ghutil.NewReleaseCreator(*username, *pat, *repo)

	err := creator.CreateRelease()
	if err != nil {
		fmt.Printf("Creating release failed: %+v", err)
		os.Exit(1)
	}
}
