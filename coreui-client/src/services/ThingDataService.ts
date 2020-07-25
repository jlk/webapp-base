import http from "../http-common";

class ThingDataService {
  getThings() {
    return http.get("/devices");
  }
}
export default new ThingDataService();