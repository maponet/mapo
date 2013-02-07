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
	"github.com/maponet/utils/log"
	"github.com/maponet/utils/conf"

	"flag"
)

func main() {

	/*
	parse flags

	In some situation we will pass path to configuration file as a command line
	value. This meaning that for first off all we need to define and parse all flags.
	The only flag that we required on this step is only conf flag ... But we
	can't distribute code with same functionality along file or files.
	*/
	var logLevel = log.FlagLevel("log")
	var confFilePath = flag.String("conf", "./conf.ini", "set path to configuration file")
	flag.Parse()

	// load config and setup application
	err := conf.ParseConfigFile(*confFilePath)
	if err != nil {
		log.Error("%v", err)
		return
	}

	// setup configuration value passed as command line arguments
	if len(*logLevel) > 0 {
		conf.GlobalConfiguration.AddOption("default", "loglevel", *logLevel)
	}

	// setup application

	// set log level
	value, _ := conf.GlobalConfiguration.GetString("default", "loglevel")
	if err := log.SetLevelString(value); err != nil {
		log.Error("%v", err)
		return
	}

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
