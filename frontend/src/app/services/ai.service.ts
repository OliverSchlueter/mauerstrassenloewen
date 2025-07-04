import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {AuthService} from './auth.service';
import {Message} from '../models/Message';
import {StartChatDTO} from '../models/StartChatDTO';

@Injectable({
  providedIn: 'root'
})
export class AiService {
  url  = "http://localhost:8082/api/v1"

  constructor(private http: HttpClient, private authService: AuthService) { }

  createSystemMessage() {
    return "You are a professional AI Assistent. Please answer the users message."
  }

  startChat(message: string) {
    const body = {
      "system_msg": this.createSystemMessage(),
      "user_msg": message
    }

    const headers = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "application/json",
      "X-Auth-Token": this.authService.authToken
    });

    return this.http.post<StartChatDTO>(this.url + "/chatbot/chat", body, {headers: headers});
  }

  sendMessage(chatID: string, message: string) {
    const body = {
      "chat_id": chatID,
      "user_msg": message
    }

    const headers = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "application/json",
      "X-Auth-Token": this.authService.authToken
    });

    return this.http.post<StartChatDTO>(this.url + "/chatbot/chat/"+chatID+"/new-message", body, {headers: headers})
  }
}
