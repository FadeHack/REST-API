document.addEventListener('DOMContentLoaded', function() {
    displayBlogs();

    // Toggle Create Form Button Click
    document.getElementById('toggle-create-button').addEventListener('click', function() {
        var createForm = document.getElementById('create-form');
        createForm.style.display = (createForm.style.display === 'none') ? 'block' : 'none';
    });

    // Cancel Create Blog
    document.getElementById('cancel-button').addEventListener('click', function() {
        document.getElementById('create-form').style.display = 'none';
    });

    // Create New Blog
    document.getElementById('create-button').addEventListener('click', createItem);

    // Search Blog by ID
    document.getElementById('search-form').addEventListener('submit', searchItem);

    // Back Arrow Click
    document.getElementById('back-arrow').addEventListener('click', function() {
        // Reset search input and display all blogs
        document.getElementById('search-id').value = '';
        displayBlogs();
        // Disable back arrow again
        document.getElementById('back-arrow').style.cursor = 'not-allowed';
    });

    // Enable or disable back arrow based on search input
    document.getElementById('search-id').addEventListener('input', function() {
        const backArrow = document.getElementById('back-arrow');
        backArrow.style.cursor = this.value === '' ? 'not-allowed' : 'pointer';
    });
});

// Function to fetch and display blogs
function displayBlogs() {
    fetch('http://localhost:8000/blogs')
    .then(response => response.json())
    .then(data => {
        const blogList = document.getElementById('blog-list');
        blogList.innerHTML = '';

        data.data.forEach(blog => {
            const card = `
    <div class="card mt-3">
        <div class="card-body">
            <div class="row">
                <div class="col-md-8">
                    <div>
                        <h5 class="card-title">${blog.Title}</h5>
                        <h6 class="card-subtitle mb-2 text-muted">${blog.SubTitle}</h6>
                        <p class="card-text">${blog.Content}</p>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="ml-3">
                        <p class="text-info">ID: ${blog.ID}</p>
                        <button class="btn btn-primary" onclick="editBlog('${blog.ID}')"><i class="fas fa-edit"></i> Edit</button>
                        <button class="btn btn-danger" onclick="deleteBlog('${blog.ID}')"><i class="fas fa-trash-alt"></i> Delete</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
`;


            blogList.innerHTML += card;
        });
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

// Function to create a new blog
function createItem(e) {
    e.preventDefault();

    var title = document.getElementById('create-title').value;
    var subtitle = document.getElementById('create-subtitle').value;
    var content = document.getElementById('create-content').value;

    fetch('http://localhost:8000/blog', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            Title: title,
            SubTitle: subtitle,
            Content: content,
        }),
    })
    .then(response => response.json())
    .then(data => {
        alert('Blog created successfully');
        document.getElementById('create-form').style.display = 'none';
        displayBlogs();
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

// Function to search for a blog by ID
function searchItem(e) {
    e.preventDefault();

    var id = document.getElementById('search-id').value;

    fetch('http://localhost:8000/blog/' + id)
    .then(response => response.json())
    .then(data => {
        data = data.data;
        if (id === data.ID) {
        const blogList = document.getElementById('blog-list');
        
        blogList.innerHTML = '';

        const card = `
            <div class="card mt-3">
                <div class="card-body">
                    <div class="d-flex justify-content-between align-items-center">
                        <div>
                            <h5 class="card-title">${data.Title}</h5>
                            <h6 class="card-subtitle mb-2 text-muted">${data.SubTitle}</h6>
                            <p class="card-text">${data.Content}</p>
                        </div>
                        <div class="ml-3">
                            <p class="text-info">ID: ${data.ID}</p>
                            <button class="btn btn-primary" onclick="editBlog('${data.ID}')"><i class="fas fa-edit"></i> Edit</button>
                            <button class="btn btn-danger" onclick="deleteBlog('${data.ID}')"><i class="fas fa-trash-alt"></i> Delete</button>
                        </div>
                    </div>
                </div>
            </div>
        `;

        blogList.innerHTML = card;
        // Enable back arrow after search
        document.getElementById('back-arrow').style.cursor = 'pointer';
        }
        else{
        alert("Not Found !! (Enter Valid ID)")
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

// Function to edit a blog
function editBlog(id) {
    // Fetch the specific blog data
    fetch('http://localhost:8000/blog/' + id)
    .then(response => response.json())
    .then(data => {
        // Display the edit form for the specific blog
        displayEditForm(data);
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

// Function to display the edit form for the specific blog
function displayEditForm(blog) {
    blog = blog.data
    const editForm = document.createElement('div');
    editForm.innerHTML = `
        <h2><i class="fas fa-edit"></i> Edit Blog</h2>
        <form id="edit-item-form">
            <div class="form-group">
                <label for="edit-title">Title</label>
                <input type="text" id="edit-title" class="form-control" value="${blog.Title}">
            </div>
            <div class="form-group">
                <label for="edit-subtitle">Subtitle</label>
                <input type="text" id="edit-subtitle" class="form-control" value="${blog.SubTitle}">
            </div>
            <div class="form-group">
                <label for="edit-content">Content</label>
                <textarea id="edit-content" class="form-control">${blog.Content}</textarea>
            </div>
            <button type="button" id="save-button" class="btn btn-primary"><i class="fas fa-save"></i> Save Changes</button>
            <button type="button" id="cancel-edit-button" class="btn btn-secondary"><i class="fas fa-times"></i> Cancel</button>
        </form>
    `;

    // Replace the existing content with the edit form
    const blogList = document.getElementById('blog-list');
    blogList.innerHTML = '';
    blogList.appendChild(editForm);

    // Add event listeners for save and cancel buttons
    document.getElementById('save-button').addEventListener('click', function() {
        saveChanges(blog.ID);
    });

    document.getElementById('cancel-edit-button').addEventListener('click', function() {
        displayBlogs(); // Display the blogs again without changes
    });
}

// Function to save changes for the specific blog
function saveChanges(id) {
    var title = document.getElementById('edit-title').value;
    var subtitle = document.getElementById('edit-subtitle').value;
    var content = document.getElementById('edit-content').value;

    fetch('http://localhost:8000/ublog/' + id, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            Title: title,
            SubTitle: subtitle,
            Content: content,
        }),
    })
    .then(response => {
        if (response.status === 200) {
            alert('Changes saved successfully');
            displayBlogs();
        } else {
            throw new Error('Failed to save changes');
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

// Function to delete a blog
function deleteBlog(id) {
    console.log('Deleting blog with ID:', id); // Add this line for debugging

    fetch('http://localhost:8000/dblog/' + id, {
        method: 'DELETE',
    })
    .then(response => {
        if (response.status === 204) {
            alert('Blog deleted successfully', id);
            displayBlogs();
        } else {
            throw new Error('Failed to delete blog');
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

