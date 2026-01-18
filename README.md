# 🚗 Cars Viewer

A full-stack web application for exploring, searching, and comparing vehicle specifications. Built with a **Go** backend, **Vanilla JavaScript** frontend, and a **Node.js** data service.

![Project Status](https://img.shields.io/badge/status-complete-green)
![Go Version](https://img.shields.io/badge/Go-1.21+-blue)
![License](https://img.shields.io/badge/license-MIT-lightgrey)

## 📖 Overview

Cars Viewer allows users to browse a catalog of vehicles, view detailed specifications (engine, transmission, etc.), and compare multiple models side-by-side. The project demonstrates the **Backend-for-Frontend (BFF)** pattern, where a Go server acts as a proxy and static file server for a client-side application, consuming data from an external mock API.

## ✨ Features

### Core Functionality
* **Dynamic Car Grid:** View all available car models with images and basic info.
* **Detailed Specs:** Click any car image or "View Details" to see comprehensive specifications (Horsepower, Drivetrain, Country of Origin).
* **Manufacturer Info:** Automatically enriches car details with manufacturer data (Founding Year, Country).

### 🚀 Advanced Features
* **Real-time Search:** Filter cars instantly by model name.
* **Smart Filtering:** Dropdown filter to browse by specific Manufacturers.
* **Side-by-Side Comparison:** Select up to 3 cars to compare their specs in a dedicated table view.
* **Async Architecture:** Uses Go **Goroutines and Channels** to handle API requests asynchronously.

---

## 🛠️ Tech Stack

### Backend (Go)
* **Language:** Go (Golang)
* **Server:** `net/http` (Standard Library)
* **Concurrency:** Goroutines & Channels for async data fetching.
* **Pattern:** Reverse Proxy (Forwards requests to the Node.js data service).

### Frontend (Client)
* **HTML5 & CSS3:** Responsive Grid Layout, Flexbox, and Modal system.
* **JavaScript (ES6+):** `fetch` API, DOM manipulation, Async/Await.
* **No Frameworks:** Pure Vanilla JS implementation.

### Data Source
* **Node.js:** Serves the raw JSON data and images via Express.

---

## 📂 Project Structure

```text
/viewer
├── /api                 # Node.js Data Service (Provided Resource)
│   ├── data.json        # Car & Manufacturer Data
│   └── /img             # Car Images
├── /cmd
│   └── /server
│       └── main.go      # Go Backend Entry Point
├── /frontend            # Static Frontend Assets
│   ├── index.html       # Main UI
│   ├── app.js           # Client logic (Fetch, Filter, Render)
│   └── style.css        # Styling
└── README.md            # Documentation

```

---

## ⚡ Getting Started

This project requires **Two Servers** running simultaneously: the Data API and the Web Server.

### Prerequisites

* [Go](https://go.dev/dl/) installed.
* [Node.js](https://nodejs.org/) installed.

### Step 1: Start the Data Service (Terminal 1)

This simulates the external API resource.

```bash
cd api
npm install   # Install dependencies (first time only)
npm start     # Runs on http://localhost:3000

```

### Step 2: Start the Go Backend (Terminal 2)

This serves the website and proxies API requests.

```bash
# From the root project folder
go run cmd/server/main.go

```

### Step 3: View the App

Open your browser and navigate to:
👉 **http://localhost:8080**

---

## 🧪 Testing the Requirements

1. **Mandatory Data Check:**
* Search for "Audi A4" -> Verify Engine (2.0L Inline-4) and HP (201).
* Search for "Mercedes-Benz" -> Verify Country (Germany) and Founding Year (1926).


2. **Concurrency Check:**
* The backend uses a `resultChan` (Channel) in `proxyDetailHandler` to fetch details asynchronously.


3. **Comparison Check:**
* Click "Compare" on 2-3 cars.
* Click "Compare Now" in the bottom bar to see the comparison table.



---

## 👥 Authors

* **Krishna Adhikari** - Backend & API Integration
* **Ayob Ali Flayih Al-Rukabi** - Frontend & UI/UX

---

*Kood/Sisu Project

```

```