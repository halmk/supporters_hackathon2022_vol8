<template>
  <div class="grey lighten-4 pa-2">
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
    unset: null,
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
    // this.getBookmarks()
    this.syncBookmarks()
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
      console.log('set bookmarks:', bookmarks)
      this.bookmarks = []
      for (const bookmark of bookmarks) {
        this.bookmarks.push(bookmark)
      }
      // this.bookmarks = bookmarks
      this.categories = this.categories.concat(this.getCategories())
      await this.getComments()
    },

    async submitBookmark(bookmarkURL) {
      this.submitLoading = true
      try {
        // APIサーバにURLを送信してメタデータを取得する
        // const apiURL = `${process.env.API_ENDPOINT}/metadata`
        const apiURL = `${process.env.API_ENDPOINT}`
        // const data = "{\"url\": \""+bookmarkURL+"\"}"
        // console.log(data)
        // const response = await axios.post(apiURL, data)
        const response = await axios.post(apiURL, {
          // data,
          url: bookmarkURL,
          // headers: {
          //   'Content-Type': 'application/json',
          //   // "Access-Control-Allow-Origin": "*",
          // },
        })
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
        // await this.getBookmarks()
      } catch (error) {
        console.log(error)
      }
      this.submitLoading = false
    },

    async getComments() {
      try {
        // プロジェクトと一致するコメントをfirestoreから取得
        const db = this.$fire.firestore
        const snapshot = await db
          .collection('comments')
          .where('project', '==', this.projectName)
          .get()
        const comments = snapshot.docs.map((doc) => {
          return doc.data()
        })

        // 一致するbookmarkにcommentsを振り分ける
        this.bookmarks.forEach((bookmark) => {
          const filteredComments = comments.filter(
            (comment) => comment.bookmark_url === bookmark.url
          )
          bookmark.comments = filteredComments
        })
        console.log(this.bookmarks)
        this.bookmarks.splice()
      } catch (error) {
        console.log(error)
      }
    },

    syncBookmarks() {
      // 対象のquery
      const db = this.$fire.firestore
      const query = db
        .collection('bookmarks')
        .where('project', '==', this.projectName)

      // リッスンの開始
      this.unset = query.onSnapshot((snapshot) => {
        snapshot.docChanges().forEach((change) => {
          // 変更のタイプ
          console.info(`*** change: type=${change.type}`)
          // 対象のドキュメント。変更の場合は、変更後
          console.info(
            `*** change: data=${JSON.stringify(change.doc.data(), null, 2)}`
          )
          // refに対する変更後のindex。削除時は-1
          console.info(`*** change: newIndex=${change.newIndex}`)
          // refに対する変更前のindex。追加時は-1
          console.info(`*** change: oldIndex=${change.oldIndex}`)

          this.getBookmarks()
          if (change.type === 'added') {
            // 追加時
          } else if (change.type === 'modified') {
            // 変更時
          } else if (change.type === 'removed') {
            // 削除時
          }
        })
      })
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
