<template>
  <v-card>
    <v-card-title>
      {{ bookmark.title }}
    </v-card-title>
    <v-card-subtitle>
      {{ bookmark.url }}
    </v-card-subtitle>
    <v-card-text>
      {{ bookmark.content }}
    </v-card-text>
    <v-card-actions>
      <v-btn text outlined @click="dialog = true">コメント</v-btn>
    </v-card-actions>
    <CommentDialog :dialog="dialog" @clickOutside="dialog = false" />
  </v-card>
</template>
<script>
export default {
  name: 'BookmarkCard',
  props: {
    bookmark: Object,
  },
  data: () => ({
    dialog: false,
  }),
  computed: {},
  methods: {},
}
</script>
