# Cars Viewer - Full Stack Web Application

A university project demonstrating a full-stack web application with a Go backend and vanilla JavaScript frontend.

## 📋 Project Overview

**Cars Viewer** is a web application that allows users to browse car models, view detailed specifications, and explore manufacturer information. The project showcases:
- RESTful API design
- Asynchronous processing with goroutines and channels
- Modern, responsive frontend design
- Clean separation of concerns

## 🛠️ Tech Stack

### Backend
- **Language**: Go (Golang)
- **Features**: Standard library HTTP server, goroutines, channels
- **Architecture**: RESTful API with concurrent request handling

### Frontend
- **HTML5**: Semantic markup
- **CSS3**: Modern, responsive design with CSS Grid and Flexbox
- **Vanilla JavaScript**: ES6+ with Fetch API for AJAX requests

## 📁 Project Structure

```
cars-viewer/
│
├── main.go                      # Backend server
│
└── frontend/
    ├── index.html               # Main HTML file
    ├── style.css                # Styling
    └── app.js                   # Frontend logic
```

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher installed
- A modern web browser (Chrome, Firefox, Safari, Edge)

### Installation & Running

1. **Navigate to the project directory**:
   ```bash
   cd viewer
   ```

2. **Run the Go server**:
   ```bash
   go run main.go
   ```

3. **Open your browser**:
   Navigate to: `http://localhost:8080/frontend/`

4. **The server will start on port 8080** with the following endpoints:
   - `GET /api/cars` - List all cars
   - `GET /api/cars/{id}` - Get specific car details (uses goroutines/channels)
   - `GET /api/manufacturers` - List all manufacturers

## 🎯 Features

### Backend Features
✅ RESTful API endpoints  
✅ Asynchronous request processing with goroutines and channels  
✅ Custom error handling (404, 500)  
✅ CORS support for frontend communication  
✅ Structured logging  
✅ Mock data with specific test cases  

### Frontend Features
✅ Responsive grid layout  
✅ Dynamic data loading with Fetch API  
✅ Modal popup for detailed car information  
✅ View switching (Cars/Manufacturers)  
✅ Loading states and error handling  
✅ Smooth animations and transitions  
✅ Mobile-friendly design  

## 📊 API Endpoints

### Get All Cars
```http
GET /api/cars
```
**Response**: Array of car objects with basic information

### Get Car Details (Asynchronous Processing)
```http
GET /api/cars/{id}
```
**Response**: Detailed car object including technical specifications
**Note**: This endpoint uses goroutines and channels for processing

### Get All Manufacturers
```http
GET /api/manufacturers
```
**Response**: Array of manufacturer objects

## 🔬 Test Data

### Required Test Cases (Mandatory)

#### Audi A4
Must include the following details:
```json
{
  "engine": "2.0L Inline-4",
  "horsepower": 201,
  "transmission": "7-speed Automatic",
  "drivetrain": "All-Wheel Drive"
}
```

#### Mercedes-Benz
Manufacturer data must include:
```json
{
  "name": "Mercedes-Benz",
  "country": "Germany",
  "foundingYear": 1926
}
```

### Sample Cars Included
1. Audi A4 (Sedan)
2. Mercedes-Benz E-Class (Luxury Sedan)
3. BMW 3 Series (Sport Sedan)
4. Audi Q5 (SUV)
5. Toyota Camry (Sedan)
6. BMW X5 (SUV)

### Sample Manufacturers
1. Audi (Germany, 1909)
2. Mercedes-Benz (Germany, 1926)
3. BMW (Germany, 1916)
4. Toyota (Japan, 1937)

## 🔄 Goroutine/Channel Implementation

The backend implements an **asynchronous event system** for handling car detail requests:

### How It Works:
1. A background goroutine (`carDetailProcessor`) runs continuously
2. When a request for `/api/cars/{id}` comes in, the handler creates a response channel
3. The request is sent to the processor through the `carRequestChan` channel
4. The processor finds the car, simulates processing time, and sends back the result
5. The handler waits for the response and returns it to the client

### Benefits:
- Decouples request handling from data processing
- Demonstrates Go's concurrency model
- Simulates real-world async operations (DB queries, API calls)
- Non-blocking architecture

See `GOROUTINE_EXPLANATION.md` for detailed technical explanation.

## 🎨 Frontend Implementation

### Key Components:

#### 1. Data Fetching
```javascript
async function fetchCarDetails(carId) {
    const response = await fetch(`${API_BASE_URL}/cars/${carId}`);
    const carData = await response.json();
    return carData;
}
```

#### 2. Event Handling
- Click events on car cards trigger detailed view
- Navigation buttons switch between views
- Modal overlay and escape key to close popups

#### 3. Dynamic Rendering
- Cars and manufacturers are rendered dynamically from API data
- Modal content is populated on-demand when user clicks "View Details"

#### 4. Error Handling
- Network errors are caught and displayed to users
- Loading indicators show during data fetching
- Graceful degradation for missing images

## 🧪 Testing

### Manual Testing Steps:

1. **Test Cars View**:
   - Open the app, verify all 6 cars are displayed
   - Click on "Audi A4" and verify the modal shows correct details
   - Check that engine is "2.0L Inline-4" and horsepower is 201

2. **Test Manufacturers View**:
   - Click "Manufacturers" navigation button
   - Verify all 4 manufacturers are displayed
   - Check that Mercedes-Benz shows "Germany" and "Founded: 1926"

3. **Test Modal**:
   - Click "View Details" on any car
   - Verify modal opens with complete information
   - Click outside modal or press Escape to close
   - Click the X button to close

4. **Test Error Handling**:
   - Stop the Go server
   - Refresh the page
   - Verify error message is displayed
   - Restart server and click "View Details" to verify recovery

5. **Test Responsive Design**:
   - Resize browser window
   - Verify layout adapts to different screen sizes
   - Test on mobile device or use browser dev tools

## 📝 Code Quality

### Backend Best Practices:
- Clear separation of concerns (handlers, data models, utilities)
- Proper error handling with custom error responses
- Structured logging for debugging
- Type-safe structs with JSON tags
- CORS handling for cross-origin requests

### Frontend Best Practices:
- Modular JavaScript with clear function separation
- Async/await for cleaner asynchronous code
- Event delegation where appropriate
- Semantic HTML5 elements
- CSS variables for maintainable styling
- Responsive design with mobile-first approach

## 🐛 Troubleshooting

### Issue: "Failed to load cars"
**Solution**: Ensure the Go server is running on port 8080

### Issue: CORS errors
**Solution**: The backend includes CORS headers. Ensure you're accessing via `http://localhost:8080/frontend/` not `file://`

### Issue: Modal doesn't close
**Solution**: Check browser console for JavaScript errors. Ensure `closeModal()` function is accessible globally.

### Issue: Images not loading
**Solution**: The app uses placeholder images from Unsplash. Check internet connection or replace with local images.

## 🎓 Learning Outcomes

This project demonstrates:
- Full-stack web development
- RESTful API design and consumption
- Go's concurrency model (goroutines and channels)
- Modern JavaScript (ES6+, Fetch API)
- Responsive web design
- Separation of concerns
- Error handling strategies

## 📚 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [MDN Web Docs - Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)
- [CSS Grid Guide](https://css-tricks.com/snippets/css/complete-guide-grid/)

## 👥 Team

- **Frontend Developer**: [Your Name]
- **Backend Developer**: [Teammate's Name]

## 📄 License

This project is created for educational purposes as part of a university assignment.

---

**Note**: This is a reference implementation. Make sure to understand how each component works rather than just copying the code. The goroutine/channel implementation is particularly important to understand for your technical review.
