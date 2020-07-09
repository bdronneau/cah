// Libraries
import Vue from 'vue'
import Vuetify from 'vuetify'

// Components
import UserHand from '@/components/UserHand.vue';

// Utilities
import {
  mount,
  createLocalVue
} from '@vue/test-utils'
Vue.use(Vuetify)

const localVue = createLocalVue()

describe('UserHand.vue', () => {
  let vuetify

  beforeEach(() => {
    vuetify = new Vuetify()
  })

  it('should have a custom title and match snapshot', () => {
    const wrapper = mount(UserHand, {
      localVue,
      vuetify,
      propsData: {
        cards: [
          {
            id: 1,
            name: "Card 1",
          },
          {
            id: 2,
            name: "Card 2",
          }
        ],
      },
    })

    // With jest we can create snapshot files of the HTML output
    expect(wrapper.html()).toMatchSnapshot()
  })
})
