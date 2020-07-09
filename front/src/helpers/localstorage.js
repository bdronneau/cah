import Vue from 'vue';

export function getStorageValue (key) {
    if (localStorage.getItem(key)) {
        try {
            return JSON.parse(localStorage.getItem(key))
        } catch (e) {
            Vue.$log.error(`${key} is not a JSON clear it`)
            localStorage.removeItem(key)
        }
    }
}

export function cleanStorageValue (key) {
  Vue.$log.debug(`Delete ${key}`)
  if (localStorage.getItem(key)) {
      localStorage.removeItem(key)
  }
}

export function persist(key, value) {
  Vue.$log.debug(`Persist ${key} with ${value}`)
  localStorage.setItem(key, JSON.stringify(value))
}
