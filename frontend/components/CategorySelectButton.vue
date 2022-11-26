<template>
  <v-overflow-btn
    :items="categories"
    label="Select Category"
    hide-details
    dense
    class="pa-0"
    @input="input"
  ></v-overflow-btn>
</template>

<script>
export default {
  name: 'CategorySelectButton',
  props: {
    categories: Array,
  },
  data: () => ({}),
  computed: {},
  methods: {
    input(category) {
      this.$emit('input', category)
    },
  },
}
</script>
