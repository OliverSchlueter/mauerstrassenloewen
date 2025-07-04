import {Component, OnInit} from '@angular/core';
import {NgClass, NgFor, NgIf} from '@angular/common';
import {MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';
import {AuthService} from '../services/auth.service';
import {AiService} from '../services/ai.service';
import {FormsModule} from '@angular/forms';
import {Message} from '../models/Message';

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
  chatID = "";
  messages: Message[] = [];

  constructor(private aiService: AiService, private authService: AuthService) {}

  ngOnInit() {

  }

  startChat() {
    const user = this.authService.user?.name
    if(user) {
      console.log("message input", this.messageInput)
      this.aiService.startChat(this.messageInput).subscribe(messages => {
        this.chatID = messages.id;
        this.messages = messages.messages;
        console.log(messages)
      })
    }
  }

  sendMessage() {
    if (this.messages.length === 0) {
      this.startChat();
      return;
    }

    const user = this.authService.user?.name
    if(user) {
      this.aiService.sendMessage(this.chatID, this.messageInput).subscribe(messages => {
        this.messages = messages.messages;
        console.log(messages)
      })
    }
  }

  getClass(message: Message) {
    if(message.role === 'user') {
      return 'fromUser'
    } else {
      return 'fromBot'
    }
  }
}
