<h1>API Rate Limiter with Golang and Frontend UI</h1>
<p>This project implements an API Rate Limiter using the Token Bucket Algorithm in Golang, along with a simple UI built using HTML, CSS, and JavaScript. The rate limiter ensures that users can only make a certain number of API requests within a given time frame, while the UI provides a user-friendly interface to interact with the API and observe rate-limiting behavior.</p>
<h2>Features:</h2>
<ul>
  <li>Rate Limiting:</li>
  <ul>
    <li>Implements a Token Bucket Algorithm for rate limiting.</li>
    <li>Configurable request limits (e.g., 5 requests per minute per user).</li>
    <li>Returns a 429 Too Many Requests status code when the limit is exceeded.</li>
  </ul>
  <li>UI:</li>
  <ul>
    <li>Simple, responsive frontend with a button to send API requests.</li>
    <li>Displays the response ("API request successful!" or "Rate limit exceeded.") directly in the browser.</li>
    <li>Built using HTML, CSS, and JavaScript.</li>
  </ul>
  <li>Backend:</li>
  <ul>
    <li>REST API built with Golang and Gorilla Mux.</li>
    <li>Middleware to handle rate limiting.</li>
    <li>Static file server to serve frontend assets.</li>
  </ul>
</ul>


