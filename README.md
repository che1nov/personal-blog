# Personal Blog

A personal blog to write and publish articles on various topics. The blog includes both a guest section and an admin section.

## Features

### Guest Section
- **Home Page**: Displays the list of articles published on the blog.
- **Article Page**: Displays the content of the article along with the date of publication.

### Admin Section
- **Dashboard**: Displays the list of articles published on the blog with options to add, edit, or delete articles.
- **Add Article Page**: Contains a form to add a new article with fields for title, content, and date of publication.
- **Edit Article Page**: Contains a form to edit an existing article with fields for title, content, and date of publication.

## Implementation Guidelines

### Storage
- Articles are stored on the filesystem.
- Each article is stored as a separate file in a directory.
- JSON or Markdown format can be used to store articles.

### Backend
- Use any programming language to build the backend of the blog.
- Render HTML directly from the server.
- Handle form submissions on the server.

### Frontend
- Use HTML and CSS for the frontend.
- Use any templating engine to render articles on the frontend.

### Authentication
- Implement basic authentication for the admin section.
- Use standard HTTP basic authentication or hardcode the username and password in the code.
- Create a simple login page that creates a session for the admin.

## Getting Started

### Prerequisites
- Go 1.15 or later

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/che1nov/personal-blog.git
   ```

2. Navigate to the project directory:
   ```sh
   cd personal-blog
   ```

### Running the Application

To start the application, run the following command:
```sh
go run main.go
```

The application will be accessible at `http://localhost:8080`.


### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

https://roadmap.sh/projects/personal-blog