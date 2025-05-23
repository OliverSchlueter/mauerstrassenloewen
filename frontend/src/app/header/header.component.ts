import { Component } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {Router} from '@angular/router';
import {MatButton} from '@angular/material/button';
import {NgIf, NgOptimizedImage} from '@angular/common';
import {MatIcon} from '@angular/material/icon';

@Component({
  selector: 'app-header',
  imports: [
    MatButton,
    NgIf,
    MatIcon,
    NgOptimizedImage
  ],
  templateUrl: './header.component.html',
  standalone: true,
  styleUrl: './header.component.scss'
})
export class HeaderComponent {

  constructor(private authService: AuthService, private router: Router) {
  }

  goHome() {
    if(this.authService.user) {
      this.router.navigate(['/home']);
    } else {
      this.router.navigate(['/login']);
    }
  }

  goAccount() {
    this.router.navigate(['/account'])
  }

  isLoggedIn() {
    return !!this.authService.user;
  }

  goTheory() {
    this.router.navigate(['/theory']);
  }

  goCoach() {
    this.router.navigate(['/coach']);
  }
}
