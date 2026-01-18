// GLOBAL VARIABLES
let allCars = [];
let manufacturers = [];
let categories = []; // NEW: Store categories here
let compareList = new Set(); 

document.addEventListener("DOMContentLoaded", () => {
    fetchManufacturers();
    fetchCategories(); // NEW: Fetch them on load
    fetchCars();

    // Event Listeners
    document.getElementById('searchInput').addEventListener('input', filterCars);
    document.getElementById('manufacturerSelect').addEventListener('change', filterCars);
    document.getElementById('resetFilters').addEventListener('click', resetFilters);

    // Modal Close
    document.querySelector('.close-btn').addEventListener('click', () => {
        document.getElementById('carModal').style.display = 'none';
    });
    window.addEventListener('click', (e) => {
        if (e.target == document.getElementById('carModal')) {
            document.getElementById('carModal').style.display = 'none';
        }
        if (e.target == document.getElementById('compareModal')) {
            closeCompareModal();
        }
    });
});

// --- API FETCHING ---
async function fetchCars() {
    try {
        const response = await fetch('/api/cars');
        allCars = await response.json();
        renderCars(allCars);
    } catch (error) {
        console.error("Error loading cars", error);
    }
}

async function fetchManufacturers() {
    try {
        const response = await fetch('/api/manufacturers');
        manufacturers = await response.json();
        const select = document.getElementById('manufacturerSelect');
        select.innerHTML = '<option value="">All Manufacturers</option>';
        manufacturers.forEach(m => {
            const option = document.createElement('option');
            option.value = m.id;
            option.textContent = m.name;
            select.appendChild(option);
        });
    } catch (error) { console.error("Error loading manufacturers", error); }
}

// NEW: Fetch Categories Function
async function fetchCategories() {
    try {
        const response = await fetch('/api/categories');
        categories = await response.json();
    } catch (error) { console.error("Error loading categories", error); }
}

// --- FILTERING & RENDERING (Same as before) ---
function filterCars() {
    const searchText = document.getElementById('searchInput').value.toLowerCase();
    const manufacturerId = document.getElementById('manufacturerSelect').value;

    const filtered = allCars.filter(car => {
        const matchesSearch = car.name.toLowerCase().includes(searchText);
        const matchesManuf = manufacturerId === "" || car.manufacturerId == manufacturerId;
        return matchesSearch && matchesManuf;
    });
    renderCars(filtered);
}

function resetFilters() {
    document.getElementById('searchInput').value = '';
    document.getElementById('manufacturerSelect').value = '';
    renderCars(allCars);
}

function renderCars(carsToDisplay) {
    const grid = document.getElementById('car-grid');
    grid.innerHTML = '';

    if (carsToDisplay.length === 0) {
        grid.innerHTML = '<p>No cars found.</p>';
        return;
    }

    carsToDisplay.forEach(car => {
        const card = document.createElement('div');
        card.className = 'car-card';
        
        const imageSrc = car.image ? `img/${car.image}` : 'https://via.placeholder.com/300x200';
        const isSelected = compareList.has(car.id);

        card.innerHTML = `
            <img src="${imageSrc}" alt="${car.name}" class="car-image" onclick="openCarDetails(${car.id})">
            <div class="card-info">
                <h3>${car.name}</h3>
                <p><strong>Year:</strong> ${car.year}</p>
                <div class="action-buttons">
                    <button class="view-btn" onclick="openCarDetails(${car.id})">Details</button>
                    <button class="compare-select-btn ${isSelected ? 'selected' : ''}" 
                            onclick="toggleCompare(${car.id}, this)">
                        ${isSelected ? 'Selected' : 'Compare'}
                    </button>
                </div>
            </div>
        `;
        grid.appendChild(card);
    });
}

// --- DETAILS LOGIC (UPDATED) ---
async function openCarDetails(id) {
    const modal = document.getElementById('carModal');
    const modalBody = document.getElementById('modal-body');
    
    modalBody.innerHTML = '<p>Loading details...</p>';
    modal.style.display = 'flex';

    try {
        const response = await fetch(`/api/cars/${id}`);
        const car = await response.json();

        // LOOKUP LOGIC: Translate IDs to Names
        // 1. Manufacturer
        const manufObj = manufacturers.find(m => m.id == car.manufacturerId);
        const manufName = manufObj ? manufObj.name : 'Unknown';
        const manufCountry = manufObj ? manufObj.country : 'Unknown';
        const manufYear = manufObj ? manufObj.foundingYear : 'Unknown';

        // 2. Category
        const catObj = categories.find(c => c.id == car.categoryId);
        const categoryName = catObj ? catObj.name : 'Unknown';

        const specs = car.specifications || {};
        const imageSrc = car.image ? `img/${car.image}` : 'https://via.placeholder.com/400';
        
        modalBody.innerHTML = `
            <img src="${imageSrc}" class="modal-image">
            <h2>${car.name}</h2>
            <p class="subtitle"><strong>Category:</strong> ${categoryName}</p>
            <hr>
            
            <div class="details-section">
                <h3>Specifications</h3>
                <div class="specs-grid">
                    <p><strong>Engine:</strong> ${specs.engine || 'N/A'}</p>
                    <p><strong>Horsepower:</strong> ${specs.horsepower || 'N/A'} hp</p>
                    <p><strong>Transmission:</strong> ${specs.transmission || 'N/A'}</p>
                    <p><strong>Drivetrain:</strong> ${specs.drivetrain || 'N/A'}</p>
                    <p><strong>Year:</strong> ${car.year}</p>
                </div>
            </div>

            <div class="details-section" style="margin-top: 20px; background: #f9f9f9; padding: 15px; border-radius: 5px;">
                <h3>Manufacturer Details</h3>
                <p><strong>Name:</strong> ${manufName}</p>
                <p><strong>Country:</strong> ${manufCountry}</p>
                <p><strong>Founding Year:</strong> ${manufYear}</p>
            </div>
        `;
    } catch (error) {
        console.error(error);
        modalBody.innerHTML = '<p style="color:red">Error loading details.</p>';
    }
}

// --- COMPARISON LOGIC (Keep as is) ---
function toggleCompare(id, btnElement) {
    if (compareList.has(id)) {
        compareList.delete(id);
        btnElement.classList.remove('selected');
        btnElement.innerText = "Compare";
    } else {
        if (compareList.size >= 3) {
            alert("Max 3 cars.");
            return;
        }
        compareList.add(id);
        btnElement.classList.add('selected');
        btnElement.innerText = "Selected";
    }
    updateCompareBar();
}

function updateCompareBar() {
    const bar = document.getElementById('compareBar');
    document.getElementById('compareCount').innerText = compareList.size;
    bar.style.display = compareList.size > 0 ? 'flex' : 'none';
}

function clearComparison() {
    compareList.clear();
    updateCompareBar();
    renderCars(allCars);
}

async function showComparisonModal() {
    const modal = document.getElementById('compareModal');
    const container = document.getElementById('compare-body');
    container.innerHTML = '<p>Loading...</p>';
    modal.style.display = 'flex';

    try {
        const promises = Array.from(compareList).map(id => fetch(`/api/cars/${id}`).then(r => r.json()));
        const cars = await Promise.all(promises);

        let html = '<div class="comparison-table-wrapper"><table class="comparison-table"><tr><th>Feature</th>';
        cars.forEach(c => html += `<td><img src="${c.image ? 'img/'+c.image : ''}" style="width:100px"><br>${c.name}</td>`);
        html += '</tr>';

        const rows = [
            { label: 'Year', key: 'year' },
            { label: 'Category', isCat: true },
            { label: 'Engine', key: 'engine', isSpec: true },
            { label: 'Horsepower', key: 'horsepower', isSpec: true },
            { label: 'Transmission', key: 'transmission', isSpec: true }
        ];

        rows.forEach(row => {
            html += `<tr><th>${row.label}</th>`;
            cars.forEach(car => {
                let val = 'N/A';
                if (row.isSpec && car.specifications) val = car.specifications[row.key] || 'N/A';
                else if (row.isCat) {
                    const c = categories.find(cat => cat.id == car.categoryId);
                    val = c ? c.name : 'N/A';
                }
                else val = car[row.key] || 'N/A';
                html += `<td>${val}</td>`;
            });
            html += '</tr>';
        });
        html += '</table></div>';
        container.innerHTML = html;
    } catch (e) { container.innerHTML = '<p>Error.</p>'; }
}

function closeCompareModal() {
    document.getElementById('compareModal').style.display = 'none';
}