const express = require("express");
const expressStatic = require("express-static");
const data = require("./data.json");

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());

// Serve images
app.use("/api/images", expressStatic("img"));

const carModels = data.carModels;
const categories = data.categories; // Loaded from JSON
const manufacturers = data.manufacturers;

app.get("/api/models", (req, res) => res.json(carModels));

app.get("/api/models/:id", (req, res) => {
  const id = parseInt(req.params.id);
  const model = carModels.find((model) => model.id === id);
  if (!model) return res.status(404).json({ message: "Car model not found" });
  res.json(model);
});

// NEW: Add this endpoint!
app.get("/api/categories", (req, res) => res.json(categories));

app.get("/api/manufacturers", (req, res) => res.json(manufacturers));

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});