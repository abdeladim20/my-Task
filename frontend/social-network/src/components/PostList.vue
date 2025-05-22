<template>
  <div>
    <CreatePost @post-created="addPost" />
    <h2>Posts</h2>
    <div v-for="post in posts" :key="post.id" class="post-card">
      <div class="post-header-title">
        <div class="post-header">
          <p class="post-user">🧑‍💻 User {{ post.user_id }}</p>
          <p>Privacy: {{ post.privacy }}</p>
          <h3 class="post-title">{{ post.title }}</h3>
        </div>
        <p class="post-content">{{ post.content }} </p>
        <!-- Display post image -->
        <p v-if="post.image">
          <img :src="`http://localhost:8080/${post.image}`" alt="Post image"
            style="max-width: 100%; max-height: 300px; display: block; margin-top: 0.5rem; margin: auto;" />
        </p>
      </div>

      <!-- Like/Dislike Buttons -->
      <div class="post-actions">
        <button @click="likePost(post.id)">👍 Like {{ post.likes || 0 }}</button>
        <button @click="dislikePost(post.id)">👎 Dislike {{ post.dislikes || 0 }}</button>
      </div>


      <!-- Comments -->
      <div v-if="post.comments && post.comments.length > 0" class="comments-section">
        <h3>Comments:</h3>
        <div v-for="comment in post.comments" :key="comment.id" class="comment">
          <p v-if="comment.image">
            <img :src="`http://localhost:8080/${comment.image}`" alt="comment image" />
          </p>
          🗨️ User {{ comment.user_id }}: {{ comment.content }}
        </div>
      </div>

      <!-- Add Comment Form -->
      <form @submit.prevent="createComment(post.id)" style="margin-top:10px;">
        <input v-model="newComments[post.id]" placeholder="Write a comment" type="text" />
        <input type="file" :id="`file-${post.id}`" accept="image/*" style="margin-left: 10px;" />
        <br />
        <button type="submit">Comment</button>
      </form>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import CreatePost from "./CreatePost.vue";

const posts = ref([]);
const newComments = ref({});

const fetchPosts = async () => {
  const res = await fetch("http://localhost:8080/posts");
  const data = await res.json();

  for (const post of data) {
    const commentRes = await fetch(
      `http://localhost:8080/posts/${post.id}/comments`
    );
    const comments = await commentRes.json();
    post.comments = comments;
  }

  posts.value = data;
};

const addPost = (newPost) => {
  posts.value.unshift({ ...newPost, comments: [] });
};

const createComment = async (postID) => {
  const content = newComments.value[postID];
  const fileInput = document.querySelector(`#file-${postID}`);
  const file = fileInput?.files[0];

  if (!content && !file) return alert("Write something or select a file!");

  const formData = new FormData();
  formData.append("post_id", postID);
  formData.append("user_id", 1);
  if (content) formData.append("content", content);
  if (file) formData.append("image", file);

  try {
    const res = await fetch("http://localhost:8080/comments", {
      method: "POST",
      body: formData,
    });

    if (!res.ok) {
      const error = await res.text();
      alert("Failed to add comment: " + error);
      return;
    }

    alert("Comment added!");
    newComments.value[postID] = "";
    fileInput.value = "";
    await fetchPosts();
  } catch (err) {
    console.error("Failed to add comment:", err);
    alert("Failed to add comment. Please try again later.");
  }
};

const userID = 1
const likePost = async (postID) => {
  try {
    const res = await fetch("http://localhost:8080/posts/react", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        post_id: postID,
        user_id: userID,
        type: "like"
      })
    });
    if (!res.ok) throw new Error(await res.text());
    const post = posts.value.find(p => p.id === postID);
    if (post) post.likes = (post.likes || 0) + 1;
    console.log("Reaction saved successfully!");
  } catch (err) {
    console.error("Failed to like post:", err);
    alert("Error liking post.");
  }
};

const dislikePost = async (postID) => {
  try {
    const res = await fetch("http://localhost:8080/posts/react", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        post_id: postID,
        user_id: userID,
        type: "dislike"
      })
    });

    if (!res.ok) throw new Error(await res.text());
    const post = posts.value.find(p => p.id === postID);
    if (post) post.dislikes = (post.dislikes || 0) + 1;
    console.log("Reaction saved successfully!");
  } catch (err) {
    console.error("Failed to dislike post:", err);
    alert("Error disliking post.");
  }
};

onMounted(fetchPosts);
</script>

<style scoped>
/* General styles */
h2 {
  text-align: center;
  margin-top: 1rem;
  font-size: 1.8rem;
  color: #333;
}

/* Post card style */
.post-card {
  border: 1px solid #ddd;
  border-radius: 10px;
  padding: 15px;
  margin: 20px auto;
  max-width: 700px;
  background-color: #fdfdfd;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition: box-shadow 0.3s ease;
}

.post-card:hover {
  box-shadow: 0 4px 14px rgba(0, 0, 0, 0.1);
}

.post-header-title {
  border: 4px solid white;
  border-radius: 10px;
  padding: 10px;
  background-color: #f0f0f5;
  /* margin-bottom: 10px; */
  /* display: flex; */
}

.post-header {
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
  margin-bottom: 10px;
  text-align: left;
}

.post-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
  text-align: left;
}

.post-user {
  font-size: 0.9rem;
  color: #777;
  margin-top: 5px;
}

.post-content {
  font-size: 1rem;
  color: #444;
  line-height: 1.6;
  text-align: left;
  white-space: pre-wrap;
  word-wrap: break-word;
  padding: 10px;
}

/* Image styles */
.post-card img {
  max-width: 100%;
  max-height: 300px;
  margin-top: 10px;
  border-radius: 6px;
}

.post-actions {
  margin-top: 10px;
  display: flex;
  justify-content: center;
  gap: 10px;
}

.post-actions button {
  padding: 6px 12px;
  border: none;
  background-color: #e0e0e0;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.95rem;
  transition: background-color 0.3s ease;
}

.post-actions button:hover {
  background-color: #d0d0d0;
}

/* Comments section */
.comments-section {
  background-color: #fafafa;
  padding: 10px;
  margin-top: 15px;
  border-radius: 6px;
}

.comment {
  background-color: #f0f0f5;
  padding: 8px 12px;
  margin-top: 10px;
  border-left: 4px solid #3498db;
  border-radius: 4px;
  font-size: 0.95rem;
}

.comment img {
  max-width: 100%;
  max-height: 200px;
  margin: 5px 0;
  border-radius: 4px;
}

/* Form styling */
form {
  margin-top: 15px;
}

form input[type="text"] {
  padding: 6px 10px;
  width: 60%;
  border: 1px solid #ccc;
  border-radius: 4px;
}

form input[type="file"] {
  margin-left: 10px;
}

form button {
  margin-top: 8px;
  padding: 6px 14px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

form button:hover {
  background-color: #2980b9;
}
</style>