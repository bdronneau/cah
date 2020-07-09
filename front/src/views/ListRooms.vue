<template>
  <v-card>
    <v-card-title>
      <div class="my-2">
        <v-btn small @click="getRooms()" color="primary">Refresh</v-btn>
      </div>

    <v-spacer></v-spacer>
      <div class="my-2">
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
      </div>
    <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="mdi-search"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>

    <v-data-table
      :headers="headers"
      :items="rooms"
      :search="search"
      item-key="name"
      class="elevation-1"
    >
    <template v-slot:top>
      <v-toolbar flat>
        <v-toolbar-title></v-toolbar-title>
        <v-spacer></v-spacer>
      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="joinRoom(item.name)"
      >
        mdi-bank-transfer-in
      </v-icon>
    </template>
  </v-data-table>
</v-card>

</template>
<script>

  import { listRooms, postNewRoom } from '@/api/main';

  export default {
    data () {
      return {
        dialog: false,
        search: '',
        expanded: [],
        headers: [
          {
            text: 'Room name',
            align: 'start',
            sortable: false,
            value: 'name',
          },
          { text: 'Description', value: 'description' },
          { text: 'Status', value: 'status' },
          { text: 'Current turn', value: 'turn' },
          { text: 'Actions', value: 'actions', sortable: false },
        ],
        rooms: [],
        room: {
          description: ""
        },
      }
    },
    mounted() {
      this.getRooms()
    },
    methods: {
      getRooms: function () {
        listRooms().then(data => {
          this.rooms = data
        })
      },
      joinRoom: function (id) {
        this.$router.push({ name: 'LobbyRoom', params: { id: id }})
      },
      newRoomAbort: function () {
        this.dialog = false
        this.room.description = ""
      },
      newRoomSave: function () {
        postNewRoom(this.room.description).then(() => {
          this.room.description = ""
          this.dialog = false
          return listRooms()
        }).then(data => {
          this.rooms = data
        })
      }
    }
  }
</script>
