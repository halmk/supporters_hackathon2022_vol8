<template>
  <div>
    <div v-for="(bookmark, i) in bookmarks" :key="i" class="mt-3">
      <BookmarkCard :bookmark="bookmark" />
    </div>
  </div>
</template>
<script>
export default {
  name: 'BookmarkList',
  props: {
    bookmarks: Array,
  },
  data: () => ({}),
  computed: {},
  methods: {},
}
</script>
