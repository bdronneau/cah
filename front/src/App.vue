<template>
  <div id="app">
    <v-app>
      <v-app-bar
        app
        clipped-right
        color="primary"
        dark
      >
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title>CAH - Card Against Humanity</v-toolbar-title>
        <v-spacer></v-spacer>
      </v-app-bar>

      <v-navigation-drawer
        v-model="drawer"
        app
      >
        <v-list dense>
          <v-list-item>
            <v-list-item-action>
              <v-icon>mdi-home</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Home</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-action>
              <v-icon>mdi-pencil</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Create room</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item :disabled="!hasUserID" @click.stop="displayUserInformations" >
            <v-list-item-action>
              <v-icon>mdi-account</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>User information</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-content>
        <v-container>
          <router-view />
        </v-container>
      </v-content>

      <v-navigation-drawer
        v-model="drawerUserInfo"
        fixed
        left
        temporary
      >
          <v-list-item>
            <v-list-item-action>
              <v-icon>mdi-information-outline</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>ID: {{ user.ID }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-action>
              <v-icon>mdi-pirate</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Name: {{ user.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item @click.stop="disconnectUser">
            <v-list-item-action>
              <v-icon>mdi-logout</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Clear user</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
      </v-navigation-drawer>

      <v-footer
        app
        color="blue-grey"
        class="white--text"
      >
        <span>CAH With Vuetify</span>
        <v-spacer></v-spacer>
        <span>&copy; 2020</span>
      </v-footer>
    </v-app>
  </div>
</template>

<script>
import { cleanStorageValue, getStorageValue } from '@/helpers/localstorage';

export default {
  data: () => ({
    drawer: false,
    drawerUserInfo: false,
    left: false,
    right: false,
    isAdmin: false,
  }),
  computed: {
    hasUserID () {
      if (this.$store.getters.userID === null) {
        return false;
      }

      return true
    },
    user () {
      return this.$store.getters.user
    }
  },
  mounted () {
    this.isAdmin = this.getStorageValue("isAdmin")
    const userID = this.getStorageValue("userID")

    if (userID !== undefined) {
      this.$log.debug(`getUserID ${userID}`)
      this.$store.commit('setUserID', userID)
      this.$store.dispatch("getUserData")
    } else {
      this.$log.debug(`No user in LS`)
    }
  },
  methods: {
    getStorageValue,
    displayUserInformations () {
      this.drawerUserInfo = true
    },
    disconnectUser() {
      cleanStorageValue("userID")
      this.$store.dispatch("resetUserState")
    },
  },
};
</script>
