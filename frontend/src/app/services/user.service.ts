import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Profile} from '../models/Profile';
import {User} from '../models/User';
import {AuthService} from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  url = "http://localhost:8082/api/v1"
  token = ""

  constructor(private http: HttpClient, private authService: AuthService) { }


  updateUser(user: User) {
    this.token = this.authService.authToken

    const headers = new HttpHeaders({
      "X-Auth-Token": this.token,
      "Content-Type": "application/json",
      "Accept": "application/json",
    });

    return this.http.put<User>(this.url + "/user/" + user.id, user, {headers});
  }


  register(user: User) {
    const headers = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "application/json",
    });

    return this.http.post(this.url + "/user/register", user, {headers, observe: 'response'})
  }
}
