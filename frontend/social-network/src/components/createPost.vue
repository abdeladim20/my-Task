<template>
  <div>
    <h2>Create Post</h2>
    <form @submit.prevent="createPost">
      <input v-model="userID" type="number" placeholder="User ID" required />
      <textarea v-model="content" placeholder="Post content" required></textarea>
      <input v-model="image" type="text" placeholder="Image URL (optional)" />
      <select v-model="privacy">
        <option value="public">Public</option>
        <option value="private">Private</option>
      </select>
      <button type="submit">Post</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userID: "",
      content: "",
      image: "",
      privacy: "public"
    };
  },
  methods: {
    async createPost() {
      const postData = {
        user_id: Number(this.userID),
        content: this.content,
        image: this.image || null,
        privacy: this.privacy
      };
      const res = await fetch("http://localhost:8080/posts", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(postData)
      });
      const data = await res.json();
      console.log("Created Post:", data);
      alert("Post created!");
    }
  }
};
</script>
