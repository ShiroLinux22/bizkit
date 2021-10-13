import path from "path";
import amqp, { Connection } from "amqplib";
import dotenv from "dotenv";
import { fileURLToPath } from "url";
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
dotenv.config({
	path: path.resolve(__dirname, "../../.env"),
});
if (!process.env.AMQP_URI) throw new Error("AMQP_URI Variable not defined");

let queue = "events";
const conn = await amqp.connect(process.env.AMQP_URI);
const channel = await conn.createChannel();
channel.assertQueue(queue, {
	durable: false,
});

channel.consume(queue, (msg) => {
	console.log(msg?.content.toString());
});

console.log("waiting for events");
