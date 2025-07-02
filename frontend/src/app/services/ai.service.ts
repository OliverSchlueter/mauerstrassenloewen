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

  startChat(user: string, message: string) {
    const body = {
      "system_msg": this.createSystemMessage(),
      "user_msg": message
    }

    const headers = new HttpHeaders({
      "Content-Type": "application/json",
      "Accept": "application/json",
      "X-Auth-Token": this.authService.authToken
    });

    return this.http.post<StartChatDTO>(this.url + "/start-chat", body, {headers: headers})
  }
}
