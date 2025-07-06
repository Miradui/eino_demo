# Poet AI

Poet AI is a Go-based application built on the Eino framework that generates poetry using a language model. It supports both non-streaming and streaming outputs for generated content.

## Features

- **Interactive CLI**: Input text and receive generated poetry in real-time.
- **Streaming Output**: Stream responses from the language model for a dynamic experience.
- **Conversation History**: Maintains a history of user inputs and model responses.
- **Eino Framework**: Utilizes Eino's `prompt` and `schema` components for chat template management.

## Prerequisites

- Go 1.20 or later
- A valid API key for the Gemini language model

## Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:Miradui/eino_demo.git
   cd eino_demo
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the `.env` file with your Gemini API credentials:
   ```dotenv
   GEMINI_MODEL_NAME=gemini-2.5-flash
   GEMINI_API_KEY=your-api-key
   ```

## Usage

1. Run the application:
   ```bash
   go run poet_ai/main.go
   ```

2. Interact with the CLI:
   - Type your input to generate poetry.
   - Type `exit` to quit the application.

## Project Structure

- `poet_ai/main.go`: Entry point for the application.
- `poet_ai/chat`: Contains functions for interacting with the language model.

## Environment Variables

The application uses the following environment variables:

- `GEMINI_MODEL_NAME`: Specifies the model name.
- `GEMINI_API_KEY`: Your API key for authentication.
