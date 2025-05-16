<template>
  <div>
    <CreatePost @post-created="addPost" />
    <h2>Posts</h2>
    <div v-for="post in posts" :key="post.id" style="border:1px solid #ccc; margin:10px; padding:10px;">
      <p><strong>{{ post.title }}</strong></p>
      <p><strong>User {{ post.user_id }}:</strong> {{ post.content }}</p>
      <p v-if="post.image">Image: {{ post.image }}</p>
      <p>Privacy: {{ post.privacy }}</p>

      <!-- Comments -->
      <div v-if="post.comments">
        <h4>Comments:</h4>
        <div v-for="comment in post.comments" :key="comment.id" style="margin-left:10px;">
          🗨️ User {{ comment.user_id }}: {{ comment.content }}
        </div>
      </div>

      <!-- Add Comment Form -->
      <form @submit.prevent="createComment(post.id)" style="margin-top:10px;">
        <input v-model="newComments[post.id]" placeholder="Write a comment" />
        <button type="submit">Comment</button>
      </form>
    </div>
  </div>
</template>

<script>
import CreatePost from "./CreatePost.vue";

export default {
  data() {
    return {
      posts: [],
      newComments: {}
    };
  },
  methods: {
    async fetchPosts() {
      const res = await fetch("http://localhost:8080/posts");
      const posts = await res.json();

      for (const post of posts) {
        const res = await fetch(`http://localhost:8080/posts/${post.id}/comments`);
        const comments = await res.json();
        post.comments = comments;
      }

      this.posts = posts;
    },

    addPost(newPost) {
      // Directly insert post with empty comments
      this.posts.unshift({ ...newPost, comments: [] });
    },

    async createComment(postID) {
      const content = this.newComments[postID];
      if (!content) return alert("Write something first!");

      const commentData = {
        post_id: postID,
        user_id: 1,
        content: content
      };

      await fetch("http://localhost:8080/comments", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(commentData)
      });

      alert("Comment added!");
      this.newComments[postID] = "";
      this.fetchPosts();
    }
  },
  mounted() {
    this.fetchPosts();
  },
  components: {
    CreatePost
  }
};
</script>

<style scoped>
h2 {
  align-items: center;
  text-align: center;
}
</style>
