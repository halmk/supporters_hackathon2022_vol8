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
        <BookmarkList :bookmarks="filteredBookmarks" />
      </v-col>
    </v-row>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ProjectPage',
  data: () => ({
    projectName: '',
    categories: ['General'],
    selectedCategory: 'General',
    bookmarks: [],
    submitLoading: false,
  }),
  computed: {
    filteredBookmarks() {
      return this.bookmarks.filter(
        (bookmark) => bookmark.category === this.selectedCategory
      )
    },
  },
  created() {
    this.projectName = this.$route.params.project
    console.log(this.projectName)
    this.getBookmarks()
  },
  methods: {
    async getBookmarks() {
      const db = this.$fire.firestore
      const snapshot = await db
        .collection('bookmarks')
        .where('project', '==', this.projectName)
        .get()
      const bookmarks = snapshot.docs.map((doc) => {
        return doc.data()
      })
      bookmarks.sort((a, b) => {
        if (a.created_at > b.created_at) {
          return -1
        } else {
          return 1
        }
      })
      console.log(bookmarks)
      this.bookmarks = bookmarks
      this.categories = this.categories.concat(this.getCategories())
    },

    async submitBookmark(bookmarkURL) {
      this.submitLoading = true
      try {
        // APIサーバにURLを送信してメタデータを取得する
        const apiURL = 'http://localhost:5000/metadata'
        const params = { url: bookmarkURL }
        console.log(apiURL, params)
        const response = await axios.get(apiURL, { params })
        const metadata = response.data
        const createdAt = new Date()
        console.log(metadata)

        // メタデータを含めたブックマークの情報をFirestoreに保存する
        const db = this.$fire.firestore
        const userUUID = this.$store.getters['auth/getUserUid']
        await db.collection('bookmarks').add({
          category: this.selectedCategory,
          content: metadata.description,
          created_at: createdAt.toISOString(),
          good_count: 0,
          project: this.projectName,
          sender_uuid: userUUID,
          thumbnail_url: metadata.image,
          title: metadata.title,
          url: bookmarkURL,
        })
        await this.getBookmarks()
      } catch (error) {
        console.log(error)
      }
      this.submitLoading = false
    },

    getCategories() {
      const categorySet = new Set(
        this.bookmarks.map((bookmark) => bookmark.category)
      )
      return [...categorySet]
    },

    setCategory(category) {
      this.selectedCategory = category
    },

    newCategory(category) {
      this.categories.push(category)
    },
  },
}
</script>
