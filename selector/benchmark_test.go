// Copyright (c) 2011 Andy Balholm. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package selector

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func MustParseHTML(doc string) *html.Node {
	dom, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		panic(err)
	}
	return dom
}

var selector = MustCompile(`div.matched`)
var doc = `<!DOCTYPE html>
<html>
<body>
<div class="matched">
  <div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
  </div>
</div>
</body>
</html>
`
var dom = MustParseHTML(doc)

func BenchmarkMatchAll(b *testing.B) {
	var matches []*html.Node
	for i := 0; i < b.N; i++ {
		matches = selector.MatchAll(dom)
	}
	_ = matches
}