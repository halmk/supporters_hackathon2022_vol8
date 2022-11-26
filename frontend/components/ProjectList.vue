<template>
  <v-card tile>
    <v-list-item
      v-for="(project, i) in projects"
      :key="i"
      :to="project.name"
      append
      router
      exact
    >
      <v-list-item-content>
        <v-list-item-title>{{ project.name }}</v-list-item-title>
        <v-list-item-subtitle
          >参加者：{{ project.participants.toString() }}</v-list-item-subtitle
        >
      </v-list-item-content>
    </v-list-item>
  </v-card>
</template>
<script>
export default {
  name: 'ProjectList',
  props: {
    // projects: Array,
  },
  data: () => ({
    projects: [],
  }),
  computed: {},
  created: function () {
    // console.log('created')
    // console.log('get')
    const db = this.$fire.firestore
    const uuid = this.$store.getters['auth/getUserUid']
    // データ取得
    db.collection('projects')
      .where('creator_uuid', '==', uuid)
      .get()
      .then((querySnapshot) => {
        querySnapshot.forEach((doc) => {
          // テーブル表示
          const project = doc.data().name
          // console.log(project)
          const users = []
          db.collection('participants')
            .where('project', '==', project)
            .get()
            .then((snapshot) => {
              snapshot.forEach((doc2) => {
                // テーブル表示
                const userUuid = doc2.data().user_uuid
                // console.log(users_uuid)

                db.collection('users')
                  .doc(userUuid)
                  .get()
                  .then((doc) => {
                    users.push(doc.data().name)
                  })
                  .catch((error) => {
                    console.log(error)
                  })
              })
            })
            .then(() => {
              this.projects.push({ name: project, participants: users })
            })
            .catch((error) => {
              // console.log(project)
              console.log(error)
            })
        })
      })
      .catch((error) => {
        console.log(error)
      })
  },
  methods: {},
}
</script>
