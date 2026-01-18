// ===========================
// Configuration & Constants
// ===========================
const API_BASE_URL = 'http://localhost:8080/api';

// DOM Elements
const elements = {
    loading: document.getElementById('loading'),
    errorMessage: document.getElementById('error-message'),
    errorText: document.getElementById('error-text'),
    carsSection: document.getElementById('cars-section'),
    manufacturersSection: document.getElementById('manufacturers-section'),
    carsGrid: document.getElementById('cars-grid'),
    manufacturersGrid: document.getElementById('manufacturers-grid'),
    modal: document.getElementById('car-modal'),
    navButtons: document.querySelectorAll('.nav-btn')
};

// ===========================
// State Management
// ===========================
let currentView = 'cars';
let carsData = [];
let manufacturersData = [];

// ===========================
// API Functions
// ===========================

/**
 * Fetch all cars from the backend
 */
async function fetchCars() {
    try {
        showLoading(true);
        hideError();

        const response = await fetch(`${API_BASE_URL}/cars`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        carsData = data;
        renderCars(data);
        
        showLoading(false);
    } catch (error) {
        console.error('Error fetching cars:', error);
        showError('Failed to load cars. Please check if the backend server is running.');
        showLoading(false);
    }
}

/**
 * Fetch specific car details by ID
 * This demonstrates the asynchronous goroutine/channel processing on the backend
 */
async function fetchCarDetails(carId) {
    try {
        showLoading(true);
        
        const response = await fetch(`${API_BASE_URL}/cars/${carId}`);
        
        if (!response.ok) {
            if (response.status === 404) {
                throw new Error('Car not found');
            }
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const carData = await response.json();
        
        showLoading(false);
        return carData;
    } catch (error) {
        console.error('Error fetching car details:', error);
        showError(`Failed to load car details: ${error.message}`);
        showLoading(false);
        throw error;
    }
}

/**
 * Fetch all manufacturers from the backend
 */
async function fetchManufacturers() {
    try {
        showLoading(true);
        hideError();

        const response = await fetch(`${API_BASE_URL}/manufacturers`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        manufacturersData = data;
        renderManufacturers(data);
        
        showLoading(false);
    } catch (error) {
        console.error('Error fetching manufacturers:', error);
        showError('Failed to load manufacturers. Please check if the backend server is running.');
        showLoading(false);
    }
}

// ===========================
// Render Functions
// ===========================

/**
 * Render cars grid
 */
function renderCars(cars) {
    if (!cars || cars.length === 0) {
        elements.carsGrid.innerHTML = '<p style="text-align: center; color: #6b7280;">No cars available.</p>';
        return;
    }

    elements.carsGrid.innerHTML = cars.map(car => `
        <div class="car-card" onclick="handleCarClick(${car.id})" style="animation-delay: ${Math.random() * 0.2}s">
            <img src="${car.imageUrl}" alt="${car.name}" class="car-image" onerror="this.src='https://via.placeholder.com/400x200?text=Car+Image'">
            <div class="car-info">
                <h3 class="car-name">${car.name}</h3>
                <p class="car-manufacturer">${car.manufacturer}</p>
                <div class="car-details">
                    <span class="car-category">${car.category}</span>
                    <span class="car-year">${car.year}</span>
                </div>
                <div class="car-details">
                    <span class="car-price">$${car.price.toLocaleString()}</span>
                    <button class="btn btn-primary" onclick="event.stopPropagation(); handleCarClick(${car.id})">
                        View Details
                    </button>
                </div>
            </div>
        </div>
    `).join('');
}

/**
 * Render manufacturers grid
 */
function renderManufacturers(manufacturers) {
    if (!manufacturers || manufacturers.length === 0) {
        elements.manufacturersGrid.innerHTML = '<p style="text-align: center; color: #6b7280;">No manufacturers available.</p>';
        return;
    }

    elements.manufacturersGrid.innerHTML = manufacturers.map((manufacturer, index) => `
        <div class="manufacturer-card" style="animation-delay: ${index * 0.1}s">
            <div class="manufacturer-logo">${manufacturer.logo}</div>
            <h3 class="manufacturer-name">${manufacturer.name}</h3>
            <div class="manufacturer-info">
                <p class="manufacturer-country">${manufacturer.country}</p>
                <p>Founded: ${manufacturer.foundingYear}</p>
            </div>
        </div>
    `).join('');
}

// ===========================
// Modal Functions
// ===========================

/**
 * Handle car card click - fetches detailed data and shows modal
 * This triggers the backend's goroutine/channel processing
 */
async function handleCarClick(carId) {
    try {
        // Fetch detailed car data from backend
        // This will trigger the goroutine/channel processing on the server
        const carData = await fetchCarDetails(carId);
        
        // Populate modal with car data
        document.getElementById('modal-car-name').textContent = carData.name;
        document.getElementById('modal-car-manufacturer').textContent = carData.manufacturer;
        document.getElementById('modal-car-category').textContent = carData.category;
        document.getElementById('modal-car-year').textContent = carData.year;
        document.getElementById('modal-car-price').textContent = `$${carData.price.toLocaleString()}`;
        document.getElementById('modal-car-image').src = carData.imageUrl;
        document.getElementById('modal-car-image').alt = carData.name;
        
        // Populate technical specifications from details object
        if (carData.details) {
            document.getElementById('modal-car-engine').textContent = carData.details.engine || 'N/A';
            document.getElementById('modal-car-horsepower').textContent = 
                carData.details.horsepower ? `${carData.details.horsepower} HP` : 'N/A';
            document.getElementById('modal-car-transmission').textContent = carData.details.transmission || 'N/A';
            document.getElementById('modal-car-drivetrain').textContent = carData.details.drivetrain || 'N/A';
        }
        
        // Show modal
        openModal();
    } catch (error) {
        console.error('Failed to show car details:', error);
    }
}

/**
 * Open the modal
 */
function openModal() {
    elements.modal.classList.remove('hidden');
    document.body.style.overflow = 'hidden'; // Prevent background scrolling
}

/**
 * Close the modal
 */
function closeModal() {
    elements.modal.classList.add('hidden');
    document.body.style.overflow = ''; // Restore scrolling
}

// Close modal when clicking outside of it
document.addEventListener('click', (e) => {
    if (e.target === elements.modal) {
        closeModal();
    }
});

// Close modal with Escape key
document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape' && !elements.modal.classList.contains('hidden')) {
        closeModal();
    }
});

// ===========================
// UI Helper Functions
// ===========================

/**
 * Show or hide loading indicator
 */
function showLoading(show) {
    if (show) {
        elements.loading.classList.remove('hidden');
    } else {
        elements.loading.classList.add('hidden');
    }
}

/**
 * Show error message
 */
function showError(message) {
    elements.errorText.textContent = message;
    elements.errorMessage.classList.remove('hidden');
}

/**
 * Hide error message
 */
function hideError() {
    elements.errorMessage.classList.add('hidden');
}

/**
 * Switch between different views (cars/manufacturers)
 */
function switchView(view) {
    currentView = view;
    
    // Update navigation buttons
    elements.navButtons.forEach(btn => {
        if (btn.dataset.view === view) {
            btn.classList.add('active');
        } else {
            btn.classList.remove('active');
        }
    });
    
    // Show/hide sections
    if (view === 'cars') {
        elements.carsSection.classList.remove('hidden');
        elements.manufacturersSection.classList.add('hidden');
        
        // Fetch cars if not already loaded
        if (carsData.length === 0) {
            fetchCars();
        }
    } else if (view === 'manufacturers') {
        elements.carsSection.classList.add('hidden');
        elements.manufacturersSection.classList.remove('hidden');
        
        // Fetch manufacturers if not already loaded
        if (manufacturersData.length === 0) {
            fetchManufacturers();
        }
    }
}

// ===========================
// Event Listeners
// ===========================

// Navigation button click handlers
elements.navButtons.forEach(btn => {
    btn.addEventListener('click', () => {
        switchView(btn.dataset.view);
    });
});

// ===========================
// Initialization
// ===========================

/**
 * Initialize the application when DOM is ready
 */
function init() {
    console.log('Cars Viewer Application Started');
    console.log('API Base URL:', API_BASE_URL);
    
    // Load initial view (cars)
    fetchCars();
    
    console.log('✅ Application initialized successfully');
}

// Start the application when DOM is fully loaded
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
} else {
    init();
}

// ===========================
// Export functions for HTML onclick handlers
// ===========================
window.handleCarClick = handleCarClick;
window.closeModal = closeModal;
