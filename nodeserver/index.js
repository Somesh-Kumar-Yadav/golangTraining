// Importing required modules
const express = require("express");

// Creating an instance of the app
const app = express();

app.use(express.json());

app.use(express.urlencoded({ extended: true }));

// Creating a simple route
app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.post("/post", (req, res) => {
  let myJson = req.body;
  res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
  let myJson = req.body;
  res.status(200).send(JSON.stringify(myJson));
});

app.get("/get", (req, res) => {
  res.status(200).send(JSON.stringify({ name: "somesh" }));
});

// Starting the server
app.listen(3000, () => {
  console.log("Server started on port 3000");
});
