import { Injectable } from '@angular/core';
import {User} from '../models/User';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  user: User | undefined;

  constructor() { }

  login(name: string, password: string): boolean {
    // Simulate a login by setting a user object
    if(name === 'testuser' && password === '123') {
      this.user = {
        id: "1",
        name: name,
        email: 'testmail@gmail.com',
        password: password
      }
      return true;
    }
    else {
      return false;
    }
  }

  isLoggedIn() {
    return this.user !== undefined;
  }
}
