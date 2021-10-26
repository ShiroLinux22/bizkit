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
