import Vue, { VNode } from 'vue'
import { CRenderFunction } from '@coreui/vue'

declare module 'vue/types/options' {
  interface ComponentOptions<V extends Vue> {
    icons?: { [key: string]: string[] }
  }
}

declare global {
  namespace JSX {
    // tslint:disable no-empty-interface
    interface Element extends VNode {}
    // tslint:disable no-empty-interface
    interface ElementClass extends Vue {}
    interface IntrinsicElements {
      [elem: string]: any
    }
  }
  type ContentToRender = CRenderFunction['contentToRender']
}
