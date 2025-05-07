import { Component } from '@angular/core';
import {InterviewComponent} from '../interview/interview.component';
import {User} from '../models/User';

@Component({
  selector: 'app-register',
  imports: [
    InterviewComponent
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss'
})
export class RegisterComponent {

  register() {

  }

  getNewUser() {
    return new User();
  }
}
