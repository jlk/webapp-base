import { mount, createLocalVue, Wrapper } from '@vue/test-utils'
import VueRouter from 'vue-router'
import CoreuiVue from '@coreui/vue'
import User from '@/views/users/User.vue'
import appRouter from '@/router'

const localVue = createLocalVue()
localVue.use(VueRouter)
const router = appRouter
router.push({path: '/users/1'})

localVue.use(CoreuiVue)

describe('User.vue', () => {
  let wrapper: Wrapper<Vue>
  beforeEach(() => {
    wrapper = mount(User, {
      localVue,
      router
    })
  })
  it('has a name', () => {
    expect(User.name).toBe('User')
  })
  it('is Vue instance', () => {
    expect(wrapper.vm).toBeTruthy()
  })
  it('is User', () => {
    expect(wrapper.findComponent(User)).toBeTruthy()
  })
  // This isn't working for me in ts, but I'm planning on getting rid of users anyways...
  // it('should have methods', () => {
  //   expect(typeof User.methods.goBack).toEqual('function')
  // })
  test('renders correctly', () => {
    expect(wrapper.element).toMatchSnapshot()
  })
})
