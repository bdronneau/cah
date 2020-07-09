import Vue from 'vue';
import Vuex from 'vuex';

import { persist } from '@/helpers/localstorage';
import { getUser, retrieveRoomStatus } from '@/api/main';


Vue.use(Vuex);

const getDefaultState = () => {
  return {
    config: {
      api: 'http://localhost:1324'
    },
    user: {
      ID: null,
      judge: false,
      inRoom: false,
      name: null,
    },
    room: {
      id: null,
      name: "N/A",
      status: "N/A",
      turn: null,
      finished: false,
      raw: {
        current_card: "",
      },
    },
    users: []
  }
}
const state = getDefaultState()

export default new Vuex.Store({
  state,
  mutations: {
    setUserID (state, userID) {
      state.user.ID = userID
      persist("userID", state.user.ID)
    },
    setJudge (state, value) {
      state.user.judge = value
    },
    nextTurn (state) {
      state.room.turn++
    },
    setRoom (state, value) {
      state.room = Object.assign(state.room, value)
    },
    setUser (state, value) {
      state.user = Object.assign(state.user, value)
    },
    setEndGame (state) {
      state.room.finished = true
    },
    setUsers (state, users) {
      state.users = users
      state.user.inRoom = false
      users.forEach(element => {
        if (state.user.ID === element.id) {
          state.user.inRoom = true
          return
        }
      })
    },
    updateConfig: (state, payload) => {
      state.config = payload
    },
  },
  actions: {
    resetState (state) {
      Object.assign(state, getDefaultState())
    },
    resetUserState ({ commit }) {
      commit("setUser", getDefaultState().user)
    },
    getRoomData ({ commit, getters }) {
      retrieveRoomStatus(getters.getRoom.id).then(data => {
        commit("setRoom", { raw: data })
      })
    },
    getUserData ({ commit, getters }) {
      getUser(getters.userID).then(data => {
        commit("setUser", { ID: data.id, name: data.name })
      })
    },
  },
  modules: {
  },
  getters: {
    userID: state => {
      return state.user.ID
    },
    user: state => {
      return state.user
    },
    isJudge: state => {
      return state.user.judge
    },
    isInRoom: state => {
      return state.user.inRoom
    },
    getRoom: state => {
      return state.room
    },
    users: state => {
      return state.users
    },
    getRoomInformations: state => {
      return {
        turn: state.room.raw.turn,
        stats: state.room.raw.status,
        name: state.room.raw.name,
        status: state.room.raw.status,
      }
    },
    config: state => {
      return state.config
    },
  },
});
