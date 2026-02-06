// Function to convert UTC text to local time string
function convertUTCToLocal() {
  document.querySelectorAll(".utc-time").forEach(function (element) {
    // The Date constructor automatically interprets ISO 8601 strings (like '...Z') as UTC
    const utcDate = new Date(element.textContent);

    // Convert to the user's local time string, using browser settings for formatting
    element.innerText = utcDate.toLocaleDateString();
  });
}

// Call the function on page load
// document.addEventListener('DOMContentLoaded', convertUTCToLocal);

// Use htmx events to re-run the script after content is swapped into the DOM
document.body.addEventListener('htmx:afterSwap', convertUTCToLocal)
