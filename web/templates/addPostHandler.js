function redirectToHomePage() {
    window.location.href = '/';
}

function previewImage(event) {
    const file = event.target.files[0];
    const imagePreview = document.getElementById('imagePreview');
    const uploadButton = document.getElementById('uploadButton');

    if (file) {
        const reader = new FileReader();
        reader.onload = function(e) {
            imagePreview.src = e.target.result;
            imagePreview.style.display = 'block';
            uploadButton.textContent = 'Remove Photo';
            uploadButton.classList.remove('upload-button');
            uploadButton.classList.add('remove-button');
            uploadButton.setAttribute('onclick', 'removeImage()');
        }
        reader.readAsDataURL(file);
    } else {
        imagePreview.style.display = 'none';
        uploadButton.textContent = 'Upload Photo';
        uploadButton.classList.remove('remove-button');
        uploadButton.classList.add('upload-button');
        uploadButton.setAttribute('onclick', 'handleImageUpload()');
    }
}

function handleImageUpload() {
    document.getElementById('fileInput').click();
}

function removeImage() {
    const imagePreview = document.getElementById('imagePreview');
    const uploadButton = document.getElementById('uploadButton');
    imagePreview.src = '';
    imagePreview.style.display = 'none';
    document.getElementById('fileInput').value = '';
    uploadButton.textContent = 'Upload Photo';
    uploadButton.classList.remove('remove-button');
    uploadButton.classList.add('upload-button');
    uploadButton.setAttribute('onclick', 'handleImageUpload()');
}

document.getElementById('addPost').addEventListener('submit', function(event) {
    event.preventDefault();
    var formData = new FormData(this);
    var fileInput = document.getElementById('fileInput');
    if (fileInput.files.length > 0) {
        formData.append('postImage', fileInput.files[0]);
    }
    fetch('/addPost', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (response.ok) {
                window.location.href = '/';
            } else {
                alert('Ошибка при добавлении поста. Пожалуйста, попробуйте ещё раз.');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Ошибка при добавлении поста. Пожалуйста, попробуйте ещё раз.');
        });
});

