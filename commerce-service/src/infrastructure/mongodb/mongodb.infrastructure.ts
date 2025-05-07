import { Injectable, Logger, OnApplicationShutdown } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import mongoose, { Mongoose } from 'mongoose';

@Injectable()
export class MongodbInfrastructure implements OnApplicationShutdown {
  private connection: Mongoose;
  private readonly logger = new Logger(MongodbInfrastructure.name);


  constructor(private readonly configService: ConfigService) {
    const url = 'mongodb://' +
      this.configService.get<string>('MONGO_USERNAME') +
      ':' +
      this.configService.get<string>('MONGO_PASSWORD') +
      '@' +
      this.configService.get<string>('MONGO_HOST') +
      ':' +
      this.configService.get<number>('MONGO_PORT') +
      '/' +
      this.configService.get<string>('MONGO_DATABASE_NAME') +
      '?authSource=admin';

    mongoose.connect(url, {}).then((r) => {
      this.logger.log('MongoDB Connected');
      this.connection = r;
    }).catch((err) => {
      console.log(err);
    });

  }

  async onApplicationShutdown(signal?: string) {
    return await this.connection.disconnect();
  }

}
