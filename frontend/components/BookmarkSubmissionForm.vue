<template>
  <v-form v-model="valid">
    <v-text-field
      v-model="bookmarkURL"
      :rules="nameRules"
      label="bookmark url"
      required
    ></v-text-field>
    <v-btn :disabled="!valid" color="primary" class="mr-4" @click="submit">
      Submit
    </v-btn>
  </v-form>
</template>
<script>
export default {
  name: 'BookmarkSubmissionForm',
  data: () => ({
    bookmarkURL: '',
    valid: false,
    nameRules: [(v) => !!v || 'required'],
  }),
  computed: {},
  methods: {
    submit() {},
  },
}
</script>
