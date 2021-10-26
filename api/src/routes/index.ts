import { Router } from "express";
import Home from "./home";

const router = Router();

router.use("/", Home);

export default router;
