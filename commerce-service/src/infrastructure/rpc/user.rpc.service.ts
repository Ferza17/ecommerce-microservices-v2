import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { Service } from '../../enum/service';
import { ClientGrpc } from '@nestjs/microservices';
import { UserServiceService } from '../../model/rpc/userService';
import { FindUserByIdRequest, User } from '../../model/rpc/userMessage';
import { lastValueFrom } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';


@Injectable()
export class UserRpcService implements OnModuleInit {
  private readonly logger = new Logger(UserRpcService.name);
  private userService: any;
  constructor(@Inject(Service.UserService.toString()) private client: ClientGrpc) {
  }
  onModuleInit() {
    this.userService = this.client.getService<typeof UserServiceService>('UserService');
  }

  async findUserById(requestId: string, req: FindUserByIdRequest): Promise<User> {
    try {
      const metadata = new Metadata();
      metadata.set(Header.X_REQUEST_ID, requestId);
      return lastValueFrom(this.userService.findUserById(req, metadata));
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}