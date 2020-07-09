import axios from 'axios'
import Vue from 'vue';
import store from '@/store/index';

const API_BASE = store.getters.config.api

var axiosInstance = axios.create({
  baseURL: `${API_BASE}/api`
});

export function retrieveRoomStatus(id) {
  return axiosInstance.get(`/rooms/${id}`).then(response => {
      return response.data
    })
    .catch((error) => {
      Vue.$log.error("On retrieve room", error)
      return Promise.reject(error)
    });
}

export function getCardsByUserID(roomID, userID) {
  return axiosInstance.get(`/rooms/${roomID}/users/${userID}/cards`).then(response => {
      return response.data
    })
    .catch((error) => {
      Vue.$log.error("On retrieve user cards", error)
      return Promise.reject(error)
    });
}

export function getUsersLobby(roomID) {
  return axiosInstance.get(`/rooms/${roomID}/users`).then(response => {
      return response.data
    }).catch(error => {
      Vue.$log.error("On retrieve users", error)
      return Promise.reject(error)
    });
}

export function listRooms() {
  return axiosInstance.get(`/rooms`).then(response => {
      return response.data
    }).catch(error => {
      Vue.$log.error("On retrieve roooms", error)
      return Promise.reject(error)
    });
}

export function headRoom(roomID) {
  return axiosInstance.head(`/rooms/${roomID}`).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On head room", error)
      return Promise.reject(error)
    });
}

export function addUserInRoom(userID, roomID) {
  return axiosInstance.post(
    `/rooms/${roomID}/users`,
    {
      id: userID,
    }).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On join room", error)
      return Promise.reject(error)
    });
}

export function postNewUser(name) {
  return axiosInstance.post(
    `/users`,
    {
      name: name,
    }).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On create user", error)
      return Promise.reject(error)
    });
}

export function postStartRoom(roomID) {
  return axiosInstance.post(
    `/rooms/${roomID}/start`).then(response => {
      return Promise.resolve(response)
    }).catch(error => {
      Vue.$log.error("On start party", error)
      return Promise.reject(error)
    });
}

export function getJudgeHand(roomID) {
  return axiosInstance.get(
    `/rooms/${roomID}/judge`).then(response => {
      return Promise.resolve(response)
    }).catch(error => {
      Vue.$log.error("On get judge hand", error)
      return Promise.reject(error)
    });
}

export function postUsersCards(roomID, userID, cardID) {
  return axiosInstance.post(
    `/rooms/${roomID}/users/${userID}/card`,
    {
      user_id: userID,
      card_id: cardID
    }).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On post card", error)
      return Promise.reject(error)
    });
}

export function postJudgeVote(roomID, userID, cardID) {
  return axiosInstance.post(
    `/rooms/${roomID}/users/${userID}/elected`,
    {
      user_id: userID,
      card_id: cardID
    }).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On post card", error)
      return Promise.reject(error)
    });
}

export function postNextTurn(roomID) {
  return axiosInstance.post(`/rooms/${roomID}/next`).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On post next", error)
      return Promise.reject(error)
    });
}

export function getRoomStats(roomID) {
  return axiosInstance.get(`/rooms/${roomID}/stats`).then(response => {
      return response.data
    }).catch(error => {
      Vue.$log.error("On getRoomStats", error)
      return Promise.reject(error)
    });
}

export function postNewRoom(roomName) {
  return axiosInstance.post(
    `/rooms`,
    {
      description: roomName
    }).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On post room", error)
      return Promise.reject(error)
    });
}

export function postEndRoom(roomID) {
  return axiosInstance.post(`/rooms/${roomID}/stop`).then(response => {
      return response
    }).catch(error => {
      Vue.$log.error("On post end game", error)
      return Promise.reject(error)
    });
}

export function getUser(id) {
  return axiosInstance.get(`/users/${id}`).then(response => {
      return response.data
    }).catch(error => {
      Vue.$log.error("On getUser", error)
      return Promise.reject(error)
    });
}
