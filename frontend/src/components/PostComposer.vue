<script setup>
import { reactive } from 'vue'
import { useForumStore } from '@/stores/forumStore'

const props = defineProps({
  boards: { type: Array, default: () => [] },
  topics: { type: Array, default: () => [] }
})
const emit = defineEmits(['created'])
const store = useForumStore()
const form = reactive({
  title: '',
  content: '',
  boardId: '',
  topicId: '',
  username: '',
  displayName: ''
})

const reset = () => {
  Object.assign(form, {
    title: '',
    content: '',
    boardId: '',
    topicId: '',
    username: '',
    displayName: ''
  })
}

const submit = async () => {
  const payload = {
    title: form.title,
    content: form.content,
    boardId: Number(form.boardId),
    topicId: Number(form.topicId),
    username: form.username,
    displayName: form.displayName
  }
  const post = await store.createPost(payload)
  emit('created', post)
  reset()
}
</script>

<template>
  <div class="panel">
    <div class="panel__header">
      <h3>发布新帖</h3>
    </div>
    <div class="panel__body">
      <input v-model="form.title" placeholder="标题" />
      <div class="form-grid">
        <select v-model="form.boardId">
          <option value="" disabled>选择板块</option>
          <option v-for="board in boards" :key="board.id" :value="board.id">{{ board.name }}</option>
        </select>
        <select v-model="form.topicId">
          <option value="" disabled>选择话题</option>
          <option v-for="topic in topics" :key="topic.id" :value="topic.id">{{ topic.title }}</option>
        </select>
      </div>
      <div class="form-grid">
        <input v-model="form.username" placeholder="用户名" />
        <input v-model="form.displayName" placeholder="显示名" />
      </div>
      <textarea v-model="form.content" rows="4" placeholder="内容" />
      <button
        :disabled="!(form.title && form.content && form.boardId && form.topicId && form.username && form.displayName)"
        @click="submit"
      >发布</button>
    </div>
  </div>
</template>