const express = require("express");
const jwt = require("jsonwebtoken");
require("dotenv").config();

const API_KEY = process.env.API_KEY;
const JWT_SECRET = process.env.JWT_SECRET;
const JWT_EXPIRATION = process.env.JWT_EXPIRATION;

const router = express.Router();

router.post("/token", (req, res) => {
  const { apiKey } = req.body;

  if (apiKey !== API_KEY) {
    return res.status(401).json({ error: "Invalid API key" });
  }

  const token = jwt.sign({}, JWT_SECRET, { expiresIn: JWT_EXPIRATION });

  res.json({ token });
});

module.exports = router;
