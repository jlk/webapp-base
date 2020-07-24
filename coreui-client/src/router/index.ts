import Vue from 'vue'
import Router from 'vue-router'

// Containers
const TheContainer = () => import('@/containers/TheContainer.vue')

// Views
const Dashboard = () => import('@/views/Dashboard.vue')

// Newly added views
const Things = () => import('@/views/Things.vue')

const Colors = () => import('@/views/theme/Colors.vue')
const Typography = () => import('@/views/theme/Typography.vue')

const Charts = () => import('@/views/charts/Charts.vue')
const Widgets = () => import('@/views/widgets/Widgets.vue')

// Views - Components
const Cards = () => import('@/views/base/Cards.vue')
const Forms = () => import('@/views/base/Forms.vue')
const Switches = () => import('@/views/base/Switches.vue')
const Tables = () => import('@/views/base/Tables.vue')
const Tabs = () => import('@/views/base/Tabs.vue')
const Breadcrumbs = () => import('@/views/base/Breadcrumbs.vue')
const Carousels = () => import('@/views/base/Carousels.vue')
const Collapses = () => import('@/views/base/Collapses.vue')
const Jumbotrons = () => import('@/views/base/Jumbotrons.vue')
const ListGroups = () => import('@/views/base/ListGroups.vue')
const Navs = () => import('@/views/base/Navs.vue')
const Navbars = () => import('@/views/base/Navbars.vue')
const Paginations = () => import('@/views/base/Paginations.vue')
const Popovers = () => import('@/views/base/Popovers.vue')
const ProgressBars = () => import('@/views/base/ProgressBars.vue')
const Tooltips = () => import('@/views/base/Tooltips.vue')

// Views - Buttons
const StandardButtons = () => import('@/views/buttons/StandardButtons.vue')
const ButtonGroups = () => import('@/views/buttons/ButtonGroups.vue')
const Dropdowns = () => import('@/views/buttons/Dropdowns.vue')
const BrandButtons = () => import('@/views/buttons/BrandButtons.vue')

// Views - Icons
const CoreUIIcons = () => import('@/views/icons/CoreUIIcons.vue')
const Brands = () => import('@/views/icons/Brands.vue')
const Flags = () => import('@/views/icons/Flags.vue')

// Views - Notifications
const Alerts = () => import('@/views/notifications/Alerts.vue')
const Badges = () => import('@/views/notifications/Badges.vue')
const Modals = () => import('@/views/notifications/Modals.vue')

// Views - Pages
const Page404 = () => import('@/views/pages/Page404.vue')
const Page500 = () => import('@/views/pages/Page500.vue')
const Login = () => import('@/views/pages/Login.vue')
const Register = () => import('@/views/pages/Register.vue')

// Users
const Users = () => import('@/views/users/Users.vue')
const User = () => import('@/views/users/User.vue')

Vue.use(Router)

export default new Router({
  mode: 'hash', // https://router.vuejs.org/api/#mode
  linkActiveClass: 'active',
  scrollBehavior: () => ({ x: 0, y: 0 }),
  routes: configRoutes()
})

function configRoutes () {
  return [
    {
      path: '/',
      redirect: '/dashboard',
      name: 'Home',
      component: TheContainer,
      children: [
        {
          path: 'things',
          name: 'Things',
          component: Things
        },
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: Dashboard
        },
        {
          path: 'theme',
          redirect: '/theme/colors',
          name: 'Theme',
          component: {
            render (c: (arg0: string) => any) { return c('router-view') }
          },
          children: [
            {
              path: 'colors',
              name: 'Colors',
              component: Colors
            },
            {
              path: 'typography',
              name: 'Typography',
              component: Typography
            }
          ]
        },
        {
          path: 'charts',
          name: 'Charts',
          component: Charts
        },
        {
          path: 'widgets',
          name: 'Widgets',
          component: Widgets
        },
        {
          path: 'users',
          meta: {
            label: 'Users'
          },
          component: {
            render(c: (arg0: string) => any) {
              return c('router-view')
            }
          },
          children: [
            {
              path: '',
              name: 'Users',
              component: Users
            },
            {
              path: ':id',
              meta: {
                label: 'User Details'
              },
              name: 'User',
              component: User
            }
          ]
        },
        {
          path: 'base',
          redirect: '/base/cards',
          name: 'Base',
          component: {
            render (c: (arg0: string) => any) { return c('router-view') }
          },
          children: [
            {
              path: 'cards',
              name: 'Cards',
              component: Cards
            },
            {
              path: 'forms',
              name: 'Forms',
              component: Forms
            },
            {
              path: 'switches',
              name: 'Switches',
              component: Switches
            },
            {
              path: 'tables',
              name: 'Tables',
              component: Tables
            },
            {
              path: 'tabs',
              name: 'Tabs',
              component: Tabs
            },
            {
              path: 'breadcrumbs',
              name: 'Breadcrumbs',
              component: Breadcrumbs
            },
            {
              path: 'carousels',
              name: 'Carousels',
              component: Carousels
            },
            {
              path: 'collapses',
              name: 'Collapses',
              component: Collapses
            },
            {
              path: 'jumbotrons',
              name: 'Jumbotrons',
              component: Jumbotrons
            },
            {
              path: 'list-groups',
              name: 'List Groups',
              component: ListGroups
            },
            {
              path: 'navs',
              name: 'Navs',
              component: Navs
            },
            {
              path: 'navbars',
              name: 'Navbars',
              component: Navbars
            },
            {
              path: 'paginations',
              name: 'Paginations',
              component: Paginations
            },
            {
              path: 'popovers',
              name: 'Popovers',
              component: Popovers
            },
            {
              path: 'progress-bars',
              name: 'Progress Bars',
              component: ProgressBars
            },
            {
              path: 'tooltips',
              name: 'Tooltips',
              component: Tooltips
            }
          ]
        },
        {
          path: 'buttons',
          redirect: '/buttons/standard-buttons',
          name: 'Buttons',
          component: {
            render (c: (arg0: string) => any) { return c('router-view') }
          },
          children: [
            {
              path: 'standard-buttons',
              name: 'Standard Buttons',
              component: StandardButtons
            },
            {
              path: 'button-groups',
              name: 'Button Groups',
              component: ButtonGroups
            },
            {
              path: 'dropdowns',
              name: 'Dropdowns',
              component: Dropdowns
            },
            {
              path: 'brand-buttons',
              name: 'Brand Buttons',
              component: BrandButtons
            }
          ]
        },
        {
          path: 'icons',
          redirect: '/icons/coreui-icons',
          name: 'CoreUI Icons',
          component: {
            render (c: (arg0: string) => any) { return c('router-view') }
          },
          children: [
            {
              path: 'coreui-icons',
              name: 'Icons library',
              component: CoreUIIcons
            },
            {
              path: 'brands',
              name: 'Brands',
              component: Brands
            },
            {
              path: 'flags',
              name: 'Flags',
              component: Flags
            }
          ]
        },
        {
          path: 'notifications',
          redirect: '/notifications/alerts',
          name: 'Notifications',
          component: {
            render (c: (arg0: string) => any) { return c('router-view') }
          },
          children: [
            {
              path: 'alerts',
              name: 'Alerts',
              component: Alerts
            },
            {
              path: 'badges',
              name: 'Badges',
              component: Badges
            },
            {
              path: 'modals',
              name: 'Modals',
              component: Modals
            }
          ]
        }
      ]
    },
    {
      path: '/pages',
      redirect: '/pages/404',
      name: 'Pages',
      component: {
        render (c: (arg0: string) => any) { return c('router-view') }
      },
      children: [
        {
          path: '404',
          name: 'Page404',
          component: Page404
        },
        {
          path: '500',
          name: 'Page500',
          component: Page500
        },
        {
          path: 'login',
          name: 'Login',
          component: Login
        },
        {
          path: 'register',
          name: 'Register',
          component: Register
        }
      ]
    }
  ]
}

