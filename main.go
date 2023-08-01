// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// vault write auth/approle/login role_id=${ROLE_ID} secret_id=${SECRET_ID}
	// use the returned token to fetch the actual secret

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hello, this is an App Role demo! My VAULT_ROLE_ID is %s\n", os.Getenv("VAULT_ROLE_ID"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
