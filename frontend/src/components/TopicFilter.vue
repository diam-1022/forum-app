<script setup>
const props = defineProps({
  topics: { type: Array, default: () => [] },
  selectedId: { type: Number, default: null }
})
const emit = defineEmits(['select'])
const select = (id) => emit('select', id)
</script>

<template>
  <div class="panel">
    <div class="panel__header">
      <h3>话题</h3>
      <small>按话题筛选帖子</small>
    </div>
    <div class="panel__body chips">
      <button
        class="chip"
        :class="{ 'chip--active': !selectedId }"
        @click="select(null)"
      >全部</button>
      <button
        v-for="topic in topics"
        :key="topic.id"
        class="chip"
        :class="{ 'chip--active': selectedId === topic.id }"
        @click="select(topic.id)"
      >
        {{ topic.title }}
      </button>
    </div>
  </div>
</template>