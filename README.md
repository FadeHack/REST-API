```markdown
# Blog Application

## Technology Specification

This blog application is built using the following technologies:

- **Backend:**
  - Language: Go (Golang)
  - Framework: gofiber (Gofr)
  - Database: MongoDB
  - Dependencies: Go Modules

- **Frontend:**
  - HTML, CSS, JavaScript
  - Bootstrap for styling
  - jQuery for DOM manipulation

## Usage

### Create New Blog
Endpoint: `POST /blog`

Create a new blog by providing the title, subtitle, and content.

### Get All Blogs
Endpoint: `GET /blogs`

Retrieve a list of all blogs.

### Get Blog By ID
Endpoint: `GET /blog/{id}`

Retrieve a specific blog by its ID.

### Update Blog By ID
Endpoint: `PUT /blog/{id}`

Update the title, subtitle, and content of a specific blog.

### Delete Blog By ID
Endpoint: `DELETE /blog/{id}`

Delete a specific blog by its ID.

## Postman Collection
[![Run in Postman][<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/31394686-037e2999-fd0f-4d4f-a6c5-76a02c6923a0?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D31394686-037e2999-fd0f-4d4f-a6c5-76a02c6923a0%26entityType%3Dcollection%26workspaceId%3D42632939-a688-4bfc-baad-b3f4c2d2d9bf)]

## Screenshots

### 1. Create New Blog

Frontend:
![Create New Blog - Frontend](path/to/frontend-screenshot.png)

Postman Response         | Postman Tests
------------------------- | -------------------------
![Create New Blog - Postman Response](path/to/postman-response-screenshot.png) | ![Create New Blog - Postman Tests](path/to/postman-tests-screenshot.png)

### 2. Get All Blogs

...

### 3. Get Blog By ID

...

### 4. Update Blog By ID

...

### 5. Delete Blog By ID

...

## How to Run

1. **Run Backend:**
   - Open terminal in the backend directory.
   - Run `go mod tidy` to tidy up dependencies.
   - Run `go run main.go` to start the backend server.

2. **Run Frontend:**
   - Open terminal in the frontend directory.
   - For local usage, use the "Live Server" extension to run the frontend.
```