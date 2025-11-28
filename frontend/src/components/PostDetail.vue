<script setup>
import { reactive, watch } from 'vue'
import { useForumStore } from '@/stores/forumStore'

const props = defineProps({
  post: { type: Object, default: null }
})
const emit = defineEmits(['comment-created'])
const store = useForumStore()
const form = reactive({
  content: '',
  username: '',
  displayName: ''
})

watch(
  () => props.post?.id,
  () => {
    form.content = ''
  }
)

const submit = async () => {
  if (!props.post?.id) return
  await store.createComment(props.post.id, { ...form })
  emit('comment-created')
  form.content = ''
}
</script>

<template>
  <div class="panel post-detail">
    <div class="panel__header">
      <h3>帖子详情</h3>
    </div>
    <div class="panel__body" v-if="post">
      <article>
        <header>
          <h2>{{ post.title }}</h2>
          <small>
            {{ post.board?.name }} / {{ post.topic?.title }} / {{ post.author?.displayName || '匿名' }}
          </small>
        </header>
        <p class="post-detail__content">{{ post.content }}</p>
      </article>
      <section class="comments">
        <h4>评论</h4>
        <p v-if="!post.comments?.length">暂无评论</p>
        <ul>
          <li v-for="comment in post.comments" :key="comment.id">
            <p>{{ comment.content }}</p>
            <small>{{ comment.author?.displayName || '匿名' }} / {{ new Date(comment.createdAt).toLocaleString() }}</small>
          </li>
        </ul>
      </section>
      <section class="comment-form">
        <h4>发布评论</h4>
        <div class="form-grid">
          <input v-model="form.username" placeholder="用户名(username)" />
          <input v-model="form.displayName" placeholder="显示名" />
        </div>
        <textarea v-model="form.content" rows="3" placeholder="写点什么..."></textarea>
        <button :disabled="!(form.content && form.username && form.displayName)" @click="submit">提交评论</button>
      </section>
    </div>
    <div class="panel__body" v-else>
      <p>选择一个帖子查看详情与评论。</p>
    </div>
  </div>
</template>