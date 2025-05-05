import { Component } from '@angular/core';
import {MatButton} from '@angular/material/button';
import {MatFormField, MatInput, MatLabel} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {MatCard, MatCardContent, MatCardTitle} from '@angular/material/card';
import {Router} from '@angular/router';

@Component({
  selector: 'app-login',
  imports: [
    MatButton,
    MatFormField,
    FormsModule,
    MatInput,
    MatLabel,
    MatCard,
    MatCardContent,
    MatCardTitle
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {
  loggedIn = false;
  username: string = "";
  password: string = "";

  constructor(private router: Router) {

  }

  login() {
    this.loggedIn = true;
    if(this.loggedIn) {
      this.router.navigate(['home'])
    }
  }

  register() {
    this.router.navigate(['register'])
  }
}
