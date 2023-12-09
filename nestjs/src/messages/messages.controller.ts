import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { CreateMessageDTO } from './dtos/create-message.dto';
import { MessagesService } from './messages.service';

@Controller('messages')
export class MessagesController {
  constructor(public messagesService: MessagesService) {}
  @Get()
  listMessages() {
    return this.messagesService.findAll();
  }
  @Post()
  createMessages(@Body() body: CreateMessageDTO) {
    return this.messagesService.create(body.content);
  }
  @Get(':id')
  getMessage(@Param('id') id: string) {
    return this.messagesService.findOne(id);
  }
}
