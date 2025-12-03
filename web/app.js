// State
let selectedLanguage = 'cpp'; // Default to C++
let manifestContent = null;
let generatedFiles = {};
let wasmReady = false;

// Language mapping for syntax highlighting
const languageMap = {
    'c': 'c',
    'cpp': 'cpp',
    'v8': 'typescript',
    'golang': 'go',
    'dotnet': 'csharp',
    'python': 'python',
    'lua': 'lua',
    'd': 'd',
};

// DOM elements
const fileInput = document.getElementById('fileInput');
const fileName = document.getElementById('fileName');
const dropZone = document.getElementById('dropZone');
const autoConvert = document.getElementById('autoConvert');
const generateClasses = document.getElementById('generateClasses');
const languageButtons = document.querySelectorAll('.lang-btn');
const convertBtn = document.getElementById('convertBtn');
const statusDiv = document.getElementById('status');
const outputDiv = document.getElementById('output');
const fileList = document.getElementById('fileList');
const downloadAllBtn = document.getElementById('downloadAllBtn');
const wasmStatus = document.getElementById('wasmStatus');

// Load WebAssembly
async function loadWasm() {
    try {
        wasmStatus.textContent = 'Loading WebAssembly module...';
        wasmStatus.className = 'wasm-status loading';

        const go = new Go();
        const result = await WebAssembly.instantiateStreaming(
            fetch('plugify-gen.wasm'),
            go.importObject
        );

        // Run Go program (this is async but doesn't return a promise)
        go.run(result.instance);

        // Wait for functions to be registered
        let attempts = 0;
        while (!window.convertManifest && attempts < 50) {
            await new Promise(resolve => setTimeout(resolve, 100));
            attempts++;
        }

        if (!window.convertManifest) {
            throw new Error('WASM functions not registered after timeout');
        }

        wasmReady = true;
        wasmStatus.textContent = 'WebAssembly module ready';
        wasmStatus.className = 'wasm-status ready';
        console.log('WASM loaded successfully');

        updateConvertButton();
    } catch (err) {
        console.error('Failed to load WebAssembly:', err);
        wasmStatus.textContent = 'Error loading WebAssembly module';
        wasmStatus.className = 'wasm-status error';
        showStatus('Failed to load WebAssembly module. Please refresh the page.', 'error');
    }
}

// Perform conversion
async function performConversion() {
    if (!manifestContent || !selectedLanguage) {
        return;
    }

    if (!window.convertManifest) {
        showStatus('WebAssembly module not ready. Please wait and try again.', 'error');
        return;
    }

    try {
        showStatus('Converting...', 'info');
        convertBtn.disabled = true;

        // Call WebAssembly function
        const options = {
            generateClasses: generateClasses.checked
        };
        const result = window.convertManifest(manifestContent, selectedLanguage, options);

        if (result && result.success) {
            generatedFiles = result.files;
            displayResults();
            showStatus(`Successfully generated ${Object.keys(generatedFiles).length} file(s)`, 'success');
        } else if (result && result.error) {
            showStatus(`Error: ${result.error}`, 'error');
            outputDiv.style.display = 'none';
        } else {
            showStatus('Conversion failed with unknown error', 'error');
            outputDiv.style.display = 'none';
        }
    } catch (err) {
        console.error('Conversion error:', err);
        showStatus(`Unexpected error: ${err.message}`, 'error');
        outputDiv.style.display = 'none';
    } finally {
        updateConvertButton();
    }
}

// Handle file processing
function handleFile(file) {
    if (!file) {
        manifestContent = null;
        fileName.textContent = 'No file selected';
        fileName.classList.remove('has-file');
        updateConvertButton();
        return;
    }

    // Check file extension
    const validExtensions = ['.pplugin', '.json', '.jsonc'];
    const fileExt = file.name.substring(file.name.lastIndexOf('.')).toLowerCase();

    if (!validExtensions.includes(fileExt)) {
        showStatus('Please select a .pplugin or .json file', 'error');
        return;
    }

    fileName.textContent = file.name;
    fileName.classList.add('has-file');

    const reader = new FileReader();
    reader.onload = (e) => {
        manifestContent = e.target.result;
        updateConvertButton();
        showStatus('Manifest file loaded successfully', 'success');

        // Auto-convert if enabled
        if (autoConvert.checked && selectedLanguage && wasmReady) {
            setTimeout(() => performConversion(), 100);
        }
    };
    reader.onerror = () => {
        showStatus('Error reading file', 'error');
        manifestContent = null;
        fileName.classList.remove('has-file');
        updateConvertButton();
    };
    reader.readAsText(file);
}

// File input handler
fileInput.addEventListener('change', (e) => {
    handleFile(e.target.files[0]);
});

// Drag and drop handlers
dropZone.addEventListener('dragover', (e) => {
    e.preventDefault();
    e.stopPropagation();
    dropZone.classList.add('drag-over');
});

dropZone.addEventListener('dragleave', (e) => {
    e.preventDefault();
    e.stopPropagation();
    dropZone.classList.remove('drag-over');
});

dropZone.addEventListener('drop', (e) => {
    e.preventDefault();
    e.stopPropagation();
    dropZone.classList.remove('drag-over');

    const files = e.dataTransfer.files;
    if (files.length > 0) {
        handleFile(files[0]);
    }
});

// Click on drop zone to trigger file input
dropZone.addEventListener('click', (e) => {
    if (e.target !== fileInput && !e.target.closest('.file-label')) {
        fileInput.click();
    }
});

// Language selection
languageButtons.forEach(btn => {
    btn.addEventListener('click', () => {
        languageButtons.forEach(b => b.classList.remove('selected'));
        btn.classList.add('selected');
        selectedLanguage = btn.dataset.lang;
        updateConvertButton();

        // Auto-convert if enabled
        if (autoConvert.checked && manifestContent && wasmReady) {
            setTimeout(() => performConversion(), 100);
        }
    });
});

// Update convert button state
function updateConvertButton() {
    convertBtn.disabled = !(wasmReady && manifestContent && selectedLanguage);
}

// Convert button handler
convertBtn.addEventListener('click', () => {
    performConversion();
});

// Get appropriate syntax highlighting language based on file extension
function getHighlightLanguage(filename, defaultLang) {
    const ext = filename.substring(filename.lastIndexOf('.')).toLowerCase();

    // Map file extensions to highlight.js languages
    const extMap = {
        '.h': 'c',
        '.c': 'c',
        '.d': 'd',
        '.hpp': 'cpp',
        '.cpp': 'cpp',
        '.go': 'go',
        '.cs': 'csharp',
        '.py': 'python',
        '.pyi': 'python',
        '.lua': 'lua',
        '.js': 'javascript',
        '.ts': 'typescript',
        '.d.ts': 'typescript'
    };

    // Check for .d.ts first (before .ts)
    if (filename.endsWith('.d.ts')) {
        return 'typescript';
    }

    return extMap[ext] || defaultLang || 'plaintext';
}

// Display results
function displayResults() {
    fileList.innerHTML = '';

    // Get the default highlight.js language name for this generator
    const defaultLang = languageMap[selectedLanguage] || 'plaintext';

    for (const [filename, content] of Object.entries(generatedFiles)) {
        const fileItem = document.createElement('div');
        fileItem.className = 'file-item';

        const header = document.createElement('div');
        header.className = 'file-item-header';

        const name = document.createElement('div');
        name.className = 'file-item-name';
        name.textContent = filename;

        const actions = document.createElement('div');
        actions.className = 'file-item-actions';

        const copyBtn = document.createElement('button');
        copyBtn.className = 'file-btn copy';
        copyBtn.textContent = 'Copy';
        copyBtn.onclick = () => copyToClipboard(content, copyBtn);

        const downloadBtn = document.createElement('button');
        downloadBtn.className = 'file-btn';
        downloadBtn.textContent = 'Download';
        downloadBtn.onclick = () => downloadFile(filename, content);

        actions.appendChild(copyBtn);
        actions.appendChild(downloadBtn);

        header.appendChild(name);
        header.appendChild(actions);

        // Determine the appropriate language for this file
        const hlLanguage = getHighlightLanguage(filename, defaultLang);

        // Create code block with syntax highlighting
        const pre = document.createElement('pre');
        const code = document.createElement('code');
        code.className = `language-${hlLanguage} file-content`;
        code.textContent = content;

        // Apply syntax highlighting
        if (typeof hljs !== 'undefined') {
            hljs.highlightElement(code);
        }

        pre.appendChild(code);
        pre.className = 'file-content-wrapper';

        fileItem.appendChild(header);
        fileItem.appendChild(pre);
        fileList.appendChild(fileItem);
    }

    outputDiv.style.display = 'block';
}

// Copy to clipboard
async function copyToClipboard(text, button) {
    try {
        // Try modern clipboard API first (requires HTTPS or localhost)
        if (navigator.clipboard && window.isSecureContext) {
            await navigator.clipboard.writeText(text);
        } else {
            // Fallback for non-secure contexts (HTTP, file://)
            const textArea = document.createElement('textarea');
            textArea.value = text;
            textArea.style.position = 'fixed';
            textArea.style.left = '-9999px';
            textArea.style.top = '-9999px';
            document.body.appendChild(textArea);
            textArea.focus();
            textArea.select();
            try {
                document.execCommand('copy');
            } finally {
                document.body.removeChild(textArea);
            }
        }
        const originalText = button.textContent;
        button.textContent = 'Copied!';
        setTimeout(() => {
            button.textContent = originalText;
        }, 2000);
    } catch (err) {
        console.error('Copy failed:', err);
        showStatus('Failed to copy to clipboard', 'error');
    }
}

// Download single file
function downloadFile(filename, content) {
    const blob = new Blob([content], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = String(filename.split('/').pop() ?? filename);
    a.click();
    URL.revokeObjectURL(url);
}

// Download all files as ZIP
downloadAllBtn.addEventListener('click', async () => {
    try {
        // Use JSZip if available, otherwise download files individually
        if (typeof JSZip !== 'undefined') {
            const zip = new JSZip();

            for (const [filename, content] of Object.entries(generatedFiles)) {
                zip.file(filename, content);
            }

            const blob = await zip.generateAsync({ type: 'blob' });
            const url = URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = 'generated-bindings.zip';
            a.click();
            URL.revokeObjectURL(url);

            showStatus('ZIP file downloaded', 'success');
        } else {
            // Fallback: download files individually
            for (const [filename, content] of Object.entries(generatedFiles)) {
                downloadFile(filename, content);
                await new Promise(resolve => setTimeout(resolve, 100));
            }
            showStatus('Files downloaded individually', 'success');
        }
    } catch (err) {
        console.error('Download error:', err);
        showStatus('Error creating download', 'error');
    }
});

// Show status message
function showStatus(message, type) {
    statusDiv.textContent = message;
    statusDiv.className = `status ${type}`;

    if (type === 'success' || type === 'info') {
        setTimeout(() => {
            statusDiv.style.display = 'none';
        }, 5000);
    }
}

// Initialize default language selection
function initializeDefaultLanguage() {
    const cppButton = document.querySelector('.lang-btn[data-lang="cpp"]');
    if (cppButton) {
        cppButton.classList.add('selected');
    }
}

// Load auto-convert preference from localStorage
function loadAutoConvertPreference() {
    const savedPreference = localStorage.getItem('autoConvert');
    if (savedPreference !== null) {
        autoConvert.checked = savedPreference === 'true';
    }
}

// Save auto-convert preference to localStorage
function saveAutoConvertPreference() {
    localStorage.setItem('autoConvert', autoConvert.checked);
}

// Load generate-classes preference from localStorage
function loadGenerateClassesPreference() {
    const savedPreference = localStorage.getItem('generateClasses');
    if (savedPreference !== null) {
        generateClasses.checked = savedPreference === 'true';
    }
}

// Save generate-classes preference to localStorage
function saveGenerateClassesPreference() {
    localStorage.setItem('generateClasses', generateClasses.checked);
}

// Auto-convert toggle change handler
autoConvert.addEventListener('change', () => {
    saveAutoConvertPreference();
});

// Generate classes toggle change handler
generateClasses.addEventListener('change', () => {
    saveGenerateClassesPreference();

    // Auto-convert if enabled
    if (autoConvert.checked && manifestContent && wasmReady) {
        setTimeout(() => performConversion(), 100);
    }
});

// Initialize
initializeDefaultLanguage();
loadAutoConvertPreference();
loadGenerateClassesPreference();
loadWasm();