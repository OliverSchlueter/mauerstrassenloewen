import { Injectable } from '@angular/core';
import {User} from '../models/User';
import {Profile} from '../models/Profile';
import {HttpClient, HttpHeaders, HttpResponse} from '@angular/common/http';
import {catchError, map, Observable, of, switchMap} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  url = "http://localhost:8082/api/v1"
  user: User | undefined;
  authToken: string = ""
  headers = new HttpHeaders({
    "Content-Type": "application/json",
    "Accept": "application/json",
  });

  constructor(private http: HttpClient) { }


  login(username: string, password: string): Observable<boolean> {
    return this.getAuthToken(username, password).pipe(
      switchMap(tokenSuccess => {
        if (tokenSuccess) {
          // Wenn Token OK, dann getUserMe aufrufen
          return this.getUserMe();
        } else {
          // Token fehlgeschlagen -> false zurückgeben
          return of(false);
        }
      })
    );
  }




  getAuthToken(username: string, password: string): Observable<boolean> {
    const tokenHeaders = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "text/plain",
      "X-Auth-Username": username,
      "X-Auth-Password": password
    });

    return this.http.post(this.url + "/auth-token", {}, {
      headers: tokenHeaders,
      observe: "response",
      responseType: "text" as const
    }).pipe(
      map(response => {
        if (response.status === 201 && response.body) {
          this.authToken = response.body;
          return true;
        }
        return false;
      }),
      catchError(err => {
        console.error("Token Error:", err);
        return of(false);
      })
    );
  }


  getUserMe(): Observable<boolean> {
    const userHeaders = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "application/json",
      "X-Auth-Token": this.authToken
    });

    return this.http.get<User>(this.url + "/user/me", {
      headers: userHeaders,
      observe: "response"
    }).pipe(
      map(response => {
        if (response.status === 200 && response.body) {
          this.user = response.body;
          console.log(response.body);
          return true;
        }
        return false;
      }),
      catchError(err => {
        console.error("Fehler bei getUserMe:", err);
        return of(false);
      })
    );
  }

  isLoggedIn() {
    return this.user !== undefined;
  }


}
