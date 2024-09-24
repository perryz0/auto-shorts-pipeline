# auto-shorts-pipeline
Auto Shorts Content Generator Project (Work in Progress)

## Project Overview
The `auto-shorts-pipeline` automates the generation of short-form video content using backend services, AI-generated scripts, and video processing. The backend manages user input, interacts with AI models like GPT-4, handles video assets, and integrates scripts with clips to produce engaging shorts.

### Folder Structure

- **cmd**: Main entry point for running the backend (`main.go`)
- **internal**: Core backend implementation
    - **handlers**: Request processing logic
    - **app**: Application-specific logic
    - **pkg**: Shared internal utilities and data models
        - **utils**: Utility functions (e.g., HTTP clients, config management)
        - **models**: Data structures for API requests, responses, and video data
        - **logging**: Centralized logging
- **pkg**: Public libraries for reusable components
- **web**: Web UI components (static assets, HTML templates, JavaScript)
- **api**: API specs (OpenAPI/Swagger)
- **configs**: Configuration files and templates
- **scripts**: Dev, testing, and deployment scripts (`setup.sh`)
- **test**: Testing resources and mock data

### Getting Started

1. **Setup the Project**:
    - Clone the repository:
        ```bash
        git clone https://github.com/your-username/auto-shorts-pipeline.git
        cd auto-shorts-pipeline
        ```
    - Set up the backend by navigating to the `cmd/backend` directory and running:
        ```bash
        go run main.go
        ```

2. **Configuration**:
    - Set up necessary environment variables (e.g., API keys) as defined in `config.yaml`

3. **Running the Pipeline**:
    - Use the terminal or web UI to input content descriptions and generate shorts using the backend services

### Contributing
Contributions are welcome! Submit pull requests or open issues for bugs or feature requests

### License
This project is licensed under the terms of the [LICENSE](LICENSE) file in this repository
