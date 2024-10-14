document.getElementById('send-request-btn').addEventListener('click', async function () {
    const responseText = document.getElementById('response-text');

    try {
        // Send a request to the API
        const response = await fetch('/api/test', { method: 'GET' });

        // Display the result
        if (response.ok) {
            const result = await response.text();
            responseText.textContent = result;
            responseText.style.color = 'green';
        } else if (response.status === 429) {
            responseText.textContent = 'Rate limit exceeded. Try again later.';
            responseText.style.color = 'red';
        }
    } catch (error) {
        responseText.textContent = 'Error sending request.';
        responseText.style.color = 'red';
    }
});
