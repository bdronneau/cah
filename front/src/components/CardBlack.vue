<template>
  <v-card
    class="mx-auto"
    max-width="344"
  >
    <v-card-text>
      <div>Current black card</div>
      <p class="display-1 text--primary">
        {{ room.raw.current_card.name }}
      </p>
      <p>1 carte</p>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'CardBlack',
  computed: {
    room () {
      return this.$store.getters.getRoom
    }
  },
};
</script>
