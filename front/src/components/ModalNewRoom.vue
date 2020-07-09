<template>
  <v-dialog v-model="dialog" persistent max-width="600px">
    <template v-slot:activator="{ on }">
      <v-btn small color="primary" v-on="on">New room</v-btn>
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">New room</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field  v-model="room.description" label="Name*" required></v-text-field>
            </v-col>
          </v-row>
        </v-container>
        <small>*indicates required field</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="newRoomAbort">Aborted</v-btn>
        <v-btn color="blue darken-1" text @click="newRoomSave">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
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
