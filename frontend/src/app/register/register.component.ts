import { Component } from '@angular/core';
import {InterviewComponent} from '../interview/interview.component';
import {User} from '../models/User';
import {MatTab, MatTabGroup} from '@angular/material/tabs';
import {MatCard, MatCardContent, MatCardHeader, MatCardTitle} from '@angular/material/card';
import {MatFormField, MatInput, MatLabel} from '@angular/material/input';
import {FormsModule} from '@angular/forms';

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
    FormsModule
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss'
})
export class RegisterComponent {
  user= new User()
  repeatPassword = '';

  register() {

  }

  getNewUser() {
    return new User();
  }
}
