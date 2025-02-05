# AI-Powered Cybersecurity Port Analyzer - Day 1 of 30 Day 30 Hack Tool Challenge

[![Day 1](https://img.shields.io/badge/Day-1-brightgreen)](https://github.com/rasperon/portscanner)

This project is the first tool developed as part of my 30-day, 30-hack tool challenge. It leverages the power of AI to analyze open network ports and provide a detailed security assessment report.  It utilizes the Gemini AI model to identify potential vulnerabilities and provide recommendations.

## Features

*   **AI-Powered Analysis:** Uses the Gemini AI model to analyze open ports.
*   **Detailed Security Assessment:** Provides a comprehensive report on security implications and vulnerabilities.
*   **Risk Level Assessment:**  Identifies the risk level (Low/Medium/High) for each open port.
*   **Specific Recommendations:**  Offers specific recommendations to secure each port.
*   **Markdown Output:**  Saves the analysis report to a Markdown file (`ai_output.md`).

## Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/rasperon/portscanner.git
    cd portscanner
    ```

2.  **Install dependencies:**

    ```bash
    go mod download
    ```

3.  **Set up your Gemini API Key:**

    *   Obtain a Gemini API key from [Google AI Studio](https://makersuite.google.com/).
    *   **Important:**  Do NOT hardcode your API key directly in the code!  Instead, set it as an environment variable.

        ```bash
        export GEMINI_API_KEY="YOUR_API_KEY"
        ```

4.  **Run the application:**

    ```bash
    go run main.go
    ```

## Usage

1.  **Prepare your open ports data:**  Provide a list of open ports in the following format:

    ```
    22/tcp open  ssh
    80/tcp open  http
    443/tcp open  https
    ```

2.  **Run the `AnalyzePorts` function** in `api/api.go` file with the open ports data and desired language.

## Configuration

You can configure the following parameters in the `api/api.go` file:

*   **Model:** Choose a Gemini model (e.g., `gemini-pro`).
*   **Temperature:** Adjust the temperature for more or less creative responses.
*   **Max Output Tokens:**  Control the maximum length of the generated report.

## Disclaimer

This tool is intended for educational and ethical hacking purposes only. Use it responsibly and only on systems you have explicit permission to test. The author is not responsible for any misuse of this tool.

## Contributing

Contributions are welcome!  Please submit pull requests with bug fixes, improvements, or new features.

## License

[MIT License](LICENSE)