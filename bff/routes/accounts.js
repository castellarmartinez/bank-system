const express = require("express");
require("dotenv").config();

const validateJWT = require("../middlewares/auth-middleware");

const ACCOUNT_SERVICE_URL = process.env.ACCOUNT_SERVICE_URL;

const router = express.Router();

router.post("/", validateJWT, async (req, res) => {
  try {
    const response = await fetch(`${ACCOUNT_SERVICE_URL}/accounts`, {
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

router.get("/:id", validateJWT, async (req, res) => {
  try {
    const response = await fetch(
      `${ACCOUNT_SERVICE_URL}/accounts/${req.params.id}`
    );

    if (!response.ok) throw new Error(`HTTP Error: ${response.status}`);

    res.json(await response.json());
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

module.exports = router;
