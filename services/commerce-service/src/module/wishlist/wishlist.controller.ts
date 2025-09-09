import { Controller, Logger } from '@nestjs/common';
import { WishlistService } from './wishlist.service';

@Controller('wishlist-controller')
export class WishlistController {
  private readonly logger = new Logger(WishlistController.name);

  constructor(
    private readonly wishlistService: WishlistService,
  ) {}
}
