import http from "../http-common";

class ThingDataService {
  getThings() {
    return http.get("/things");
  }
}
export default new ThingDataService();