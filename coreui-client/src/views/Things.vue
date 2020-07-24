<template>
  <div>
    <CRow>
      <CCol md="12">
        <CCard>
          <CCardHeader>
            Things
          </CCardHeader>
          <CCardBody>
            <CDataTable
              class="mb-0 table-outline"
              hover
              :items="things.data"
              :fields="tableFields"
              head-color="light"
              no-sorting
            >
            </CDataTable>
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  </div>
</template>

<script lang="ts">

import ThingDataService from "../services/ThingDataService";
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Things extends Vue {
  private things: any[] = [];
  private currentIndex = -1;
  private title = "";
  public name = "Things";
  
  private tableFields: {key: string, label: string, _classes:string}[] = [
    { key: 'addr', label: 'Address', _classes: 'text-center' },
    { key: 'name', label: 'Full name', _classes: '' },
    { key: 'phone', label: 'Phone Number', _classes: 'text-center' },
  ]

  retrieveThings() {
      ThingDataService.getThings()
        .then(response => {
          this.things = response.data;
          console.log(response.data);
        })
        .catch(e => {
          console.log(e);
        });
  }

  mounted() {
    this.retrieveThings();
  }
}
</script>
