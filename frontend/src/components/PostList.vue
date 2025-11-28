<script setup>
const props = defineProps({
  posts: { type: Array, default: () => [] },
  loading: { type: Boolean, default: false },
  error: { type: String, default: '' }
})
const emit = defineEmits(['select'])
const select = (post) => emit('select', post)
</script>

<template>
  <div class="panel">
    <div class="panel__header">
      <h3>帖子列表</h3>
      <small v-if="posts.length">共 {{ posts.length }} 条</small>
    </div>
    <div class="panel__body">
      <p v-if="loading">加载中...</p>
      <p v-else-if="error" class="text-error">{{ error }}</p>
      <p v-else-if="!posts.length">暂无帖子，快来发布第一条内容吧！</p>
      <ul class="post-list" v-else>
        <li
          v-for="post in posts"
          :key="post.id"
          class="post-card"
          @click="select(post)"
        >
          <header>
            <h4>{{ post.title }}</h4>
            <small>
              {{ post.board?.name }} / {{ post.topic?.title }} / {{ post.author?.displayName || '匿名用户' }}
            </small>
          </header>
          <p>
            {{ (post.content || '').slice(0, 120) }}
            <span v-if="post.content && post.content.length > 120">...</span>
          </p>
          <footer>
            <small>{{ new Date(post.updatedAt || post.createdAt).toLocaleString() }}</small>
            <small>{{ post.comments?.length || 0 }} 评论</small>
          </footer>
        </li>
      </ul>
    </div>
  </div>
</template>