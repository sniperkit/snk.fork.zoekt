/*
Sniperkit-Bot
- Date: 2018-08-12 12:11:26.372554071 +0200 CEST m=+0.045728207
- Status: analyzed
*/

// Zoekt: a fast text search engine, intended for use with source code. (Pronunciation: roughly as you would pronounce "zooked" in English)

// Package cmd of the Zoekt kit contains the following CLI and Webservices:
//
// Zoekt's clients:
//
// - zoekt - Searching - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt
// $ $GOPATH/bin/zoekt 'ngram f:READ'
//
// - zoekt-index - Indexing - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-index
// $ $GOPATH/bin/zoekt-index .
//
// - zoekt-git-index - Indexing local git repositories - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-git-index
// $ $GOPATH/bin/zoekt-git-index -branches master,stable-1.4 -prefix origin/ .
//
// - zoekt-repo-index - Indexing repo repositories - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-mirror-gitiles
// $ $GOPATH/bin/zoekt-mirror-gitiles -dest ~/repos/ https://gfiber.googlesource.com
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-repo-index
// $ zoekt-repo-index \
//   -name gfiber \
//   -base_url https://gfiber.googlesource.com/ \
//   -manifest_repo ~/repos/gfiber.googlesource.com/manifests.git \
//   -repo_cache ~/repos \
//   -manifest_rev_prefix=refs/heads/ --rev_prefix= \
//   master:default_unrestricted.xml
//
// - zoekt-mirror-gitiles - Mirror remote Gitiles - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-mirror-gitiles
// $ $GOPATH/bin/zoekt-mirror-gitiles -dest ~/repos/ --user avelino
// $ $GOPATH/bin/zoekt-mirror-gitiles -dest ~/repos/ --org segmentio
//
// - zoekt-mirror-gerrit - Mirror remote Gerrit - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-mirror-gerrit
// $ $GOPATH/bin/zoekt-mirror-gerrit -dest ~/repos/ ssh://<USERNAME>@gerrit.wikimedia.org:29418/mediawiki/core.git
//
// - zoekt-mirror-github - Mirror remote Github repositories - CLI
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-mirror-github
// $ $GOPATH/bin/zoekt-mirror-github -dest ~/repos/ https://github.com/google/go-github
// $ $GOPATH/bin/zoekt-mirror-github -dest ~/repos/ git://github.com/google/go-github.git
//
// Zoekt's webservices:
//
// - zoekt-indexserver - Indexing - SERVICE
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-indexserver
// $ $GOPATH/bin/zoekt-indexserver .
//
// - zoekt - Web Interface
// $ go install github.com/sniperkit/snk.fork.zoekt/cmd/zoekt-webserver
// $ $GOPATH/bin/zoekt-webserver -listen :6070
//
package cmd

// EOF
