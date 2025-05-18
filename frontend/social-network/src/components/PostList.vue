<template>
  <div>
    <CreatePost @post-created="addPost" />
    <h2>Posts</h2>
    <div
      v-for="post in posts"
      :key="post.id"
      style="border:1px solid #ccc; margin:10px; padding:10px;"
    >
      <p><strong>{{ post.title }}</strong></p>
      <p><strong>User {{ post.user_id }}:</strong> {{ post.content }}</p>

      <!-- Display post image -->
      <p v-if="post.image">
        <img
          :src="post.image"
          alt="Post image"
          style="max-width: 100%; max-height: 300px; display: block; margin-top: 0.5rem;"
        />
      </p>

      <p>Privacy: {{ post.privacy }}</p>

      <!-- Comments -->
      <div v-if="post.comments && post.comments.length > 0">
        <h4>Comments:</h4>
        <div
          v-for="comment in post.comments"
          :key="comment.id"
          style="margin-left:10px;"
        >
          🗨️ User {{ comment.user_id }}: {{ comment.content }}
        </div>
      </div>

      <!-- Add Comment Form -->
      <form @submit.prevent="createComment(post.id)" style="margin-top:10px;">
        <input
          v-model="newComments[post.id]"
          placeholder="Write a comment"
          type="text"
        />
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
  if (!content) return alert("Write something first!");

  const commentData = {
    post_id: postID,
    user_id: 1,
    content: content,
  };

  const res = await fetch("http://localhost:8080/comments", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(commentData),
  });

  if (!res.ok) {
    const error = await res.text();
    alert("Failed to add comment: " + error);
    return;
  }
  
  alert("Comment added!");
  newComments.value[postID] = "";
  await fetchPosts();
};

onMounted(fetchPosts);
</script>

<style scoped>
h2 {
  align-items: center;
  text-align: center;
}
</style>