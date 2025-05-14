<template>
  <div class="post-container">
    <h2>Create Post</h2>
    <form @submit.prevent="createPost">
      <textarea v-model="title" placeholder="Post title" required></textarea>
      <textarea v-model="content" placeholder="Post content" required></textarea>
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
      title: "",
      content: "",
      image: "",
      privacy: "public",
    };
  },
  methods: {
    async createPost() {
      const postData = {
        user_id: Number(this.userID),
        title: this.title,
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

<style scoped>
.post-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 2rem auto;
  max-width: 500px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 100%;
}

textarea,
select,
button {
  width: 100%;
  padding: 0.75rem;
  border-radius: 0.5rem;
  border: 1px solid #ccc;
  font-size: 1rem;
}

button {
  background-color: #007bff;
  color: #fff;
  border: none;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #0056b3;
}
</style>