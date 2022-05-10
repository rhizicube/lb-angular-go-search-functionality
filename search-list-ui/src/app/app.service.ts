import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class AppService {
  constructor(private http: HttpClient) {}

  rootURL = '/api';

  // getUsers() {
  //   return this.http.get('app_backend/api/' + '/movies');
  // }

  getUsers() {
    return this.http.get(this.rootURL + '/movies');
  }

  getSpecificUser(name: any) {
    return this.http.post(this.rootURL + '/sendMovies', { name });
  }
}
