<template>
  <v-card
    class="mx-auto"
    max-width="344"
    outlined
  >
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4">Informations</div>
        <v-list-item-title class="headline mb-1">{{ room.name }} </v-list-item-title>
        <v-list-item-subtitle>Turn: {{ room.turn }}</v-list-item-subtitle>
      </v-list-item-content>

      <v-list-item-avatar
        tile
        size="80"
        color="grey"
      ></v-list-item-avatar>
    </v-list-item>

    <v-card-actions>
      <v-btn :disabled="!isRefreshable" text @click="refreshRoom">Refresh</v-btn>
      <v-btn disabled text>{{ room.status }}</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'LifeStats',
  computed: {
    room () {
      return this.$store.getters.getRoomInformations
    },
    isRefreshable () {
      return !(this.$store.getters.getRoomInformations.status === "finished" )
    }
  },
  methods: {
    refreshRoom: function() {
      this.$store.dispatch("getRoomData")
    }
  },
};
</script>
