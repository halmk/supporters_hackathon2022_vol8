<template>
  <v-dialog :value="dialog" max-width="500" @click:outside.stop="clickOutside">
    <v-card tile>
      <v-card-title class="text-h7">
        {{ title }}
      </v-card-title>
      <v-list-item v-for="(comment, i) in showComments" :key="i" tile>
        <v-list-item-content>
          <v-list-item-title>
            <span class="font-weight-black">{{ comment.username }}</span>
            <span class="text--secondary text-subtitle-2">{{
              comment.formattedCreatedAt
            }}</span>
          </v-list-item-title>
          <v-list-item-subtitle>{{ comment.comment }}</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-card-actions class="mt-3">
        <v-text-field
          v-model="comment"
          label="Comment"
          required
          outlined
          dense
        ></v-text-field>
        <v-btn color="primary" class="mb-4 ml-3" @click="send">Send</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script>
export default {
  name: 'CommentDialog',
  props: {
    title: String,
    project: String,
    url: String,
    comments: Array,
    dialog: Boolean,
  },
  data: () => ({
    showComments: [],
    comment: '',
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
        const newComment = comment
        for (const user of users) {
          if (user.uuid === comment.sender_uuid) {
            newComment.username = user.name
          }
        }
        const createdAt = new Date(newComment.created_at)
        newComment.formattedCreatedAt = createdAt.toLocaleString()
        this.showComments.splice(index, 1, newComment)
      })
      this.showComments.sort((a, b) => {
        if (a.created_at < b.created_at) {
          return -1
        } else {
          return 1
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

    send() {
      this.$emit('clickSend', this.comment)
    },
  },
}
</script>
