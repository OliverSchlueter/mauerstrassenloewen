import { Component } from '@angular/core';
import {NgClass, NgFor, NgIf} from '@angular/common';
import {Message} from '../models/Message'
import {MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';

@Component({
  selector: 'app-coach',
  imports: [
    NgFor,
    NgClass,
    MatIconButton,
    MatIcon,
    NgIf
  ],
  templateUrl: './coach.component.html',
  standalone: true,
  styleUrl: './coach.component.scss'
})
export class CoachComponent {
  messages: Message[] = [
    {
      content: "testesetsetsetsetsetsetsetsetset",
      fromUser: true,
    },
    {
      content: "loplolololololololololololololol",
      fromUser: false,
    },
    {
      content: "so uncivilized",
      fromUser: true,
    },
  ]

  getClass(message: Message) {
    if(message.fromUser) {
      return 'fromUser'
    } else {
      return 'fromBot'
    }
  }
}
