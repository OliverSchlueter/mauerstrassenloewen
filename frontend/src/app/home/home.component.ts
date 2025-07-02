import { Component } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {MatButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';
import {Router} from '@angular/router';

@Component({
  selector: 'app-home',
  imports: [
    MatButton,
    MatIcon
  ],
  templateUrl: './home.component.html',
  standalone: true,
  styleUrl: './home.component.scss'
})
export class HomeComponent {

  constructor(private authService: AuthService, private router: Router) {
  }

  getUserName(): string {
    if(this.authService.user) {
      return this.authService.user?.name
    } else {
      return '';
    }
  }

  navigate(location: string) {
    this.router.navigate([location]);
  }
}
