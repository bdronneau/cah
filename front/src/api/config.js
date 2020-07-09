import axios from 'axios'
import Vue from 'vue'

export function getConfig () {
  return axios({ method: 'GET', url: '/config' }).then(result => {
    return result.data
  }).catch(error => {
    Vue.$log.error(`Can not load config from url ${error}`)
  })
}
