<template>
  <div class="post-container">
    <h2>Create Post</h2>
    <form @submit.prevent="createPost">
      <textarea v-model="title" placeholder="Post title" required></textarea>
      <textarea v-model="content" placeholder="Post content" required></textarea>
      <input type="file" @change="onFileChange" />
      <select v-model="privacy">
        <option value="public">Public</option>
        <option value="private">Private</option>
      </select>
      <button type="submit">Post</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";

const emit = defineEmits(["post-created"]);

const title = ref("");
const content = ref("");
const imageFile = ref(null);
const privacy = ref("public");

const onFileChange = (event) => {
  const files = event.target.files;
  if (files.length > 0) {
    imageFile.value = files[0];
  } else {
    imageFile.value = null;
  }
};

const createPost = async () => {
  const formData = new FormData();
  formData.append("user_id", 1);
  formData.append("title", title.value);
  formData.append("content", content.value);
  formData.append("privacy", privacy.value);

  if (imageFile.value) {
    formData.append("image", imageFile.value);
  }

  const res = await fetch("http://localhost:8080/posts", {
    method: "POST",
    body: formData,
  });

  if (!res.ok) {
    const error = await res.text();
    console.error("Post creation failed:", error);
    alert("Post creation failed");
    return;
  }

  const data = await res.json();
  emit("post-created", data);

  // Reset form
  title.value = "";
  content.value = "";
  imageFile.value = null;
  privacy.value = "public";

  // Clear the file input element
  document.querySelector('input[type="file"]').value = null;
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