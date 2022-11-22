<template>
  <div>
    <div v-for="(bookmark, i) in bookmarks" :key="i" class="mt-2">
      <BookmarkCard :bookmark="bookmark" />
    </div>
  </div>
</template>
<script>
export default {
  name: 'BookmarkList',
  props: {
    // bookmarks: Array,
  },
  data: () => ({
    bookmarks: [
      {
        url: 'http://example.com',
        thumbnailURL: 'http://example.com/thumbnail.png',
        title: 'example.com',
        content: 'abcdefg',
        cratedAt: '2022-11-22T9:00:00',
        project: 'project1',
        category: 'frontend',
        sender_uuid: 'abc',
      },
      {
        url: 'http://example.com',
        thumbnailURL: 'http://example.com/thumbnail.png',
        title: 'example.com',
        content: 'abcdefg',
        cratedAt: '2022-11-22T9:00:00',
        project: 'project1',
        category: 'frontend',
        sender_uuid: 'xyz',
      },
    ],
  }),
  computed: {},
  methods: {},
}
</script>
