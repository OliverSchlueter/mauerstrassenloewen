import { Component } from '@angular/core';
import {MatButton} from '@angular/material/button';
import {MatError, MatFormField, MatInput, MatLabel} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {MatCard, MatCardContent, MatCardTitle} from '@angular/material/card';
import {Router} from '@angular/router';
import {AuthService} from '../services/auth.service';
import {NgIf} from '@angular/common';

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
    MatCardTitle,
    MatError,
    NgIf
  ],
  templateUrl: './login.component.html',
  standalone: true,
  styleUrl: './login.component.scss'
})
export class LoginComponent {
  errorMessage: string = "";
  username: string = "";
  password: string = "";

  constructor(private router: Router, private authService: AuthService) {

  }

  login() {
    if(this.authService.login(this.username, this.password)) {
      this.router.navigate(['home'])
    }
    else {
      this.errorMessage = "wrong username or password"
    }
  }

  register() {
    this.router.navigate(['register'])
  }
}
