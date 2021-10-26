/*
	Web Server struct for the api
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

import express, { Express } from "express";
import { Server } from "http";
import Routes from "./routes";

export default class Web {
	public app: Express;
	public server: Server;

	constructor() {
		const app = express();
		const server = new Server(app);

		app.disable("x-powered-by");

		app.use("/", Routes);

		this.app = app;
		this.server = server;

		return this;
	}

	public listen(port: number): void {
		this.server.listen(port);
	}
}
