<template>
  <v-container>
    <v-row>
      <v-col>
        <RoomInformations :id=roomID />
      </v-col>
      <v-col>
        <div class="my-2">
          <v-btn v-if="!isStarted" :disabled="isStarted || (!isInRoom && !isStarted)" v-on:click="startRoom" color="primary">Start party</v-btn>
          <v-btn v-if="isStarted" :disabled="!isInRoom" v-on:click="startRoom" color="primary">Enter in party</v-btn>
        </div>
        <div class="my-2">
          <v-btn v-if="userID" :disabled="isInRoom" v-on:click="joinRoom" color="primary">Join room</v-btn>
          <v-dialog v-model="dialog" persistent max-width="600px">
            <template v-slot:activator="{ on }">
              <v-btn v-if="userID === null" :disabled="disableCreateUser || isStarted" color="primary" v-on="on">New user</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">New user</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <v-text-field v-model="newUser.name" label="Name*" required></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
                <small>*indicates required field</small>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="dialog = false">Aborted</v-btn>
                <v-btn color="blue darken-1" text v-on:click="createUser">Save</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </div>
      </v-col>
    </v-row>

    <v-row>
      <v-col
        xs12
        sm12
        md12
      >
        <UsersLobby :id=roomID />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import RoomInformations from '@/components/RoomInformations.vue';
  import UsersLobby from '@/components/UsersLobby.vue';

  import { addUserInRoom, getUsersLobby, retrieveRoomStatus, postNewUser, postStartRoom } from '@/api/main';

  export default {
    components: {
      RoomInformations,
      UsersLobby,
    },
    data() {
      return {
        roomID: null,
        newUser: {
          name: null,
        },
        dialog: false,
        disableCreateUser: false,
        room: {},
      }
    },
    created() {
      this.roomID = this.$route.params.id
      this.$store.commit('setRoom', { id: this.roomID })
      if (this.$store.getters.userID) {
        this.disableCreateUser = true
      }
    },
    computed: {
      users () {
        return this.$store.getters.users
      },
      userID () {
        return this.$store.getters.userID
      },
      isStarted () {
        return this.$store.getters.getRoomInformations.status === "started"
      },
      isInRoom () {
        return this.$store.getters.isInRoom
      }
    },
    methods: {
      poolingData() {
        this.intervalID = setInterval(() => {
            this.$store.dispatch("getRoomData")
          },
          10000)
      },
      joinRoom: function() {
        addUserInRoom(this.$store.getters.userID, this.roomID).then(() => {
          return getUsersLobby(this.roomID)
        }).then(data => {
          this.$store.commit("setUsers", data)
          this.poolingData()
        })
      },
      createUser: function() {
        postNewUser(this.newUser.name).then(resp => {
          this.$store.commit('setUserID', resp.data.id)
          return addUserInRoom(this.$store.getters.userID, this.roomID)
        }).then(() => {
          this.poolingData()
          this.dialog = false
          return getUsersLobby(this.roomID)
        }).then(data => {
          this.$store.commit("setUsers", data)
        })
      },
      startRoom: function() {
        this.$store.dispatch("getRoomData")

        if (this.isStarted) {
          this.$router.push({ name: 'Room', params: { id: this.roomID }})
        } else {
          postStartRoom(this.roomID).then(() => {
            this.$router.push({ name: 'Room', params: { id: this.roomID }})
          })
        }
      },
      enterRoom: function() {
        this.$router.push({ name: 'Room', params: { id: this.roomID }})
      },
    },
    beforeDestroy () {
      clearInterval(this.intervalID)
    },
    mounted() {
      this.$store.dispatch("getRoomData")
      if (this.$store.getters.userID) {
        this.disableCreateUser = true
      }

      retrieveRoomStatus(this.roomID).then(data => {
        this.room = data
        this.$log.debug("end get room status")
        return getUsersLobby(this.roomID)
      }).then(data => {
        this.$store.commit("setUsers", data)

        if (this.isInRoom) {
          if (this.isStarted) {
            this.$router.push({ name: 'Room', params: { id: this.roomID }})
          } else {
            this.poolingData()
          }
        }
      })
    },
  }
</script>
