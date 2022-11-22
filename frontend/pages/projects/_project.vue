<template>
  <div>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <span class="text-h3">{{ projectName }}</span>
      </v-col>
      <v-col cols="12">
        <BookmarkList :bookmarks="bookmarks" />
      </v-col>
    </v-row>
    <v-footer absolute class="grey lighten-5">
      <BookmarkSubmissionForm />
    </v-footer>
  </div>
</template>

<script>
export default {
  name: 'ProjectPage',
  data: () => ({
    projectName: '',
    categories: ['frontend'],
    bookmarks: [
      {
        url: 'http://example.com',
        thumbnailURL: 'http://example.com/thumbnail.png',
        title: 'example.com',
        content: 'abcdefg',
        project: 'project1',
        category: 'frontend',
      },
    ],
  }),
  mounted() {
    this.projectName = this.$route.params.project
    console.log(this.projectName)
  },
  methods: {},
}
</script>
