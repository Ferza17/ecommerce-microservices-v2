import { Injectable, Logger } from '@nestjs/common';

@Injectable()
export class WishlistMongoRepository {
  private readonly logger = new Logger(WishlistMongoRepository.name);

}