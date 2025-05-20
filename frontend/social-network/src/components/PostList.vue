<template>
  <div>
    <CreatePost @post-created="addPost" />
    <h2>Posts</h2>
    <div v-for="post in posts" :key="post.id"
      style="border:1px solid #ccc; margin:10px; padding:10px; align-items: center; justify-content: center; text-align: center;">
      <p><strong>Title: {{ post.title }}</strong></p>
      <strong>User {{ post.user_id }}:</strong>
      <P><strong>Post Content:</strong> {{ post.content }}</P>

      <!-- Display post image -->
      <p v-if="post.image">
        <img :src="`http://localhost:8080/${post.image}`" alt="Post image"
          style="max-width: 100%; max-height: 300px; display: block; margin-top: 0.5rem; margin: auto;" />
      </p>

      <p>Privacy: {{ post.privacy }}</p>

      <!-- Comments -->
      <div v-if="post.comments && post.comments.length > 0">
        <h3>Comments:</h3>
        <div v-for="comment in post.comments" :key="comment.id" style="margin-left:10px;">
          <p v-if="comment.image">
            <img :src="`http://localhost:8080/${comment.image}`" alt="comment image"
              style="max-width: 100%; max-height: 300px; display: block; margin-top: 0.5rem; margin: auto;" />
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

onMounted(fetchPosts);
</script>

<style scoped>
h2 {
  align-items: center;
  text-align: center;
}
</style>