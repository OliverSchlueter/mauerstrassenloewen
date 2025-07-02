export class StartChatDTO {
  id: string = "";
  messages: Message[] = []
}

export class Message {
  role: string = ""
  content: string = ""
  sent_at: string = ""
}
