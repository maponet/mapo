/*
Copyright 2013 Petru Ciobanu, Francesco Paglia, Lorenzo Pierfederici

This file is part of Mapo.

Mapo is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

Mapo is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Mapo.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"mapo/log"
)

func main() {
	// parse flags

	// load config and setup application
	log.SetLevel(log.DEBUG)
	log.Info("Setting log level to DEBUG")

	log.Info("Starting application")

	// register with supervisor
	log.Info("Joining supervisor")

	// init db
	log.Info("Initializing db")

	// load addons
	log.Info("Loading addons")

	// register handlers
	log.Info("Registering handlers")

	// start server
	log.Info("Accepting requests")

	// inform supervisor that we are up

	// for each request
		// check authentication/authorization

		// extract request operation

		// extract request arguments

		// pass operation and arguments to api.router

			// find function mapped to operation

			// call function with arguments

		// return result to user


	// close on signal
	log.Info("Closing application")
}
