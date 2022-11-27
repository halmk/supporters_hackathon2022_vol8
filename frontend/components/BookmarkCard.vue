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
    <CommentDialog
      :title="bookmark.title"
      :project="bookmark.project"
      :url="bookmark.url"
      :comments="bookmark.comments"
      :dialog="dialog"
      @clickOutside="dialog = false"
      @clickSend="sendComment"
    />
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
  methods: {
    async sendComment(comment) {
      try {
        // コメントの情報をFirestoreに保存する
        const db = this.$fire.firestore
        const userUUID = this.$store.getters['auth/getUserUid']
        const createdAt = new Date()

        await db.collection('comments').add({
          comment,
          category: this.bookmark.category,
          created_at: createdAt.toISOString(),
          project: this.bookmark.project,
          sender_uuid: userUUID,
          bookmark_url: this.bookmark.url,
        })
      } catch (error) {
        console.log(error)
      }
    },
  },
}
</script>
