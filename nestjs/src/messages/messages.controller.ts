import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { CreateMessageDTO } from './dtos/create-message.dto';

@Controller('messages')
export class MessagesController {
  @Get()
  listMessages() {
    return 'list';
  }
  @Post()
  createMessages(@Body() body: CreateMessageDTO) {
    console.log(body.content);
    return 'create: ' + body.content;
  }
  @Get(':id')
  getMessage(@Param('id') id: string) {
    return 'id: ' + id;
  }
}
