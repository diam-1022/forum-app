<script setup>
import { onMounted, reactive } from 'vue'
import { useForumStore } from '@/stores/forumStore'
import BoardList from '@/components/BoardList.vue'
import TopicFilter from '@/components/TopicFilter.vue'
import PostList from '@/components/PostList.vue'
import PostComposer from '@/components/PostComposer.vue'
import PostDetail from '@/components/PostDetail.vue'

const store = useForumStore()
const filters = reactive({ boardId: null, topicId: null })

const applyFilters = () => {
  const params = {}
  if (filters.boardId) params.boardId = filters.boardId
  if (filters.topicId) params.topicId = filters.topicId
  return store.fetchPosts(params)
}

const handleBoardSelect = (id) => {
  filters.boardId = id
  applyFilters()
}

const handleTopicSelect = (id) => {
  filters.topicId = id
  applyFilters()
}

const handlePostSelect = (post) => {
  if (post?.id) {
    store.fetchPost(post.id)
  }
}

const handlePostCreated = (post) => {
  filters.boardId = post.boardId
  filters.topicId = post.topicId
  applyFilters()
  store.fetchPost(post.id)
}

const refreshActivePost = () => {
  if (store.activePost?.id) {
    store.fetchPost(store.activePost.id)
  }
}

onMounted(async () => {
  await Promise.all([store.fetchBoards(), store.fetchTopics()])
  applyFilters()
})
</script>

<template>
  <div class="app-shell">
    <header class="app-header">
      <div>
        <h1>反正是个Forum</h1>
        <p>BangDream</p>
      </div>
      <span class="badge">Full stack demo</span>
    </header>
    <main class="layout">
      <aside class="sidebar">
        <BoardList :boards="store.boards" :selected-id="filters.boardId" @select="handleBoardSelect" />
        <TopicFilter :topics="store.topics" :selected-id="filters.topicId" @select="handleTopicSelect" />
        <PostComposer :boards="store.boards" :topics="store.topics" @created="handlePostCreated" />
      </aside>
      <section class="content">
        <PostList :posts="store.posts" :loading="store.loading" :error="store.error" @select="handlePostSelect" />
        <PostDetail :post="store.activePost" @comment-created="refreshActivePost" />
      </section>
    </main>
  </div>
</template>
