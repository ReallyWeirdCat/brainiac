// Copyright (c) 2026 Nikolai Papin
//
// This file is part of Brainiac gamification and education platform
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package transport

import (
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/ReallyWeirdCat/brainiac/web"
)

// spaHandler serves the embedded Svelte build and falls back to index.html for
// unknown paths so client-side routing works on refresh/deep links.
func spaHandler() http.Handler {
	sub, err := fs.Sub(web.Dist, "dist")
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServerFS(sub)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if p == "." || p == "" {
			p = "index.html"
		}
		if _, err := fs.Stat(sub, p); err != nil {
			// Unknown path: let the SPA router handle it.
			r.URL.Path = "/"
			p = "index.html"
		} else if p != "index.html" {
			// Serve a concrete asset (js/css/...) through the file server.
			fileServer.ServeHTTP(w, r)
			return
		}
		// Serve index.html (root or fallback).
		r.URL.Path = "/"
		fileServer.ServeHTTP(w, r)
	})
}
