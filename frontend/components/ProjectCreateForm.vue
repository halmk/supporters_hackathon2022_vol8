<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-text-field
            v-model="projectName"
            :rules="nameRules"
            label="Project name"
            required
          ></v-text-field>
        </v-col>
        <v-col cols="12">
          <v-btn
            :disabled="!valid || !loggedIn"
            color="primary"
            class="mr-4"
            @click="submit"
          >
            Create
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  name: 'ProjectCreateForm',
  data: () => ({
    projectName: '',
    valid: false,
    nameRules: [(v) => !!v || 'required'],
  }),
  computed: {
    ...mapGetters({
      loggedIn: 'auth/isLoggedIn',
    }),
  },
  methods: {
    submit() {
      // console.log('submit')
      const db = this.$fire.firestore
      const dbProjects = db.collection('projects')
      const uuid = this.$store.getters['auth/getUserUid']
      // console.log(uuid)
      dbProjects
        .add({
          creator_uuid: uuid,
          name: this.projectName,
        })
        .then((ref) => {
          // console.log('Add ID: ', ref.id)
          db.collection('participants')
            .add({
              user_uuid: uuid,
              project: this.projectName,
            })
            .then(() => {
              this.$router.push('/projects/' + this.projectName)
            })
        })
    },
  },
}
</script>
