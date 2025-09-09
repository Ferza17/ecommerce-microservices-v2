import { CallHandler, ExecutionContext, Injectable, NestInterceptor } from '@nestjs/common';
import { Observable } from 'rxjs';
import { Header } from '../../../enum/header'
import { v4 as uuidv4 } from 'uuid';


@Injectable()
export class RequestIdInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    const metadata = context.switchToRpc().getContext();

    let requestId: string = metadata.get(Header.X_REQUEST_ID)[0];
    if (!requestId) {
      requestId = uuidv4();
    }

    metadata.set(Header.X_REQUEST_ID, requestId);
    return next.handle();
  }
}
