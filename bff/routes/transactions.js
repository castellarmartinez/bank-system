const express = require("express");
require("dotenv").config();

const validateJWT = require("../middlewares/auth-middleware");

const TRANSACTION_SERVICE_PORT = process.env.TRANSACTION_SERVICE_PORT;
const HOST = process.env.HOST;

const router = express.Router();

router.post("/", validateJWT, async (req, res) => {
  try {
    const response = await fetch(`http://${HOST}:${TRANSACTION_SERVICE_PORT}/transactions`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(req.body),
    });

    if (!response.ok) throw new Error(`HTTP Error: ${response.status}`);

    res.json(await response.json());
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

router.get("/:account_id", validateJWT, async (req, res) => {
  try {
    const response = await fetch(
      `http://${HOST}:${TRANSACTION_SERVICE_PORT}/transactions/${req.params.account_id}`
    );

    if (!response.ok) throw new Error(`HTTP Error: ${response.status}`);

    res.json(await response.json());
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

module.exports = router;
