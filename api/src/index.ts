/*
	Index file for the api
    Copyright (C) 2021 Jack C <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
import path from "path";
import dotenv from "dotenv";
import { init as sentryInit, Handlers, Integrations } from "@sentry/node";
import { Integrations as TracingIntergrations } from "@sentry/tracing";
import Web from "./web";

// Load dotenv
dotenv.config({
	path: path.resolve(__dirname, "../../.env"),
});
if (!process.env.PORT) {
	console.error("PORT not set, defaulting to 3000");
}

const production = process.env.NODE_ENV == "production";

// Create Web Server
const app = new Web();

if (production && process.env.SENTRY_DSN != "") {
	sentryInit({
		dsn: process.env.SENTRY_DSN,
		environment: "api",

		integrations: [
			new Integrations.Http({ tracing: true }),
			new TracingIntergrations.Express({
				app: app.app,
			}),
		],
	});

	app.app.use(Handlers.requestHandler());
	app.app.use(Handlers.tracingHandler());
	app.app.use(Handlers.errorHandler());
} else
	console.log(
		"NODE_ENV is not production or SENTRY_DSN is undefined, skipping sentry",
	);

// Listen on Web Server
const port: number = parseInt(process.env.PORT || "3000");
app.listen(port);
console.log(`Listening on port ${port}`);
