import { Component } from '@angular/core';
import {MatCard, MatCardContent, MatCardHeader, MatCardSubtitle, MatCardTitle} from '@angular/material/card';
import {MatProgressSpinner} from '@angular/material/progress-spinner';
import {MatProgressBar} from '@angular/material/progress-bar';
import {CircularProgressComponent} from '../util/circular-progress/circular-progress.component';
import {AuthService} from '../services/auth.service';
import {User} from '../models/User';

@Component({
  selector: 'app-home',
  imports: [
    MatCard,
    MatCardTitle,
    MatCardContent,
    MatCardSubtitle,
    MatCardHeader,
    MatProgressSpinner,
    MatProgressBar,
    CircularProgressComponent
  ],
  templateUrl: './home.component.html',
  standalone: true,
  styleUrl: './home.component.scss'
})
export class HomeComponent {

  constructor(private authService: AuthService) {
  }

  getUserName(): string {
    if(this.authService.user) {
      return this.authService.user?.name
    } else {
      return '';
    }
  }
}
