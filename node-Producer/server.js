const express = require("express");
const app = express();
const port = 3000;
const database = require("./pkg/database/database");
const router = require("./router/index");
let rabbitMQ = require("./broker/rabbitProducer");
const StartServer = async () => {
  app.use(express.json());
  await database();
  app.use(router);
  app.post("/createRabbit", (req, res) => {
    const { name, surname } = req.body;
    rabbitMQ("hello", { name, surname });
    res.json({message:"sent"})
  });
  app.listen(port, () => console.log(`Example app listening on port ${port}!`));
};
StartServer();
