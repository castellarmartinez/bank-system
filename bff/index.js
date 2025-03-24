const express = require("express");
require("dotenv").config();

const authRoutes = require("./routes/auth");
const accountRoutes = require("./routes/accounts");
const transactionRoutes = require("./routes/transactions");

const PORT = process.env.PORT;

const app = express();
app.use(express.json());

app.get("/", (req, res) => res.send("BFF is running!"));

app.use("/auth", authRoutes);
app.use("/accounts", accountRoutes);
app.use("/transactions", transactionRoutes);

app.listen(PORT, () => console.log(`BFF running on port ${PORT}`));
