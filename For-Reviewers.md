# Notes For Reviewers
Core Features:
1. HTTP Server listening on port 8080 for HTTP traffic.
2. CRUD Endpoints:
   - GET /v1/risks to list all risks.
   - POST /v1/risks to create a new risk.
   - GET /v1/risks/{id} to retrieve a specific risk by its ID.
   - PUT /v1/risks/{id} to update an existing risk.
3. Risk Data Structure:
   - Risk ID (UUID auto-generated).
   - State (open, closed, accepted, investigating).
   - Title and Description.

Improvements Beyond Requirements:
- Modular Architecture: Core components separated into:
  - v1 (API version), common, logger, and validation packages.
- Logging: Middleware to log HTTP requests, methods, and processing times.
- In-Memory Storage: Thread-safe storage using a mutex to ensure concurrency.
- Validation: Custom state validation to enforce correct risk states.
- Added an update API, as we would need to update the risk records at some point

Testing:
- Unit Tests: Created for core API handlers to ensure all individual functionalities operate correctly.

Extra Enhancements:
- Detailed Logging for all events and errors.
- README Documentation: Clear instructions, usage examples, architecture details, and API references.
- Future Enhancements Suggested: Persistent database, token-based authentication, and external logging integration.

