import { Module } from '@nestjs/common';
import { CacheModule } from '@nestjs/cache-manager';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UserController } from './users/users.controller';
import { UserService } from './users/users.service';
import { UserCacheRepository } from './users/user.cache.repository';
import { USER_REPOSITORY } from './users/user.repository.token';
import KeyvRedis from '@keyv/redis';
import Keyv from 'keyv';

@Module({
  imports: [
    CacheModule.registerAsync({
      useFactory: async () => {
        const keyvRedis = new KeyvRedis('redis://localhost:6379');
        const keyv = new Keyv({ store: keyvRedis, ttl: 600 });
        return {
          store: keyv,
        };
      },
    }),
  ],

  controllers: [AppController, UserController],
  providers: [AppService, UserService,
    {
      provide: USER_REPOSITORY,
      useClass: UserCacheRepository
    }
  ],
})
export class AppModule { }
