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
            :disabled="!valid"
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
// import firebase from '@nuxtjs/firebase'
export default {
  name: 'ProjectCreateForm',
  data: () => ({
    projectName: '',
    valid: false,
    nameRules: [(v) => !!v || 'required'],
  }),
  computed: {},
  methods: {
    submit() {
      console.log('submit')
      const db = this.$fire.firestore
      console.log('submit')
      const dbProjects = db.collection('projects')
      console.log('submit')
      const uuid = this.$store.getters['auth/getUserUid']
      console.log(uuid)
      dbProjects
        .add({
          creator_uuid: uuid,
          name: this.projectName,
        })
        .then((ref) => {
          console.log('Add ID: ', ref.id)
        })
    },
  },
}
</script>
