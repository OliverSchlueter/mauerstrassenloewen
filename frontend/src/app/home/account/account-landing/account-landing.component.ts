import {Component, Input, OnInit} from '@angular/core';
import {User} from '../../../models/User';
import {AuthService} from '../../../services/auth.service';

@Component({
  selector: 'app-account-landing',
  imports: [],
  templateUrl: './account-landing.component.html',
  standalone: true,
  styleUrl: './account-landing.component.scss'
})
export class AccountLandingComponent implements OnInit{
  user: User | undefined;

  constructor(private authService: AuthService) {
  }

  ngOnInit() {
    this.user = this.authService.user;
  }
}
