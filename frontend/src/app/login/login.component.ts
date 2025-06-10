import {Component, signal} from '@angular/core';
import {MatButton, MatIconButton} from '@angular/material/button';
import {MatError, MatFormField, MatInput, MatLabel, MatSuffix} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {MatCard, MatCardContent, MatCardTitle} from '@angular/material/card';
import {Router} from '@angular/router';
import {AuthService} from '../services/auth.service';
import {NgIf} from '@angular/common';
import {MatIcon} from '@angular/material/icon';

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
    NgIf,
    MatIcon,
    MatSuffix,
    MatIconButton,
  ],
  templateUrl: './login.component.html',
  standalone: true,
  styleUrl: './login.component.scss'
})
export class LoginComponent {
  hidePassword = signal(true)
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

  clickPassword($event: MouseEvent) {
    this.hidePassword.set(!this.hidePassword);
    $event.stopPropagation()
  }
}
