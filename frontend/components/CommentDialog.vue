<template>
  <v-dialog :value="dialog" max-width="500" @click:outside.stop="clickOutside">
    <v-card tile>
      <v-card-title class="text-h5">
        {{ title }}
      </v-card-title>
      <v-list-item v-for="(comment, i) in comments" :key="i">
        <v-list-item-content>
          <v-list-item-title>{{ comment.username }}</v-list-item-title>
          <v-list-item-subtitle>{{ comment.content }}</v-list-item-subtitle>
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
    // title: String,
    // comments: Array,
    dialog: Boolean,
  },
  data: () => ({
    title: 'title1',
    comments: [
      {
        username: 'A',
        content: 'abcdefg',
      },
      {
        username: 'B',
        content: 'hijklmn',
      },
    ],
  }),
  computed: {},
  methods: {
    clickOutside() {
      this.$emit('clickOutside')
    },
  },
}
</script>
