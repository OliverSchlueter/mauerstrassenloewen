import { Component } from '@angular/core';
import {InterviewComponent} from '../interview/interview.component';
import {User} from '../models/User';
import {MatTab, MatTabGroup} from '@angular/material/tabs';
import {MatCard, MatCardContent, MatCardFooter, MatCardHeader, MatCardTitle} from '@angular/material/card';
import {MatFormField, MatInput, MatLabel} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {MatButton, MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';
import {Router} from '@angular/router';
import {UserService} from '../services/user.service';

@Component({
  selector: 'app-register',
  imports: [
    InterviewComponent,
    MatTabGroup,
    MatTab,
    MatCard,
    MatCardTitle,
    MatCardHeader,
    MatCardContent,
    MatFormField,
    MatLabel,
    MatInput,
    FormsModule,
    MatCardFooter,
    MatButton,
    MatIcon,
    MatIconButton
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss'
})
export class RegisterComponent {
  user= new User()
  repeatPassword = '';

  constructor(private router: Router, private userService: UserService) {
  }

  register() {
    this.userService.register(this.user).subscribe((response: any) => {
      console.log(response)
      if(response) {

      }
      this.router.navigate(['login'])
    })
  }

  getNewUser() {
    return new User();
  }

  navigateLogin() {
    this.router.navigate(['login'])
  }
}
