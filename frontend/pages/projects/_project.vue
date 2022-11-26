<template>
  <div>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <span class="text-h3">{{ projectName }}</span>
      </v-col>
      <v-col cols="12" class="mt-2">
        <CategoryCreateForm @submit="newCategory" />
        <BookmarkSubmissionForm
          :loading="submitLoading"
          @submit="submitBookmark"
        />
        <v-row>
          <v-col cols="12" sm="6">
            <CategorySelectButton
              :categories="categories"
              @input="setCategory"
            />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <BookmarkList :bookmarks="bookmarks" />
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'ProjectPage',
  data: () => ({
    projectName: '',
    categories: ['General'],
    selectedCategory: 'General',
    bookmarks: [],
    submitLoading: false,
  }),
  mounted() {
    this.projectName = this.$route.params.project
    console.log(this.projectName)
  },
}
</script>
