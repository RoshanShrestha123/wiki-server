# Simple Go Web Application

This is a simple web application written in Go that demonstrates basic file handling, HTTP request handling.

## Features

- View a page
- Edit a page
- Save a page

## Getting Started

### Prerequisites

- Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/RoshanShrestha123/wiki-server
    ```
2. Navigate into the project directory:
    ```sh
    cd wiki-server
    ```

### Usage

1. Run the application:
    ```sh
    go run main.go
    ```
2. Open your web browser and navigate to `http://localhost:8000`.

### URL Endpoints

- `/view/{title}`: View a page with the given title.
- `/edit/{title}`: Edit a page with the given title.
- `/save/{title}`: Save a page with the given title and body.
- `/`: Root endpoint, displays a simple greeting.

## Example

1. Navigate to `http://localhost:8000/view/TestPage` to view a page named `TestPage`.
2. Navigate to `http://localhost:8000/edit/TestPage` to edit the `TestPage`.
3. Submit the form to save changes to `TestPage`.

## Code Structure

- `Page` struct: Represents a web page with a title and body.
- `save` method: Saves the page content to a text file.
- `loadPage` function: Loads a page from a text file.
- `viewHandler` function: Handles viewing a page.
- `saveHandler` function: Handles saving a page.
- `editHandler` function: Handles editing a page.
- `handler` function: Handles the root endpoint.


