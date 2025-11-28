<script setup>
const props = defineProps({
  boards: { type: Array, default: () => [] },
  selectedId: { type: Number, default: null }
})
const emit = defineEmits(['select'])
const select = (id) => emit('select', id)
</script>

<template>
  <div class="panel">
    <div class="panel__header">
      <h3>板块</h3>
      <small>选择不同板块查看帖子</small>
    </div>
    <div class="panel__body chips">
      <button
        class="chip"
        :class="{ 'chip--active': !selectedId }"
        @click="select(null)"
      >
        全部
      </button>
      <button
        v-for="board in boards"
        :key="board.id"
        class="chip"
        :class="{ 'chip--active': selectedId === board.id }"
        @click="select(board.id)"
      >
        {{ board.name }}
      </button>
    </div>
  </div>
</template>