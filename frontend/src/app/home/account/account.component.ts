import {Component, OnInit} from '@angular/core';
import {MatDivider} from '@angular/material/divider';
import {Router, RouterOutlet} from '@angular/router';
import {MatButton} from '@angular/material/button';
import {AuthService} from '../../services/auth.service';
import {User} from '../../models/User';

@Component({
  selector: 'app-account',
  imports: [
    MatDivider,
    RouterOutlet,
    MatButton
  ],
  templateUrl: './account.component.html',
  styleUrl: './account.component.scss'
})
export class AccountComponent implements OnInit{
  user: User | undefined;

  constructor(private router: Router, private authService: AuthService) {
  }

  ngOnInit() {
    this.user = this.authService.user
  }

  navigate(location: string) {
    this.router.navigate(['account/'+location]);
  }
}
