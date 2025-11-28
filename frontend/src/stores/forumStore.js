import { defineStore } from 'pinia'
import client from '@/api/http'

export const useForumStore = defineStore('forum', {
  state: () => ({
    boards: [],
    topics: [],
    posts: [],
    activePost: null,
    loading: false,
    error: null,
  }),
  actions: {
    async fetchBoards() {
      const { data } = await client.get('/boards')
      this.boards = data
    },
    async fetchTopics() {
      const { data } = await client.get('/topics')
      this.topics = data
    },
    async fetchPosts(params = {}) {
      this.loading = true
      this.error = null
      try {
        const { data } = await client.get('/posts', { params })
        this.posts = data
      } catch (error) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    },
    async fetchPost(id) {
      this.loading = true
      this.error = null
      try {
        const { data } = await client.get('/posts/' + id)
        this.activePost = data
        return data
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },
    async createPost(payload) {
      const { data } = await client.post('/posts', payload)
      this.posts.unshift(data)
      return data
    },
    async createComment(postId, payload) {
      const { data } = await client.post('/posts/' + postId + '/comments', payload)
      if (this.activePost && this.activePost.id === postId) {
        this.activePost.comments = [data, ...(this.activePost.comments || [])]
      }
      return data
    }
  }
})


