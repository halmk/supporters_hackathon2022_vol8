<template>
  <v-dialog :value="dialog" max-width="500" @click:outside.stop="clickOutside">
    <v-card tile>
      <v-card-title class="text-h7">
        {{ title }}
      </v-card-title>
      <v-list-item v-for="(comment, i) in showComments" :key="i">
        <v-list-item-content>
          <v-list-item-title>{{ comment.username }}</v-list-item-title>
          <v-list-item-subtitle>{{ comment.comment }}</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-card-actions>
        <v-text-field></v-text-field>
        <v-btn>Send</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script>
export default {
  name: 'CommentDialog',
  props: {
    title: String,
    comments: Array,
    dialog: Boolean,
  },
  data: () => ({
    showComments: [],
  }),
  computed: {},
  watch: {
    async comments() {
      this.showComments = this.comments
      if (!this.comments) {
        return []
      }
      const users = await this.getUsers()
      console.log(users)
      this.showComments.forEach((comment, index) => {
        for (const user of users) {
          if (user.uuid === comment.sender_uuid) {
            const newComment = comment
            newComment.username = user.name
            this.showComments.splice(index, 1, newComment)
          }
        }
      })
    },
  },
  methods: {
    clickOutside() {
      this.$emit('clickOutside')
    },

    async getUsers() {
      try {
        const users = []
        for (const comment of this.comments) {
          // UUIDと一致するユーザをfirestoreから取得
          const db = this.$fire.firestore
          const snapshot = await db
            .collection('users')
            .doc(comment.sender_uuid)
            .get()
          const user = snapshot.data()
          console.log(user)
          users.push({ uuid: comment.sender_uuid, name: user.name })
        }
        return users
      } catch (error) {
        console.log(error)
      }
    },
  },
}
</script>
