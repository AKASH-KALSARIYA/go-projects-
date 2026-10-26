const CONFIG = {
    BASE_URL: 'http://localhost:4000',
    ENDPOINTS: {
        REGISTER: '/register',
        LOGIN: '/login',
        PROFILE: '/api/profile'
    }
};

const state = {
    token: localStorage.getItem('nexus_token') || null,
    user: null
};

// UI Elements
const views = {
    placeholder: document.getElementById('view-placeholder'),
    register: document.getElementById('view-register'),
    login: document.getElementById('view-login'),
    profile: document.getElementById('view-profile')
};

const terminal = document.getElementById('responseLog');
const userStatusEl = document.getElementById('userStatus');

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    updateAuthUI();

    // Form Listeners
    document.getElementById('registerForm').addEventListener('submit', handleRegister);
    document.getElementById('loginForm').addEventListener('submit', handleLogin);
});

// Routing
function showView(viewName) {
    // Hide all main containers
    document.querySelectorAll('.view').forEach(v => v.classList.add('hidden'));

    const view = document.getElementById(`view-${viewName}`);
    if (view) {
        view.classList.remove('hidden');
    } else {
        // Handle special views like designer/docs if needed
        createDynamicView(viewName);
    }

    // Update nav links
    document.querySelectorAll('.nav-links a').forEach(a => {
        a.classList.toggle('active', a.getAttribute('href') === `#${viewName}`);
    });

    log(`Switched view to: [${viewName.toUpperCase()}]`, 'system');
}

function createDynamicView(name) {
    const container = document.getElementById('view-placeholder');
    container.classList.remove('hidden');

    if (name === 'designer') {
        container.innerHTML = `
            <h2>API Blueprint Designer</h2>
            <p class="section-desc">Visualizing your authentication architecture.</p>
            <div class="blueprint-grid">
                <div class="blueprint-card">
                    <div class="method-badge post">POST</div>
                    <h4>/register</h4>
                    <p>Parameters: <code>name, email, password</code></p>
                    <div class="flow-line"></div>
                </div>
                <div class="blueprint-card">
                    <div class="method-badge post">POST</div>
                    <h4>/login</h4>
                    <p>Parameters: <code>email, password</code></p>
                    <div class="flow-line"></div>
                </div>
                <div class="blueprint-card">
                    <div class="method-badge get">GET</div>
                    <h4>/api/profile</h4>
                    <p>Headers: <code>Authorization: Bearer &lt;token&gt;</code></p>
                </div>
            </div>
            <div class="designer-footer">
                <p>Status: <span class="gradient-text">Architect Mode Active</span></p>
            </div>
        `;
    } else if (name === 'docs') {
        container.innerHTML = `
            <h2>Documentation</h2>
            <div class="docs-content">
                <h3>Authentication Flow</h3>
                <ol>
                    <li>Register a new identity via <code>/register</code>.</li>
                    <li>Exchange credentials for a JWT via <code>/login</code>.</li>
                    <li>Use the token in the <code>Authorization</code> header for protected routes.</li>
                </ol>
            </div>
        `;
    }
}

function selectEndpoint(endpoint) {
    showView(endpoint);
    log(`Selected endpoint: ${CONFIG.ENDPOINTS[endpoint.toUpperCase()]}`, 'system');
}

// API Functions
async function handleRegister(e) {
    e.preventDefault();
    const name = document.getElementById('regName').value;
    const email = document.getElementById('regEmail').value;
    const password = document.getElementById('regPass').value;

    log(`POST ${CONFIG.ENDPOINTS.REGISTER}...`, 'system');

    try {
        const response = await fetch(`${CONFIG.BASE_URL}${CONFIG.ENDPOINTS.REGISTER}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, password })
        });

        const data = await response.json();
        logResponse(response.status, data);

        if (response.ok) {
            showToast('Registration Account Created Successfully!');
            showView('login');
        }
    } catch (err) {
        log(`Error: ${err.message}`, 'error');
    }
}

async function handleLogin(e) {
    e.preventDefault();
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPass').value;

    log(`POST ${CONFIG.ENDPOINTS.LOGIN}...`, 'system');

    try {
        const response = await fetch(`${CONFIG.BASE_URL}${CONFIG.ENDPOINTS.LOGIN}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await response.json();
        logResponse(response.status, data);

        if (response.ok && data.token) {
            state.token = data.token;
            localStorage.setItem('nexus_token', data.token);
            showToast('Authentication Successful');
            updateAuthUI();
            fetchProfile();
        }
    } catch (err) {
        log(`Error: ${err.message}`, 'error');
    }
}

async function fetchProfile() {
    if (!state.token) {
        log('Error: Authorization token missing. Please login.', 'error');
        showView('login');
        return;
    }

    log(`GET ${CONFIG.ENDPOINTS.PROFILE}...`, 'system');

    try {
        const response = await fetch(`${CONFIG.BASE_URL}${CONFIG.ENDPOINTS.PROFILE}`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${state.token}`,
                'Content-Type': 'application/json'
            }
        });

        const data = await response.json();
        logResponse(response.status, data);

        if (response.ok) {
            updateProfileUI(data);
            showView('profile');
        } else if (response.status === 401) {
            logout();
        }
    } catch (err) {
        log(`Error: ${err.message}`, 'error');
    }
}

// Helpers
function log(msg, type = '') {
    const entry = document.createElement('div');
    entry.className = `log-entry ${type}`;
    entry.textContent = `> ${msg}`;
    terminal.appendChild(entry);
    terminal.scrollTop = terminal.scrollHeight;
}

function logResponse(status, data) {
    const type = status < 300 ? 'success' : 'error';
    log(`HTTP ${status} Response:`, type);
    log(JSON.stringify(data, null, 2), type);
}

function clearTerminal() {
    terminal.innerHTML = '<div class="log-entry system">> Terminal Cleared.</div>';
}

function updateAuthUI() {
    if (state.token) {
        userStatusEl.innerHTML = '<span class="status-indicator online"></span> Authenticated';
    } else {
        userStatusEl.innerHTML = '<span class="status-indicator"></span> Guest Mode';
    }
}

function updateProfileUI(data) {
    document.getElementById('profName').textContent = 'User Authenticated';
    document.getElementById('profEmail').textContent = 'Profile data retrieved from protected route';
    document.getElementById('profStatus').textContent = 'Active';
    document.getElementById('profStatus').className = 'value success';
}

function logout() {
    state.token = null;
    localStorage.removeItem('nexus_token');
    updateAuthUI();
    showView('placeholder');
    log('Session terminated.', 'system');
}

function showToast(msg) {
    const container = document.getElementById('toast-container');
    const toast = document.createElement('div');
    toast.className = 'toast';
    toast.textContent = msg;
    container.appendChild(toast);
    setTimeout(() => toast.remove(), 3000);
}
