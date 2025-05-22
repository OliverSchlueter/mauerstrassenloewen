import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Profile} from '../models/Profile';
import {User} from '../models/User';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  url = "http://localhost:8082/api/v1"
  token = "msl_932ee7e7-3881-4810-b08f-be36cceea50f"

  constructor(private http: HttpClient) { }

  updateUser(user: User) {

    const headers = new HttpHeaders({
      "X-Auth-Token": this.token,
      "Content-Type": "application/json",
      "Accept": "application/json",
    });

    return this.http.put<User>(this.url + "/user", user, {headers});
  }

  register(user: User) {
    const headers = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "application/json",
    });

    return this.http.post(this.url + "/user/register", user, {headers})
  }
}
