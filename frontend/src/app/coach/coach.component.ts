import {Component, OnInit} from '@angular/core';
import {NgClass, NgFor, NgIf} from '@angular/common';
import {Message} from '../models/Message'
import {MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';
import {AuthService} from '../services/auth.service';
import {AiService} from '../services/ai.service';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-coach',
  imports: [
    NgFor,
    NgClass,
    MatIconButton,
    MatIcon,
    NgIf,
    FormsModule
  ],
  templateUrl: './coach.component.html',
  standalone: true,
  styleUrl: './coach.component.scss'
})
export class CoachComponent implements OnInit {
  messageInput = "";
  messages: Message[] = [];

  constructor(private aiService: AiService, private authService: AuthService) {}

  ngOnInit() {
    const user = this.authService.user?.name
    if(user) {
      this.aiService.getChatByUser(user, this.messageInput).subscribe(messages => {
        console.log(messages)
        this.messages = messages[0]
      })
    }
  }

  getClass(message: Message) {
    if(message.fromUser) {
      return 'fromUser'
    } else {
      return 'fromBot'
    }
  }
}
