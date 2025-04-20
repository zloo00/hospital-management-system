// Common JavaScript functions

// Set up CSRF token for all AJAX requests
document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('token');
    if (token) {
        // Set token for all fetch requests
        const originalFetch = window.fetch;
        window.fetch = function(url, options = {}) {
            options.headers = options.headers || {};
            options.headers['Authorization'] = `Bearer ${token}`;
            return originalFetch(url, options);
        };
    }

    // Check if user is logged in
    checkAuthStatus();
});

function checkAuthStatus() {
    const token = localStorage.getItem('token');
    const loginLinks = document.querySelectorAll('a[href="/login"], a[href="/register"]');
    const logoutLinks = document.querySelectorAll('a[href="/logout"]');

    if (token) {
        loginLinks.forEach(link => link.style.display = 'none');
        logoutLinks.forEach(link => link.style.display = 'block');
    } else {
        loginLinks.forEach(link => link.style.display = 'block');
        logoutLinks.forEach(link => link.style.display = 'none');
    }
}

// Logout function
document.querySelectorAll('a[href="/logout"]').forEach(link => {
    link.addEventListener('click', function(e) {
        e.preventDefault();
        localStorage.removeItem('token');
        window.location.href = '/login';
    });
});

// Form handling utilities
function handleFormSubmit(formId, endpoint, method = 'POST', successCallback) {
    const form = document.getElementById(formId);
    if (!form) return;

    form.addEventListener('submit', async function(e) {
        e.preventDefault();

        const formData = new FormData(form);
        const data = {};
        formData.forEach((value, key) => data[key] = value);

        try {
            const response = await fetch(endpoint, {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            if (response.ok) {
                const result = await response.json();
                if (successCallback) successCallback(result);
            } else {
                const error = await response.json();
                alert(error.error || 'An error occurred');
            }
        } catch (err) {
            console.error('Error:', err);
            alert('An error occurred');
        }
    });
}
