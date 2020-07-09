<template>
  <v-container>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="cards"
      :single-select=true
      item-key="name"
      show-select
      hide-default-footer
      class="elevation-1"
    >
    </v-data-table>
    <v-btn color="primary" :disabled="selected.length === 0" @click="sendCard">Confirm</v-btn>
    <v-btn text>Reset</v-btn>
  </v-container>
</template>

<script>
export default {
  name: 'UserHand',
  props: {
    cards: Array
  },
  data() {
    return {
      selected: [],
      headers: [
        { text: 'id', value: 'id' },
        { text: 'Name', value: 'name' },
      ],
    }
  },
  methods: {
    sendCard: function() {
      this.$emit("played", this.selected[0]);
    }
  },
};
</script>
