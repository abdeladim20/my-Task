<template>
  <div>
    <h1>📝 My Social Network</h1>
    <CreatePost />
    <PostList />
  </div>
</template>

<script>
import CreatePost from "./components/CreatePost.vue";
import PostList from "./components/PostList.vue";

export default {
  components: {
    CreatePost,
    PostList
  }
};
</script>
