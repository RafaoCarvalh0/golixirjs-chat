import { Inject, Injectable } from '@nestjs/common';
import { CACHE_MANAGER } from '@nestjs/cache-manager';
import { UserRepository } from './user.repository';
import { User } from './user.entity';
import type { Cache } from 'cache-manager';


@Injectable()
export class UserCacheRepository implements UserRepository {
    constructor(@Inject(CACHE_MANAGER) private cacheManager: Cache) { }

    async createUser(user: User): Promise<void> {
        await this.cacheManager.set(user.id, user);
    }

    async deleteUser(userId: string): Promise<void> {
        await this.cacheManager.del(userId);
    }

    async getUserByUsername(username: string): Promise<any> {
        await this.cacheManager.get(username);
    }
}