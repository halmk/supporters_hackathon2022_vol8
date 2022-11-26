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
    projects: [
      {
        name: 'project1',
        participants: ['taro', 'hanako'],
      },
    ],
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

          db.collection('participants')
            .where('project', '==', project)
            .get()
            .then((snapshot) => {
              snapshot.forEach((doc2) => {
                // テーブル表示
                const usersUuid = doc2.data().userUuid
                // console.log(users_uuid)
                const users = []
                for (let i = 0; i < usersUuid.length; i++) {
                  db.collection('users')
                    .where('uuid', '==', usersUuid[i])
                    .get()
                    .then((userSnapshot) => {
                      userSnapshot.forEach((doc3) => {
                        const name = doc3.data().name
                        // console.log(name)
                        users.push(name)
                      })
                    })
                }
                this.projects.push({ name: project, participants: users })
              })
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
