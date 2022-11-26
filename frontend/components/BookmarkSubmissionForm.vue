<template>
  <v-form v-model="valid">
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="bookmarkURL"
          :rules="nameRules"
          label="Bookmark URL"
          required
          dense
        ></v-text-field>
      </v-col>
      <v-col cols="12" sm="6">
        <v-btn
          :disabled="!valid"
          color="primary"
          class="mr-4"
          :loading="loading"
          @click="submit"
        >
          Submit
        </v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
export default {
  name: 'BookmarkSubmissionForm',
  props: {
    loading: Boolean,
  },
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
