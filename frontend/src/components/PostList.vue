<template>
  <div>
    <h2 class="text-xl font-bold mb-4">Posts</h2>

    <div v-for="post in posts" :key="post.id" class="mb-6 p-4 border rounded">
      <p class="font-semibold">Post #{{ post.id }} by User {{ post.user_id }}</p>
      <p class="mb-2">{{ post.content }}</p>
      <p class="text-sm text-gray-500">Visibility: {{ post.visibility }}</p>

      <div class="ml-4 mt-2">
        <p class="font-medium">Comments:</p>
        <div v-for="comment in post.comments" :key="comment.id" class="mb-1">
          - User {{ comment.user_id }}: {{ comment.content }}
        </div>

        <form @submit.prevent="submitComment(post.id)">
          <input v-model="newComments[post.id]" placeholder="Write a comment..." class="border p-1 mr-2"/>
          <button type="submit" class="bg-blue-500 text-white px-2 py-1 rounded">Comment</button>
        </form>
      </div>
    </div>

    <div class="mt-8 p-4 border rounded">
      <h3 class="font-semibold mb-2">Create New Post</h3>
      <textarea v-model="newPostContent" placeholder="Post content..." class="border p-2 w-full mb-2"></textarea>
      <select v-model="newPostVisibility" class="border p-2 mb-2">
        <option value="public">Public</option>
        <option value="private">Private</option>
        <option value="custom">Custom</option>
      </select>
      <input
        v-if="newPostVisibility === 'custom'"
        v-model="customUserIds"
        placeholder="User IDs comma separated (e.g. 2,3,5)"
        class="border p-2 w-full mb-2"
      />
      <button @click="submitPost" class="bg-green-500 text-white px-4 py-2 rounded">Post</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const posts = ref([])
const newComments = ref({})
const newPostContent = ref('')
const newPostVisibility = ref('public')
const customUserIds = ref('')
const userId = 1

const fetchPosts = async () => {
  const res = await fetch(`http://localhost:8080/api/get_posts?user_id=${userId}`)
  posts.value = await res.json()
}

onMounted(fetchPosts)

const submitComment = async (postId) => {
  const content = newComments.value[postId]
  if (!content) return

  await fetch('http://localhost:8080/api/comments', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      user_id: userId,
      post_id: postId,
      content
    })
  })

  newComments.value[postId] = ''
  fetchPosts()
}

const submitPost = async () => {
  if (!newPostContent.value) return

  const body = {
    user_id: userId,
    content: newPostContent.value,
    visibility: newPostVisibility.value
  }

  if (newPostVisibility.value === 'custom') {
    body.audience = customUserIds.value.split(',').map(id => parseInt(id.trim()))
  }

  await fetch('http://localhost:8080/api/posts', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })

  newPostContent.value = ''
  customUserIds.value = ''
  newPostVisibility.value = 'public'
  fetchPosts()
}
</script>

<style scoped>
</style>