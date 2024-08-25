document.addEventListener("DOMContentLoaded", function() {
    var postList = document.getElementById("posts");
    var postsArray = Array.from(postList.children);
    postsArray.reverse();
    postList.innerHTML = "";
    postsArray.forEach(function(post) {
        postList.appendChild(post);
    });


});
document.querySelectorAll('.blog-post').forEach(post => {
    post.addEventListener('click', function() {
       const postId = this.getAttribute('data-post-id');
       redirectToPostPage(postId);
    });
});


function redirectToAddPostPage() {
    window.location.href = '/add';
}
function redirectToHomePage(){
    window.location.href = "/"
}
function redirectToPostPage(postId) {
    window.location.href = "/posts/" + postId;
}
